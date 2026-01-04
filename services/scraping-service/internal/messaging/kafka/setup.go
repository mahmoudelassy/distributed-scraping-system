package kafka

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (k *Kafka) SetupTopic(
	topic string,
	partitions int,
	replicationFactor int,
) error {

	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": k.brokers,
	})
	if err != nil {
		return err
	}
	defer admin.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := admin.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{
			{
				Topic:             topic,
				NumPartitions:     partitions,
				ReplicationFactor: replicationFactor,
			},
		},
	)
	if err != nil {
		return err
	}

	for _, res := range results {
		if res.Error.Code() != kafka.ErrNoError &&
			res.Error.Code() != kafka.ErrTopicAlreadyExists {
			return res.Error
		}
	}

	return nil
}
