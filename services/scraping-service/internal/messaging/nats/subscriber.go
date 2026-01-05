package nats

import (
	"time"

	"github.com/nats-io/nats.go"
)

type NATSSubscriber struct {
	js  nats.JetStreamContext
	sub *nats.Subscription
}

func NewNATSSubscriber(n *NATS) *NATSSubscriber {
	return &NATSSubscriber{
		js: n.js,
	}
}

func (s *NATSSubscriber) Subscribe(stream, consumer string, handler func([]byte) error) error {
	sub, err := s.js.Subscribe(
		stream,
		func(msg *nats.Msg) {
			go func(m *nats.Msg) {
				if handler(m.Data) == nil {
					m.Ack() // success
				} else {
					m.Nak() // retry later
				}
			}(msg)
		},
		nats.Durable(consumer),
		nats.ManualAck(),
		nats.AckWait(30*time.Second),
	)
	if err != nil {
		return err
	}

	s.sub = sub
	return nil
}

func (s *NATSSubscriber) Close() {
	if s.sub != nil {
		s.sub.Unsubscribe()
	}
}
