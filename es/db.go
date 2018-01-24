package es

import (
	"context"
	"fmt"

	"gopkg.in/Clever/kayvee-go.v6/logger"
	elastic "gopkg.in/olivere/elastic.v6"
)

var ESReservedFields = map[string]bool{
	"uid":         true,
	"_id":         true,
	"_type":       true,
	"_source":     true,
	"_all":        true,
	"_parent":     true,
	"_fieldnames": true,
	"_routing":    true,
	"_index":      true,
	"_size":       true,
	"_timestamp":  true,
	"_ttl":        true,
}

// DBConfig specifies how the client should connect to ElasticSearch
type DBConfig struct {
	URL string
}

// DB exposes functionality to read and write from ElasticSearch
type DB struct {
	client *elastic.Client
	config *DBConfig
	lg     logger.KayveeLogger
}

// NewDB creates a new DB instance
func NewDB(config *DBConfig, lg logger.KayveeLogger) (*DB, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.URL),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to cluster: %s", err)
	}

	return &DB{
		client: client,
		config: config,
		lg:     lg,
	}, nil
}

type Doc struct {
	Op   string
	Id   string
	Item interface{}
}

func (db *DB) WriteDocs(docs []Doc) error {
	bulkRequest := db.client.Bulk()

	for _, doc := range docs {
		bulkRequest.Add(toESRequest(doc))
	}

	if bulkRequest.NumberOfActions() == 0 {
		return nil
	}

	resp, err := bulkRequest.Do(context.Background())
	if err != nil {
		db.lg.ErrorD("write-failed", logger.M{
			"error-type":   "UNKNOWN",
			"error-reason": err.Error(),
		})
		return err
	}

	if !resp.Errors {
		return nil
	}

	// log all errors
	for _, failed := range resp.Failed() {
		db.lg.ErrorD("document-write-failed", logger.M{
			"error-type":   failed.Error.Type,
			"doc-id":       failed.Id,
			"error-reason": failed.Error.Reason,
		})
	}

	return fmt.Errorf("errors-during-write")
}

func toESRequest(doc Doc) elastic.BulkableRequest {
	switch doc.Op {
	case "insert":
		fallthrough
	case "update":
		return elastic.NewBulkIndexRequest().Index("testing").Type("testing").Id(doc.Id).Doc(doc.Item)
	case "delete":
		return elastic.NewBulkDeleteRequest().Index("testing").Type("testing").Id(doc.Id)
	default:
		return nil
	}
}