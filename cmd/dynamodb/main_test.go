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
		docs    []es.Doc
		err     error
	}{
		{
			request: loadDynamoDBEvent(t),
			docs: []es.Doc{
				es.Doc{
					Op: "insert",
					ID: "data|binary",
					Item: map[string]interface{}{
						"asdf1": []uint8{0x0, 0x1, 0x2a, 0x41},
						"asdf2": [][]uint8{[]uint8{0x0, 0x1, 0x2a, 0x41}, []uint8{0x41, 0x2a, 0x1, 0x0}},
						"key":   "binary", "val": "data"},
				},
				es.Doc{
					Op: "insert",
					ID: "data|binary",
					Item: map[string]interface{}{
						"Binary":         []uint8{0x0, 0x1, 0x2a, 0x41},
						"BinarySet":      [][]uint8{[]uint8{0x0, 0x1, 0x2a, 0x41}, []uint8{0x0, 0x1, 0x2a, 0x41}},
						"Boolean":        true,
						"EmptyStringSet": []string{},
						"FloatNumber":    "123.45",
						"IntegerNumber":  "123",
						"List":           []interface{}{"Cookies", "Coffee", "3.14159"},
						"Map": map[string]interface{}{
							"Age":  "35",
							"Name": "Joe",
							"Workflow": map[string]interface{}{
								"workflowDefinition": map[string]interface{}{
									"stateMachine": map[string]interface{}{
										"NumStates": "1",
										"State1":    "state 1",
									},
								},
							},
						},
						"Workflow": map[string]interface{}{
							"workflowDefinition": map[string]interface{}{},
						},
						"NumberSet": []string{"1234", "567.8"},
						"String":    "Hello",
						"StringSet": []string{"Giraffe", "Zebra"},
						"asdf1":     []uint8{0x0, 0x1, 0x2a, 0x41},
						"asdf2":     [][]uint8{[]uint8{0x0, 0x1, 0x2a, 0x41}, []uint8{0x41, 0x2a, 0x1, 0x0}, []uint8{0x0, 0x1, 0x2a, 0x41}},
						"b2":        []uint8{0xb5, 0xeb, 0x2d}, "key": "binary", "val": "data",
					},
				},
			},
			err: nil,
		},
		{
			request: events.DynamoDBEvent{
				Records: []events.DynamoDBEventRecord{},
			},
			err: ErrNoRecords,
		},
	}

	for _, test := range tests {
		docs, err := processRecords(test.request.Records, &MockDB{})
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.docs, docs)
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
