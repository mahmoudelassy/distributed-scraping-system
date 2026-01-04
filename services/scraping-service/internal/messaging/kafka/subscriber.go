package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaSubscriber struct {
	consumer *kafka.Consumer
}

func NewKafkaSubscriber(k *Kafka, groupID string) (*KafkaSubscriber, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  k.brokers,
		"group.id":           groupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaSubscriber{consumer: c}, nil
}

func (s *KafkaSubscriber) Subscribe(topic string, handler func([]byte) error) error {
	if err := s.consumer.Subscribe(topic, nil); err != nil {
		return err
	}

	go func() {
		for {
			msg, err := s.consumer.ReadMessage(-1)
			if err != nil {
				continue
			}

			if err := handler(msg.Value); err != nil {
				// no commit → message will be retried
				continue
			}

			_, _ = s.consumer.CommitMessage(msg)
		}
	}()

	return nil
}

func (s *KafkaSubscriber) Close() error {
	return s.consumer.Close()
}
