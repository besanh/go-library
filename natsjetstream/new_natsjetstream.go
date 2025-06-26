package natsjetstream

import (
	"github.com/besanh/go-library/logger"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type (
	NatsJetStream struct {
		NC     *nats.Conn
		Client *jetstream.JetStream
		Config Config
		lg     logger.ILogger
	}

	Config struct {
		Host string
	}
)

func NewNatsJetstream(config Config) INatsJetstream {
	nat := &NatsJetStream{
		Config: config,
		lg:     logger.Default(),
	}

	return nat
}
