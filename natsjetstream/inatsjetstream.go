package natsjetstream

import (
	"time"

	"github.com/nats-io/nats.go"
)

type INatsJetstream interface {
	Connect() error
	Ping()
}

var NatJetstream INatsJetstream

func (n *NatsJetStream) Connect() error {
	nc, err := nats.Connect(n.Config.Host)
	if err != nil {
		return err
	}
	n.NC = nc
	n.Ping()

	return nil
}

func (n *NatsJetStream) Ping() {
	if err := nats.PingInterval(5 * time.Second); err != nil {
		n.lg.Error("ping nats error", "error", err)
		return
	}
	n.lg.Info("ping nats success")
}
