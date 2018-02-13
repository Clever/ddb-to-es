package es

import (
	"context"
	"fmt"
	"strings"

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

// OpType specifies the kind of operation a Doc represents
type OpType string

var (
	OpTypeInsert OpType = "insert"
	OpTypeUpdate OpType = "modify"
	OpTypeDelete OpType = "delete"
)

// Doc is a representation of an object that can be written via the DB interface
type Doc struct {
	Op    OpType
	ID    string
	Index string
	Item  interface{}
}

// DBConfig specifies how the client should connect to ElasticSearch
type DBConfig struct {
	URL string
}

// DB allows for the writing Doc's to a backend
type DB interface {
	WriteDocs(context.Context, []Doc) error
}

// Elasticsearch exposes functionality to read and write from ElasticSearch
type Elasticsearch struct {
	client *elastic.Client
	config *DBConfig
	lg     logger.KayveeLogger
}

// NewDB creates a new DB instance
func NewDB(config *DBConfig, lg logger.KayveeLogger) (DB, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.URL),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to cluster: %s", err)
	}

	return &Elasticsearch{
		client: client,
		config: config,
		lg:     lg,
	}, nil
}

// WriteDocs implements the writing Doc's to elasticsearch as a batch
func (db *Elasticsearch) WriteDocs(ctx context.Context, docs []Doc) error {
	bulkRequest := db.client.Bulk()

	for _, doc := range docs {
		req := toESRequest(doc)
		// TODO: handle nil (error) cases better. For now let's just keep going
		if req != nil {
			bulkRequest.Add(req)
		}
	}

	if bulkRequest.NumberOfActions() == 0 {
		return nil
	}

	resp, err := bulkRequest.Do(ctx)
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
		if failed.Error != nil {
			db.lg.ErrorD("document-write-failed", logger.M{
				"error-type":   failed.Error.Type,
				"doc-id":       failed.Id,
				"error-reason": failed.Error.Reason,
			})
		} else {
			db.lg.ErrorD("document-write-failed", logger.M{
				"error-type":   "UNKNOWN",
				"doc-id":       failed.Id,
				"error-reason": "UNKNOWN",
			})
		}
	}

	return fmt.Errorf("errors-during-write")
}

// SanitizeKey makes sure that document keys meet Elasticsearch requirements
func SanitizeKey(key string) string {
	if _, ok := ESReservedFields[key]; ok {
		// add another _ as prefix
		return fmt.Sprintf("_%s", key)
	}
	return key
}

func toESRequest(doc Doc) elastic.BulkableRequest {
	// make sure we don't have invalid indexes
	index := strings.ToLower(doc.Index)
	if index == "" {
		index = "unknown"
	}

	switch doc.Op {
	case OpTypeInsert:
		fallthrough
	case OpTypeUpdate:
		return elastic.NewBulkIndexRequest().Index(index).Type("default").Id(doc.ID).Doc(doc.Item)
	case OpTypeDelete:
		return elastic.NewBulkDeleteRequest().Index(index).Type("default").Id(doc.ID)
	default:
		fmt.Printf("INVALID DOC TYPE %s; %s", doc.Op, doc.ID)
		return nil
	}
}
