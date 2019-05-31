package es

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	elastic "gopkg.in/olivere/elastic.v6"
)

// Delete the indices if they already exist, and create new indices for testing.
func setupIndices(t *testing.T, client *elastic.Client, indices []string) {
	deleteIndices(client, indices)
	createIndices(t, client, indices)
}

// Create new indices for testing.
func createIndices(t *testing.T, client *elastic.Client, indices []string) {
	for _, index := range indices {
		createIndex(t, client, index)
	}
}

// Delete test indices.
func deleteIndices(client *elastic.Client, indices []string) {
	for _, index := range indices {
		deleteIndex(client, index)
	}
}

// Create an Elasticsearch index, and assert that it is created successfully.
func createIndex(t *testing.T, client *elastic.Client, index string) {
	resp, err := elastic.NewIndicesCreateService(client).Index(index).Do(context.TODO())
	assert.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.Acknowledged)
}

// Delete the given index if it exists.
// Ignore errors so this is acts as a no-op when the index doesn't exist.
// The API accepts multiple indices in a single request, but the API call errors if any of the
// provided indices are not found. For that reason, we only pass a single index per API call.
func deleteIndex(client *elastic.Client, index string) {
	_, _ = elastic.NewIndicesDeleteService(client).Index([]string{index}).Do(context.TODO())
}

// Assert the existence of a document in an index.
func assertIndexHasDoc(t *testing.T, client *elastic.Client, index, id string) {
	resp, err := elastic.NewGetService(client).Index(index).Id(id).Do(context.TODO())
	assert.NoError(t, err, fmt.Sprintf("%s %s", index, id))
	require.NotNil(t, resp)
	assert.True(t, resp.Found)
}
