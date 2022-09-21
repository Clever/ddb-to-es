package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Clever/ddb-to-es/es"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

//go:generate $PWD/bin/go-bindata -pkg $GOPACKAGE -o bindata.go kvconfig.yml
//go:generate gofmt -w bindata.go

var log = logger.New(os.Getenv("APP_NAME"))

// FailOnError specifies if to only log errors or to also return to lambda
var FailOnError bool
var DBClient es.DB

// ErrNoRecords is an example error you could generate in handling an event.
var ErrNoRecords = errors.New("no records contained in event")

// Handler is your Lambda function handler.
// The return signature can be empty, a single error, or a return value (struct or string) and error.
func Handler(ctx context.Context, event events.DynamoDBEvent) error {
	log.InfoD("process-records-start", logger.M{"count": len(event.Records)})
	if docs, err := processRecords(event.Records, DBClient); err != nil {
		if FailOnError {
			return err
		}
		errorMsg := err.Error()
		if len(errorMsg) > 50 {
			errorMsg = errorMsg[:50]
		}
		log.InfoD("process-records-failure", logger.M{
			"error": errorMsg,
		})
	} else {
		log.InfoD("process-records-success", logger.M{"count": len(event.Records), "docs": len(docs)})
	}

	return nil
}

func main() {
	err := logger.SetGlobalRoutingFromBytes(MustAsset("kvconfig.yml"))
	if err != nil {
		log.ErrorD("kvconfig-err", logger.M{"error": err})
		os.Exit(1)
	}
	if FailOnError, err = strconv.ParseBool(os.Getenv("FAIL_ON_ERROR")); err != nil {
		FailOnError = false
	}

	// Parse a comma separated list of Elasticsearch indices.
	rawESIndices := strings.Split(os.Getenv("ELASTICSEARCH_INDICES"), ",")
	esIndices := []string{}
	for _, index := range rawESIndices {
		if index != "" {
			esIndices = append(esIndices, strings.TrimSpace(index))
		}
	}
	if len(esIndices) < 1 {
		log.Error("missing-elasticsearch-indices")
		os.Exit(1)
	}
	esURL := os.Getenv("ELASTICSEARCH_URL")

	dbConfig := &es.DBConfig{URL: esURL}
	DBClient, err = es.NewDB(dbConfig, esIndices, log)
	if err != nil {
		log.ErrorD("elasticsearch-connect-error", logger.M{
			"message": err.Error(),
			"url":     dbConfig.URL,
		})
		os.Exit(1)
	}

	lambda.Start(Handler)
}

// processRecords converts DynamoDB stream records to es.Doc and writes them to the db
func processRecords(records []events.DynamoDBEventRecord, db es.DB) ([]es.Doc, error) {
	if len(records) == 0 {
		return nil, ErrNoRecords
	}

	docs := []es.Doc{}
	// TODO: we can parallalize this
	for _, record := range records {
		id, err := toId(record.Change.Keys)
		if err != nil {
			return nil, err
		}
		item := map[string]interface{}{}
		for k, v := range record.Change.NewImage {
			if i := toItem(v, k); i != nil {
				item[santizeKey(k)] = i
			}
		}
		switch events.DynamoDBOperationType(record.EventName) {
		case events.DynamoDBOperationTypeInsert:
			docs = append(docs, es.Doc{Op: es.OpTypeInsert, ID: id, Item: item})
		case events.DynamoDBOperationTypeModify:
			docs = append(docs, es.Doc{Op: es.OpTypeUpdate, ID: id, Item: item})
		case events.DynamoDBOperationTypeRemove:
			docs = append(docs, es.Doc{Op: es.OpTypeDelete, ID: id, Item: item})
		case "":
			continue
		default:
			return nil, fmt.Errorf("Unsupported eventName %s", record.EventName)
		}
	}

	log.InfoD("write-docs-start", logger.M{"count": len(docs)})
	if err := db.WriteDocs(docs); err != nil {
		return nil, err
	}

	return docs, nil
}

// toId generates a deterministic Id for each record
func toId(ddbKeys map[string]events.DynamoDBAttributeValue) (string, error) {
	values := []string{}
	keysSorted := []string{}
	for k := range ddbKeys {
		keysSorted = append(keysSorted, k)
	}
	sort.Strings(keysSorted)
	for _, k := range keysSorted {
		key := ddbKeys[k]
		item := toItem(key, "")
		if key.DataType() == events.DataTypeMap ||
			key.DataType() == events.DataTypeList ||
			key.DataType() == events.DataTypeBinary ||
			key.DataType() == events.DataTypeBinarySet ||
			key.DataType() == events.DataTypeNumberSet ||
			key.DataType() == events.DataTypeStringSet {

			val, err := json.Marshal(item)
			if err != nil {
				return "", err
			}
			values = append(values, string(val[:]))
		} else {
			switch item.(type) {
			//case int:
			// TODO: aws-sdk-go treats number as string
			case string:
				values = append(values, item.(string))
			case bool:
				values = append(values, strconv.FormatBool(item.(bool)))
			}
		}
	}

	return strings.Join(values, "|"), nil
}

// toItem recursively walks through DynamoDBAttributeValue
// to convert it to a standard object
func toItem(value events.DynamoDBAttributeValue, pathSoFar string) interface{} {
	switch value.DataType() {
	case events.DataTypeList:
		doc := []interface{}{}
		for _, item := range value.List() {
			if i := toItem(item, pathSoFar); i != nil {
				doc = append(doc, i)
			}
		}
		return doc
	case events.DataTypeMap:
		doc := map[string]interface{}{}
		for k, v := range value.Map() {
			path := fmt.Sprintf("%s.%s", pathSoFar, k)
			// When we send workflows to ES, including the state machine explodes the number of fields.
			if path == "Workflow.workflowDefinition.stateMachine" {
				continue
			}
			if i := toItem(v, path); i != nil {
				doc[santizeKey(k)] = i
			}
		}
		return doc
	case events.DataTypeNull:
		return nil
	case events.DataTypeNumber:
		return value.Number()
	case events.DataTypeNumberSet:
		return value.NumberSet()
	case events.DataTypeBinary:
		return value.Binary()
	case events.DataTypeBoolean:
		return value.Boolean()
	case events.DataTypeBinarySet:
		return value.BinarySet()
	case events.DataTypeString:
		return value.String()
	case events.DataTypeStringSet:
		return value.StringSet()
	default:
		return nil
	}

}

// santizeKey makes sure that document keys meet Elasticsearch requirements
func santizeKey(key string) string {
	if _, ok := es.ESReservedFields[key]; ok {
		// add another _ as prefix
		return fmt.Sprintf("_%s", key)
	}
	return key
}
