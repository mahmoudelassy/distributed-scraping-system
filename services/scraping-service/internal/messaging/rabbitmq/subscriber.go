package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQSubscriber struct {
	ch *amqp.Channel
}

func NewRabbitMQSubscriber(r *RabbitMQ) *RabbitMQSubscriber {
	return &RabbitMQSubscriber{
		ch: r.ch,
	}
}

func (s *RabbitMQSubscriber) Subscribe(queue string, handler func([]byte) error) error {
	msgs, err := s.ch.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := handler(d.Body); err != nil {
				d.Nack(false, true)
				continue
			}
			d.Ack(false)
		}
	}()

	return nil
}
