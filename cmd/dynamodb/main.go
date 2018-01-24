package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Clever/ddb-to-es/es"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

var log = logger.New(os.Getenv("APP_NAME"))
var db *es.DB

// ErrNoRecords is an example error you could generate in handling an event.
var ErrNoRecords = errors.New("no records contained in event")

// Handler is your Lambda function handler.
// The return signature can be empty, a single error, or a return value (struct or string) and error.
func Handler(ctx context.Context, event events.DynamoDBEvent) error {
	if len(event.Records) == 0 {
		log.Error("no-records-found")
		return ErrNoRecords
	}

	return processRecords(event.Records, db)
}

// TODO: tmp fixthis
func initDB() *es.DB {
	dbConfig := &es.DBConfig{URL: os.Getenv("ELASTICSEARCH_URL")}
	db, err := es.NewDB(dbConfig, log)
	if err != nil {
		log.ErrorD("elasticsearch-connect-error", logger.M{
			"message": err.Error(),
			"url":     dbConfig.URL,
		})
		os.Exit(1)
	}
	return db
}

func main() {
	db = initDB() // TODO: initiate the DB as an interface and outside of main()
	lambda.Start(Handler)
}

func processRecords(records []events.DynamoDBEventRecord, db *es.DB) error {
	docs := []es.Doc{}

	// TODO: we can parallalize this
	for _, record := range records {
		id, err := toId(record.Change.Keys)
		if err != nil {
			return err
		}
		item := map[string]interface{}{}
		for k, v := range record.Change.NewImage {
			item[santizeKey(k)] = toItem(v)
		}
		switch events.DynamoDBOperationType(record.EventName) {
		case events.DynamoDBOperationTypeInsert:
			docs = append(docs, es.Doc{Op: "insert", Id: id, Item: item})
		case events.DynamoDBOperationTypeModify:
			docs = append(docs, es.Doc{Op: "modify", Id: id, Item: item})
		case events.DynamoDBOperationTypeRemove:
			docs = append(docs, es.Doc{Op: "delete", Id: id, Item: item})
		case "":
			continue
		default:
			return fmt.Errorf("Unsupported eventName %s", record.EventName)
		}
	}

	out, err := json.MarshalIndent(docs, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(out[:]))

	return nil
}

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

func toItem(value events.DynamoDBAttributeValue) interface{} {
	switch value.DataType() {
	case events.DataTypeList:
		doc := []interface{}{}
		for _, item := range value.List() {
			doc = append(doc, toItem(item))
		}
		return doc
	case events.DataTypeMap:
		doc := map[string]interface{}{}
		for k, v := range value.Map() {
			doc[santizeKey(k)] = toItem(v)
		}
		return doc
	case events.DataTypeNull:
		return value.IsNull()
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

func santizeKey(key string) string {
	if _, ok := es.ESReservedFields[key]; ok {
		// add another _ as prefix
		return fmt.Sprintf("_%s", key)
	}
	return key
}
