package nats

import (
	"github.com/nats-io/nats.go"
)

type NATS struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

func NewNATS(url string) (*NATS, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, err
	}

	return &NATS{
		conn: nc,
		js:   js,
	}, nil
}

func (n *NATS) Close() {
	n.conn.Close()
}
