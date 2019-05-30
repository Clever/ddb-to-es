package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Clever/ddb-to-es/es"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type MockDB struct{}

func (db *MockDB) WriteDocs(docs []es.Doc) error {
	return nil
}

func TestProcessRecords(t *testing.T) {
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
		err := processRecords(test.request.Records, &MockDB{})
		assert.Equal(t, test.err, err)
	}
}

func loadDynamoDBEvent(t *testing.T) events.DynamoDBEvent {
	// 1. read JSON from file
	inputJson, err := ioutil.ReadFile("./testdata/dynamodb-event.json")
	if err != nil {
		t.Fatal(err)
	}

	// 2. de-serialize into Go object
	var inputEvent events.DynamoDBEvent
	if err = json.Unmarshal(inputJson, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	return inputEvent
}
