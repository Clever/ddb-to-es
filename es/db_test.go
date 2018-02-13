package es

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

// this test currently requires a running elasticsaerch
// at localhost:9200
func TestWriteDocs(t *testing.T) {
	tests := []struct {
		docs []Doc
		err  error
	}{
		{
			docs: []Doc{},
			err:  nil,
		},
		{
			docs: []Doc{
				Doc{
					Op: "insert",
					ID: "data|binary",
					Item: map[string]interface{}{
						"hello": "there",
						"boo":   1,
					},
				},
			},
			err: nil,
		},
	}

	db, err := NewDB(&DBConfig{URL: "http://localhost:9200"}, logger.New("test"))
	assert.NoError(t, err)

	for _, test := range tests {
		err = db.WriteDocs(context.Background(), test.docs)
		assert.NoError(t, err)
	}
}

func TestWriteDocsComplexBatch(t *testing.T) {
	docs := &[]Doc{}
	err := json.Unmarshal([]byte(docsJSON), docs)
	assert.NoError(t, err)

	db, err := NewDB(&DBConfig{URL: "http://localhost:9200"}, logger.New("test"))
	assert.NoError(t, err)
	err = db.WriteDocs(context.Background(), *docs)
	assert.NoError(t, err)
}

// a complex document pulled from a test DynamoDB stream
var docsJSON = `[{
            "Op": "modify",
            "ID": "0bf37be0-56bc-4164-8cd1-7de725be01af",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:39:02.015286312Z",
                    "id": "0bf37be0-56bc-4164-8cd1-7de725be01af",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.108042824Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954742",
                "_gsi-ca": "2018-01-31T01:39:02.015286312Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.108042824Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "0bf37be0-56bc-4164-8cd1-7de725be01af"
            }
        },
        {
            "Op": "modify",
            "ID": "7a07b4e9-9b96-4b06-8ddc-9c6eb64006ce",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:41:59.06380604Z",
                    "id": "7a07b4e9-9b96-4b06-8ddc-9c6eb64006ce",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.128468419Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954919",
                "_gsi-ca": "2018-01-31T01:41:59.06380604Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.128468419Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "7a07b4e9-9b96-4b06-8ddc-9c6eb64006ce"
            }
        },
        {
            "Op": "modify",
            "ID": "83a5b3c5-123c-4ded-b065-cb30a2d0004c",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:41:33.822233885Z",
                    "id": "83a5b3c5-123c-4ded-b065-cb30a2d0004c",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.142578974Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954893",
                "_gsi-ca": "2018-01-31T01:41:33.822233885Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.142578974Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "83a5b3c5-123c-4ded-b065-cb30a2d0004c"
            }
        },
        {
            "Op": "modify",
            "ID": "06f9a831-9d1a-45ab-bf7a-1ff52a0240f7",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:42:03.581024938Z",
                    "id": "06f9a831-9d1a-45ab-bf7a-1ff52a0240f7",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.174928506Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954923",
                "_gsi-ca": "2018-01-31T01:42:03.581024938Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.174928506Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "06f9a831-9d1a-45ab-bf7a-1ff52a0240f7"
            }
        },
        {
            "Op": "modify",
            "ID": "b1dd41a6-18d7-4399-b5a0-6a607c60138d",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:41:59.412706666Z",
                    "id": "b1dd41a6-18d7-4399-b5a0-6a607c60138d",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.191842705Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954919",
                "_gsi-ca": "2018-01-31T01:41:59.412706666Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.191842705Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "b1dd41a6-18d7-4399-b5a0-6a607c60138d"
            }
        },
        {
            "Op": "modify",
            "ID": "7a07b4e9-9b96-4b06-8ddc-9c6eb64006ce",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:41:59.06380604Z",
                    "id": "7a07b4e9-9b96-4b06-8ddc-9c6eb64006ce",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.198208298Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954919",
                "_gsi-ca": "2018-01-31T01:41:59.06380604Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.198208298Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "7a07b4e9-9b96-4b06-8ddc-9c6eb64006ce"
            }
        },
        {
            "Op": "modify",
            "ID": "c1b9124f-fa2f-4156-83c7-e5ffa7bd3f27",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:41:59.2254977Z",
                    "id": "c1b9124f-fa2f-4156-83c7-e5ffa7bd3f27",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.226590324Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954919",
                "_gsi-ca": "2018-01-31T01:41:59.2254977Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.226590324Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "c1b9124f-fa2f-4156-83c7-e5ffa7bd3f27"
            }
        },
        {
            "Op": "delete",
            "ID": "cdb6a76d-60cb-4d72-95fc-c944a7433a3a",
            "Item": {}
        },
        {
            "Op": "delete",
            "ID": "42c62471-6c1a-41c0-8e2a-096f21b0e199",
            "Item": {}
        },
        {
            "Op": "delete",
            "ID": "2f4783f0-f037-4f9a-ac03-fccc78f6f603",
            "Item": {}
        },
        {
            "Op": "modify",
            "ID": "06f9a831-9d1a-45ab-bf7a-1ff52a0240f7",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:42:03.581024938Z",
                    "id": "06f9a831-9d1a-45ab-bf7a-1ff52a0240f7",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.25199596Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954923",
                "_gsi-ca": "2018-01-31T01:42:03.581024938Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.25199596Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "06f9a831-9d1a-45ab-bf7a-1ff52a0240f7"
            }
        },
        {
            "Op": "modify",
            "ID": "b1dd41a6-18d7-4399-b5a0-6a607c60138d",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:41:59.412706666Z",
                    "id": "b1dd41a6-18d7-4399-b5a0-6a607c60138d",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.267915868Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954919",
                "_gsi-ca": "2018-01-31T01:41:59.412706666Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.267915868Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "b1dd41a6-18d7-4399-b5a0-6a607c60138d"
            }
        },
        {
            "Op": "modify",
            "ID": "9fa43d01-bd32-4860-a5bb-41a21fe05dff",
            "Item": {
                "Workflow": {
                    "createdAt": "2018-01-31T01:42:17.414758833Z",
                    "id": "9fa43d01-bd32-4860-a5bb-41a21fe05dff",
                    "input": "{}",
                    "jobs": true,
                    "lastUpdated": "2018-01-31T01:42:36.30535706Z",
                    "namespace": "production",
                    "queue": "production",
                    "retries": true,
                    "status": "running",
                    "stoppedAt": "0001-01-01T00:00:00Z",
                    "workflowDefinition": {
                        "createdAt": "2017-12-27T01:52:53.922879594Z",
                        "id": "9294ea0b-2d6f-4b1f-91b4-fab601d86b8b",
                        "manager": "step-functions",
                        "name": "multiverse:master",
                        "stateMachine": {
                            "StartAt": "istrict-sharing",
                            "States": {
                                "heck-preemption": {
                                    "Catch": true,
                                    "Choices": [
                                        {
                                            "And": true,
                                            "BooleanEquals": true,
                                            "Next": "reempt",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        },
                                        {
                                            "And": true,
                                            "BooleanEquals": false,
                                            "Next": "ummaries",
                                            "Or": true,
                                            "Variable": "$.preempted"
                                        }
                                    ],
                                    "Retry": true,
                                    "Type": "Choice"
                                },
                                "inalizer": {
                                    "Catch": true,
                                    "Choices": true,
                                    "End": true,
                                    "HeartbeatSeconds": "30",
                                    "Resource": "inalizer",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "600",
                                    "Type": "Task"
                                },
                                "ocker": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "heck-preemption",
                                    "Resource": "ocker",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "14400",
                                    "Type": "Task"
                                },
                                "reempt": {
                                    "Catch": true,
                                    "Choices": true,
                                    "Retry": true,
                                    "Type": "Succeed"
                                },
                                "lasticsearch-sis": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "inalizer",
                                    "Resource": "lasticsearch-sis",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "vents": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "lasticsearch-sis",
                                    "Resource": "vents",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "is": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "vents",
                                    "Resource": "is",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ummaries": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "iew-differ",
                                    "Resource": "ummaries",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "pp-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ata-gator",
                                    "Resource": "pp-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ata-gator": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ead-links",
                                    "Resource": "ata-gator",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "ead-links": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "etadata",
                                    "Resource": "ead-links",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "istrict-sharing": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "copes",
                                    "Resource": "istrict-sharing",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "etadata": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "ocker",
                                    "Resource": "etadata",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "copes": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "pp-sharing",
                                    "Resource": "copes",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                },
                                "iew-differ": {
                                    "Catch": true,
                                    "Choices": true,
                                    "HeartbeatSeconds": "30",
                                    "Next": "is",
                                    "Resource": "iew-differ",
                                    "Retry": [
                                        {
                                            "ErrorEquals": [
                                                "States.ALL"
                                            ],
                                            "MaxAttempts": "2"
                                        }
                                    ],
                                    "TimeoutSeconds": "7200",
                                    "Type": "Task"
                                }
                            },
                            "TimeoutSeconds": "259200",
                            "Version": "1.0"
                        },
                        "version": "8"
                    }
                },
                "__ttl": "1519954937",
                "_gsi-ca": "2018-01-31T01:42:17.414758833Z",
                "_gsi-lastUpdated": "2018-01-31T01:42:36.30535706Z",
                "_gsi-status": "running",
                "_gsi-wn": "multiverse:master",
                "_gsi-wn-and-resolvedbyuser": "multiverse:master:false",
                "_gsi-wn-and-status": "multiverse:master:running",
                "id": "9fa43d01-bd32-4860-a5bb-41a21fe05dff"
            }
        }
    ]`
