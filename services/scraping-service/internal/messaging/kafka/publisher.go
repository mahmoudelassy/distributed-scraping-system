package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaPublisher struct {
	producer *kafka.Producer
}

func NewKafkaPublisher(k *Kafka) (*KafkaPublisher, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  k.brokers,
		"acks":               "all",
		"enable.idempotence": true,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaPublisher{producer: p}, nil
}

func (p *KafkaPublisher) Publish(topic string, message []byte) error {
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
	}, nil)
}

func (p *KafkaPublisher) RunDeliveryLoop() {
	go func() {
		for e := range p.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					// log / retry / DLQ
				}
			}
		}
	}()
}
