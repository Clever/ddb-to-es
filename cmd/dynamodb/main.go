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
	"github.com/aws/aws-lambda-go/lambdacontext"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

// AppName is application name provided by catapult
var AppName = os.Getenv("APP_NAME")
var IndexPattern = regexp.MustCompile("arn:aws:dynamodb:.*?:.*?:table/([0-9a-zA-Z_-]+)/.+")

// FailOnError specifies if to only log errors or to also return to lambda
var FailOnError bool
var IndexPrefix string
var DBClient es.DB

// ErrNoRecords is an example error you could generate in handling an event.
var ErrNoRecords = errors.New("no records contained in event")

func main() {
	log := logger.New(AppName)
	var err error
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

// Handler is your Lambda function handler.
// The return signature can be empty, a single error, or a return value (struct or string) and error.
func Handler(ctx context.Context, event events.DynamoDBEvent) error {
	// create a request specific logger, with lambda request id
	ctx = logger.NewContext(ctx, logger.New(AppName))
	if lambdacontext, ok := lambdacontext.FromContext(ctx); ok {
		logger.FromContext(ctx).AddContext("aws_request_id", lambdacontext.AwsRequestID)
	}
	if err := processRecords(ctx, event.Records, DBClient); err != nil {
		if FailOnError {
			return err
		}
		logger.FromContext(ctx).ErrorD("failed-process-records-but-continue", logger.M{
			"message": err.Error(),
		})
	}

	return nil
}

// processRecords converts DynamoDB stream records to es.Doc and writes them to the db
func processRecords(ctx context.Context, records []events.DynamoDBEventRecord, db es.DB) error {
	if len(records) == 0 {
		return ErrNoRecords
	}

	docs := []es.Doc{}
	index := ""
	// TODO: we can parallalize this
	for _, record := range records {
		id, err := toId(record.Change.Keys)
		if err != nil {
			return err
		}
		item := map[string]interface{}{}
		for k, v := range record.Change.NewImage {
			if i := toItem(v); i != nil {
				item[es.SanitizeKey(k)] = i
			}
		}

		// index name does not change in a single invocation
		if index == "" || index == fmt.Sprintf("%s-unknown-table-name", IndexPrefix) {
			index, err = indexName(record)
			if err != nil {
				logger.FromContext(ctx).WarnD("table-name-not-found", logger.M{
					"record_id":           record.EventID,
					"record_event_source": record.EventSourceArn,
				})
				index = fmt.Sprintf("%s-unknown-table-name", IndexPrefix)
			}
		}

		switch events.DynamoDBOperationType(record.EventName) {
		case events.DynamoDBOperationTypeInsert:
			docs = append(docs, es.Doc{Op: es.OpTypeInsert, ID: id, Item: item, Index: index})
		case events.DynamoDBOperationTypeModify:
			docs = append(docs, es.Doc{Op: es.OpTypeUpdate, ID: id, Item: item, Index: index})
		case events.DynamoDBOperationTypeRemove:
			docs = append(docs, es.Doc{Op: es.OpTypeDelete, ID: id, Item: item, Index: index})
		case "":
			continue
		default:
			return fmt.Errorf("Unsupported eventName %s", record.EventName)
		}
	}

	if err := db.WriteDocs(ctx, docs); err != nil {
		// print out docs on error
		out, _ := json.Marshal(docs)
		fmt.Println(string(out[:]))
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
				doc[es.SanitizeKey(k)] = i
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

// indexName computes an doc index name based on the EventSourceArn
// this should be the same for all records in a lambda invocation
func indexName(record events.DynamoDBEventRecord) (string, error) {
	results := IndexPattern.FindStringSubmatch(record.EventSourceArn)
	if len(results) == 0 {
		return "", fmt.Errorf("table-name-not-found")
	}

	return fmt.Sprintf("%s%s", IndexPrefix, results[1]), nil
}
