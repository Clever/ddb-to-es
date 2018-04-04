package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
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
var indexPattern = regexp.MustCompile("arn:aws:dynamodb:.*?:.*?:table/([0-9a-zA-Z_-]+)/.+")

// FailOnError specifies if to only log errors or to also return to lambda
var FailOnError bool
var IndexPrefix string
var DBClient es.DB

// ErrNoRecords is an example error you could generate in handling an event.
var ErrNoRecords = errors.New("no records contained in event")

// Handler is your Lambda function handler.
// The return signature can be empty, a single error, or a return value (struct or string) and error.
func Handler(ctx context.Context, event events.DynamoDBEvent) error {
	if err := processRecords(event.Records, DBClient); err != nil {
		if FailOnError {
			return err
		}
		log.CounterD("process-records-failure", 1, logger.M{
			"error": err.Error(),
		})
	} else {
		log.Counter("process-records-success")
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
	IndexPrefix = os.Getenv("INDEX_PREFIX")
	esURL := os.Getenv("ELASTICSEARCH_URL")

	dbConfig := &es.DBConfig{URL: esURL}
	DBClient, err = es.NewDB(dbConfig, log)
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
func processRecords(records []events.DynamoDBEventRecord, db es.DB) error {
	if len(records) == 0 {
		return ErrNoRecords
	}

	docs := []es.Doc{}
	// TODO: we can parallalize this
	for _, record := range records {
		id, err := toId(record.Change.Keys)
		if err != nil {
			return err
		}
		item := map[string]interface{}{}
		for k, v := range record.Change.NewImage {
			if i := toItem(v); i != nil {
				item[santizeKey(k)] = i
			}
		}
		switch events.DynamoDBOperationType(record.EventName) {
		case events.DynamoDBOperationTypeInsert:
			docs = append(docs, es.Doc{Op: es.OpTypeInsert, ID: id, Item: item, Index: indexName(record)})
		case events.DynamoDBOperationTypeModify:
			docs = append(docs, es.Doc{Op: es.OpTypeUpdate, ID: id, Item: item, Index: indexName(record)})
		case events.DynamoDBOperationTypeRemove:
			docs = append(docs, es.Doc{Op: es.OpTypeDelete, ID: id, Item: item, Index: indexName(record)})
		case "":
			continue
		default:
			return fmt.Errorf("Unsupported eventName %s", record.EventName)
		}
	}

	if err := db.WriteDocs(docs); err != nil {
		// print out docs on error
		out, _ := json.Marshal(docs)
		strOut := string(out[:])
		if len(strOut) > 10000 {
			strOut = strOut[:10000]
		}
		fmt.Println(strOut)
		return err
	}

	return nil
}

// toId generates a deterministic Id for each record
func toId(ddbKeys map[string]events.DynamoDBAttributeValue) (string, error) {
	values := []string{}
	for _, key := range ddbKeys {
		item := toItem(key)
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
func toItem(value events.DynamoDBAttributeValue) interface{} {
	switch value.DataType() {
	case events.DataTypeList:
		doc := []interface{}{}
		for _, item := range value.List() {
			if i := toItem(item); i != nil {
				doc = append(doc, i)
			}
		}
		return doc
	case events.DataTypeMap:
		doc := map[string]interface{}{}
		for k, v := range value.Map() {
			if i := toItem(v); i != nil {
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

// indexName computes index name. It does change in a single invocation but for now,
// calculate it for each record since it's simpler
func indexName(record events.DynamoDBEventRecord) string {
	results := indexPattern.FindStringSubmatch(record.EventSourceArn)
	if len(results) == 0 {
		log.ErrorD("table-name-not-found", logger.M{
			"record_id":           record.EventID,
			"record_event_source": record.EventSourceArn,
		})
		return fmt.Sprintf("%s-unknown-table-name", IndexPrefix)
	}

	return fmt.Sprintf("%s%s", IndexPrefix, results[1])
}

// santizeKey makes sure that document keys meet Elasticsearch requirements
func santizeKey(key string) string {
	if _, ok := es.ESReservedFields[key]; ok {
		// add another _ as prefix
		return fmt.Sprintf("_%s", key)
	}
	return key
}
