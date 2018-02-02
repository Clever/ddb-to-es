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

func TestIndexNameParsing(t *testing.T) {
	tests := []struct {
		arn   events.DynamoDBEventRecord
		table string
	}{
		{
			arn: events.DynamoDBEventRecord{
				EventSourceArn: "arn:aws:dynamodb:us-west-2:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899",
			},
			table: "ExampleTableWithStream",
		},
		{
			arn: events.DynamoDBEventRecord{
				EventSourceArn: "arn:aws:dynamodb:us-west-1:589:table/workflow-manager-prod-workflows/stream/2017-12-07T23:33:03.914",
			},
			table: "workflow-manager-prod-workflows",
		},
	}

	for _, test := range tests {
		assert.Equal(t, indexName(test.arn), test.table)
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
