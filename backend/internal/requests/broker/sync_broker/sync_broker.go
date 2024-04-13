package sync_broker

import (
	"context"
	"cookdroogers/internal/requests/broker"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

const CookDroogersSyncBrokerConsumerGroup = "CDsyncB"

type SyncBroker struct {
	Producer      sarama.SyncProducer
	ConsumerGroup sarama.ConsumerGroup

	ctx           context.Context
	ctxCancelFunc context.CancelFunc
}

func NewSyncBroker(kafkaEndpoints []string, kafkaConfig *sarama.Config) (broker.IBroker, error) {
	producer, err := sarama.NewSyncProducer(kafkaEndpoints, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create sync producer with err %w", err)
	}

	consumerGroup, err := sarama.NewConsumerGroup(kafkaEndpoints, CookDroogersSyncBrokerConsumerGroup, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create consumer group with err %w", err)
	}

	return &SyncBroker{Producer: producer, ConsumerGroup: consumerGroup}, nil
}

func (sb *SyncBroker) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return sb.Producer.SendMessage(msg)
}

func (sb *SyncBroker) Close() error {

	sb.ctxCancelFunc()

	if err := sb.Producer.Close(); err != nil {
		return err
	}
	if err := sb.ConsumerGroup.Close(); err != nil {
		return err
	}

	return nil
}

func (sb *SyncBroker) AddHandler(topics []string, handler broker.IConsumerGroupHandler) {

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := sb.ConsumerGroup.Consume(sb.ctx, topics, handler); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if sb.ctx.Err() != nil {
				return
			}
			handler.Ready()
		}
	}()

	handler.WaitReady()
}
