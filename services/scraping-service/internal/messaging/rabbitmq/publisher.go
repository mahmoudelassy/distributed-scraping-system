package rabbitmq

import (
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TopicConfig struct {
	Exchange   string
	RoutingKey string
}

var ErrUnknownTopic = errors.New("unknown topic")

type RabbitMQPublisher struct {
	ch       *amqp.Channel
	topicMap map[string]TopicConfig
}

func NewRabbitMQPublisher(r *RabbitMQ, topicMap map[string]TopicConfig) *RabbitMQPublisher {
	return &RabbitMQPublisher{
		ch:       r.ch,
		topicMap: topicMap,
	}
}

func (p *RabbitMQPublisher) Publish(topic string, message []byte) error {
	cfg, ok := p.topicMap[topic]
	if !ok {
		return fmt.Errorf("%w: %s", ErrUnknownTopic, topic)
	}

	return p.ch.Publish(
		cfg.Exchange,
		cfg.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         message,
		},
	)
}
