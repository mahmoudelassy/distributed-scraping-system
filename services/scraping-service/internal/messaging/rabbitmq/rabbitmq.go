package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) Publish(exchange, routingKey string, message []byte) error {
	return r.ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         message,
		},
	)
}

func (r *RabbitMQ) Subscribe(queue string, handler func([]byte) error) error {
	_, err := r.ch.QueueDeclare(
		queue, true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	msgs, err := r.ch.Consume(
		queue, "", false, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := handler(d.Body); err == nil {
				d.Ack(false)
			}
		}
	}()

	return nil
}

func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}
