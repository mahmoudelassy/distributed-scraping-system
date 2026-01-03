package rabbitmq

func (r *RabbitMQ) SetupQueue(exchange string, queue string, routingKey string) error {

	if err := r.ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	q, err := r.ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return r.ch.QueueBind(
		q.Name,
		routingKey,
		exchange,
		false,
		nil,
	)
}
