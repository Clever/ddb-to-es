package es

import (
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
					Id: "data|binary",
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
		err = db.WriteDocs(test.docs)
		assert.NoError(t, err)
	}
}
