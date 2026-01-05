package nats

import (
	"github.com/nats-io/nats.go"
)

type NATSPublisher struct {
	js nats.JetStreamContext
}

func NewNATSPublisher(n *NATS) *NATSPublisher {
	return &NATSPublisher{
		js: n.js,
	}
}

func (p *NATSPublisher) Publish(stream string, message []byte) error {
	_, err := p.js.Publish(stream, message)
	return err
}
