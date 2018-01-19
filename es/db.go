package es

import (
	"fmt"

	"gopkg.in/Clever/kayvee-go.v6/logger"
	elastic "gopkg.in/olivere/elastic.v6"
)

// DBConfig specifies how the client should connect to ElasticSearch
type DBConfig struct {
	URL string
}

// DB exposes functionality to read and write from ElasticSearch
type DB struct {
	bulkProcessor *elastic.BulkProcessor
	client        *elastic.Client
	config        *DBConfig
	errC          chan error
	lg            logger.KayveeLogger
}

// NewDB creates a new DB instance
func NewDB(config *DBConfig, lg logger.KayveeLogger, errC chan error) (*DB, error) {
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
		errC:   errC,
		lg:     lg,
	}, nil
}

type Doc struct {
	Op   string
	Id   string
	Item interface{}
}

func ToESRequest(op string, doc Doc) elastic.BulkableRequest {
	switch op {
	case "insert":
		fallthrough
	case "update":
		return elastic.NewBulkIndexRequest().Index("testing").Type("testing").Id(doc.Id).Doc(doc.Item)
	case "delete":
		return elastic.NewBulkDeleteRequest().Index("testing").Type("testing").Id(doc.Id)
	}
	return nil
}
