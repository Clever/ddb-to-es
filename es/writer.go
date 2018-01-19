package es

import (
	"context"
	"fmt"
	"time"

	elastic "gopkg.in/olivere/elastic.v6"
)

// StartWorker initializers a BulkProcessor for the DB to dispatch any CRUD events to
func (es *DB) StartWorker() error {
	bulkProcessor, err := es.client.BulkProcessor().
		Name("elasticsearch-worker").
		BulkActions(1000).              // commit if # requests >= 1000
		BulkSize(2 << 20).              // commit if size of requests >= 2 MB
		FlushInterval(5 * time.Second). // commit every 5s
		Before(es.beforeCommit).
		After(es.afterCommit).
		Stats(true).
		Do(context.Background())
	if err != nil {
		return err
	}
	es.bulkProcessor = bulkProcessor

	return nil
}

// FlushWorker manually flushes the elasticsearch worker's underlying bulkProcessor
func (es *DB) FlushWorker() error {
	return es.bulkProcessor.Flush()
}

// StopWorker stops the elasticsearch worker's underlying bulkProcessor
func (es *DB) StopWorker() error {
	return es.bulkProcessor.Stop()
}

// TODO do something useful here
func (es *DB) beforeCommit(id int64, requests []elastic.BulkableRequest) {
	// fmt.Println(id, "committing: ", len(requests))
}

// TODO lets add error handling, throttling, etc
func (es *DB) afterCommit(
	id int64,
	requests []elastic.BulkableRequest,
	response *elastic.BulkResponse,
	err error,
) {
	if err != nil {
		// lg something here
		// we will prefer to fail fast
		es.errC <- err
	}
	if response != nil {
		failedReq := response.Failed()
		for _, failure := range failedReq {
			// 404s have no `Error` field, so ignore them
			if failure.Status != 404 {
				fmt.Println(failure.Error.Reason)
			}
		}
	}
}

// LogWorkerStats logs the stats of the ElasticSearch client's bulkProcessor operations
func (es *DB) LogWorkerStats() {
	// TODO use kayvee
	stats := es.bulkProcessor.Stats()

	fmt.Printf("Number of times flush has been invoked: %d\n", stats.Flushed)
	fmt.Printf("Number of times workers committed reqs: %d\n", stats.Committed)
	fmt.Printf("Number of requests indexed            : %d\n", stats.Indexed)
	fmt.Printf("Number of requests reported as created: %d\n", stats.Created)
	fmt.Printf("Number of requests reported as updated: %d\n", stats.Updated)
	fmt.Printf("Number of requests reported as success: %d\n", stats.Succeeded)
	fmt.Printf("Number of requests reported as failed : %d\n", stats.Failed)

	for i, w := range stats.Workers {
		fmt.Printf("Worker %d: Number of requests queued: %d\n", i, w.Queued)
		fmt.Printf("Worker %d: Last response time       : %v\n", i, w.LastDuration)
	}
}
