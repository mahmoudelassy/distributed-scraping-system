package pulsar

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarPublisher struct {
	producer pulsar.Producer
}

func NewPulsarPublisher(p *Pulsar, topic string) (*PulsarPublisher, error) {
	prod, err := p.client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return nil, err
	}

	return &PulsarPublisher{
		producer: prod,
	}, nil
}

func (p *PulsarPublisher) Publish(topic string, message []byte) error {
	_, err := p.producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: message,
	})
	return err
}

func (p *PulsarPublisher) Close() {
	p.producer.Close()
}
