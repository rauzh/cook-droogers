package broker

import (
	"errors"
	"fmt"
	"github.com/IBM/sarama"
)

//go:generate mockery --name IBroker --with-expecter
type IBroker interface {
	SignConsumerToTopic(topic string) error
	GetConsumerByTopic(topic string) sarama.PartitionConsumer
	Close() error
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
}

var (
	ErrNoBroker   error = errors.New("no broker")
	ErrNoConsumer error = errors.New("no consumer")
)

type SyncBroker struct {
	Producer  sarama.SyncProducer
	Master    sarama.Consumer
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

	return &SyncBroker{Producer: producer, Master: master}, nil
}

func (sb *SyncBroker) SignConsumerToTopic(topic string) error {
	consumer, err := sb.Master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return fmt.Errorf("can't sign consumer with err %w", err)
	}
	sb.Consumers[topic] = consumer
	return nil
}

func (sb *SyncBroker) Close() error {
	if err := sb.Producer.Close(); err != nil {
		return err
	}
	if err := sb.Master.Close(); err != nil {
		return err
	}

	return nil
}

func (sb *SyncBroker) GetConsumerByTopic(topic string) sarama.PartitionConsumer {
	return sb.Consumers[topic]
}

func (sb *SyncBroker) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return sb.Producer.SendMessage(msg)
}
