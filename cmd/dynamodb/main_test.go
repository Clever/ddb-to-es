package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.DynamoDBEvent
		err     error
	}{
		{
			request: loadDynamoDBEvent(t),
			err:     nil,
		},
		{
			request: events.DynamoDBEvent{
				Records: []events.DynamoDBEventRecord{},
			},
			err: ErrNoRecords,
		},
	}

	for _, test := range tests {
		err := Handler(context.Background(), test.request)
		assert.Equal(t, test.err, err)
	}
}

func loadDynamoDBEvent(t *testing.T) events.DynamoDBEvent {
	// 1. read JSON from file
	inputJson := readJsonFromFile(t, "./testdata/dynamodb-event.json")

	// 2. de-serialize into Go object
	var inputEvent events.DynamoDBEvent
	if err := json.Unmarshal(inputJson, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	return inputEvent
}

func readJsonFromFile(t *testing.T, inputFile string) []byte {
	inputJson, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJson
}
