// mongodb/client_test.go
package mongodb

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewClient_InvalidURI(t *testing.T) {
	cfg := Config{
		URI:      "invalid-uri",
		Database: "testdb",
	}
	client, err := NewClient(cfg)
	require.Error(t, err)
	require.Nil(t, client)
}

func TestNewClient_Success(t *testing.T) {
	// This test requires a running MongoDB instance. Set MONGO_URI in environment, e.g.:
	//   export MONGO_URI="mongodb://localhost:27017"
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		t.Skip("MONGO_URI not set, skipping integration test")
	}

	cfg := Config{
		URI:               uri,
		Database:          "testdb",
		ConnectionTimeout: 2 * time.Second,
	}
	client, err := NewClient(cfg)
	require.NoError(t, err)
	require.NotNil(t, client)

	// The database name should match
	require.Equal(t, "testdb", client.DB.Name())

	// Close should not error
	require.NoError(t, client.Close())
}
