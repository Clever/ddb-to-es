package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

var log = logger.New(os.Getenv("APP_NAME"))

// ErrNoRecords is an example error you could generate in handling an event.
var ErrNoRecords = errors.New("no records contained in event")

type Doc struct {
	Op   string
	Item interface{}
}

// Handler is your Lambda function handler.
// The return signature can be empty, a single error, or a return value (struct or string) and error.
func Handler(ctx context.Context, event events.DynamoDBEvent) error {
	if len(event.Records) == 0 {
		log.Error("no-records-found")
		return ErrNoRecords
	}

	return processRecords(event.Records)
}

func main() {
	lambda.Start(Handler)
}

func processRecords(records []events.DynamoDBEventRecord) error {
	docs := []Doc{}

	// TODO: we can parallalize this
	for _, record := range records {
		item := map[string]interface{}{}
		for k, v := range record.Change.NewImage {
			item[k] = toDoc(v)
		}
		switch events.DynamoDBOperationType(record.EventName) {
		case events.DynamoDBOperationTypeInsert:
			docs = append(docs, Doc{"insert", item})
		case events.DynamoDBOperationTypeModify:
			docs = append(docs, Doc{"modify", item})
		case events.DynamoDBOperationTypeRemove:
			docs = append(docs, Doc{"delete", item})
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

func toDoc(value events.DynamoDBAttributeValue) interface{} {
	switch value.DataType() {
	case events.DataTypeList:
		doc := []interface{}{}
		for _, item := range value.List() {
			doc = append(doc, toDoc(item))
		}
		return doc
	case events.DataTypeMap:
		doc := map[string]interface{}{}
		for k, v := range value.Map() {
			doc[k] = toDoc(v)
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
