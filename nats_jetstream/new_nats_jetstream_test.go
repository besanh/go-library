package natsjetstream

import (
	"testing"

	natsserver "github.com/nats-io/nats-server/v2/test"
	"github.com/stretchr/testify/require"
)

func TestNewNatsJetstream_Defaults(t *testing.T) {
	cfg := Config{
		Host: "nats://localhost:4222",
	}
	nj := NewNatsJetstream(cfg)

	impl, ok := nj.(*NatsJetStream)
	require.True(t, ok, "expected *NatsJetStream, got %T", nj)
	require.Equal(t, cfg, impl.Config)

	require.Nil(t, impl.NC, "expected NC to be nil before initialization")
	require.Nil(t, impl.Client, "expected Client to be nil before initialization")

	require.NotNil(t, impl.lg, "expected logger not to be non-nil")
}

func TestNatsJetstream_SetupConnection(t *testing.T) {
	s := natsserver.RunDefaultServer()
	defer s.Shutdown()

	cfg := Config{Host: s.ClientURL()}
	nj := NewNatsJetstream(cfg).(*NatsJetStream)

	err := nj.Connect()
	require.NoError(t, err, "connect should succeed against test server")

	require.NotNil(t, nj.NC)
	require.True(t, nj.NC.IsConnected(), "NC.IsConnected() should be true")

	js, err := nj.NC.JetStream()
	require.NoError(t, err, "JetStream() should succeed")
	require.NotNil(t, js)

	nj.NC.Close()
}
