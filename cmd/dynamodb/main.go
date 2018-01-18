package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

var log = logger.New(os.Getenv("APP_NAME"))

func convert(from interface{}, to interface{}) {
	bs, _ := json.Marshal(from)
	json.Unmarshal(bs, to)
}

func toMap(i interface{}) logger.M {
	var m logger.M
	convert(i, &m)
	return m
}

// ErrNoRecords is an example error you could generate in handling an event.
var ErrNoRecords = errors.New("no records contained in event")

// Handler is your Lambda function handler.
// The return signature can be empty, a single error, or a return value (struct or string) and error.
func Handler(ctx context.Context, event events.DynamoDBEvent) error {
	log.InfoD("event", toMap(event))
	if len(event.Records) == 0 {
		return ErrNoRecords
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
