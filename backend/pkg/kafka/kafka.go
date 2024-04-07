package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

type SyncBroker struct {
	producer  sarama.SyncProducer
	master    sarama.Consumer
	Consumers map[string]sarama.PartitionConsumer // topic : consumer
}

func NewSyncBroker(kafkaEndpoints []string, kafkaConfig *sarama.Config) (*SyncBroker, error) {
	producer, err := sarama.NewSyncProducer(kafkaEndpoints, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create sync producer with err %w", err)
	}

	master, err := sarama.NewConsumer(kafkaEndpoints, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create consumer with err %w", err)
	}

	return &SyncBroker{producer: producer, master: master}, nil
}

func (sb *SyncBroker) SignConsumerToTopic(topic string) error {
	consumer, err := sb.master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return fmt.Errorf("can't sign consumer with err %w", err)
	}
	sb.Consumers[topic] = consumer
	return nil
}

func (sb *SyncBroker) Close() error {
	if err := sb.producer.Close(); err != nil {
		return err
	}
	if err := sb.master.Close(); err != nil {
		return err
	}

	return nil
}

func (sb *SyncBroker) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return sb.producer.SendMessage(msg)
}
