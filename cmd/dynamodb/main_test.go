package main

import (
	"context"
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
			request: events.DynamoDBEvent{
				Records: []events.DynamoDBEventRecord{{}},
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
		err := Handler(context.Background(), test.request)
		assert.Equal(t, test.err, err)
	}
}
