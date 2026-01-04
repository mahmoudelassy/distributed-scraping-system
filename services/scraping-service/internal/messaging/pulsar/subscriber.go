package pulsar

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarSubscriber struct {
	consumer pulsar.Consumer
}

func NewPulsarSubscriber(p *Pulsar, topic, subscription string) (*PulsarSubscriber, error) {
	consumer, err := p.client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subscription,
		Type:             pulsar.Shared, // like Kafka consumer group
	})
	if err != nil {
		return nil, err
	}

	return &PulsarSubscriber{
		consumer: consumer,
	}, nil
}

func (s *PulsarSubscriber) Subscribe(topic string, handler func([]byte) error) error {
	go func() {
		for {
			msg, err := s.consumer.Receive(context.Background())
			if err != nil {
				continue
			}

			if handler(msg.Payload()) == nil {
				s.consumer.Ack(msg)
			} else {
				s.consumer.Nack(msg)
			}
		}
	}()

	return nil
}

func (s *PulsarSubscriber) Close() {
	s.consumer.Close()
}
