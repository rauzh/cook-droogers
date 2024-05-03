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
	Producer       sarama.SyncProducer
	ConsumerGroups map[string]sarama.ConsumerGroup

	endpoints []string
	config    *sarama.Config

	ctx           context.Context
	ctxCancelFunc context.CancelFunc
}

func NewSyncBroker(kafkaEndpoints []string, kafkaConfig *sarama.Config) (broker.IBroker, error) {
	producer, err := sarama.NewSyncProducer(kafkaEndpoints, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create sync producer with err %w", err)
	}

	return &SyncBroker{
		Producer:       producer,
		ConsumerGroups: make(map[string]sarama.ConsumerGroup),
		config:         kafkaConfig,
		endpoints:      kafkaEndpoints,
	}, nil
}

func (sb *SyncBroker) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return sb.Producer.SendMessage(msg)
}

func (sb *SyncBroker) Close() error {

	sb.ctxCancelFunc()

	if err := sb.Producer.Close(); err != nil {
		return err
	}

	for _, cg := range sb.ConsumerGroups {
		if err := cg.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (sb *SyncBroker) AddHandler(topics []string, handler broker.IConsumerGroupHandler) error {
	sb.ctx = context.Background()

	consumerGroup, err := sarama.NewConsumerGroup(sb.endpoints,
		fmt.Sprintf("%s_%d", CookDroogersSyncBrokerConsumerGroup, len(sb.ConsumerGroups)), sb.config)
	if err != nil {
		return fmt.Errorf("can't create consumer group with err %w", err)
	}

	for _, topic := range topics {
		sb.ConsumerGroups[topic] = consumerGroup
	}

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := consumerGroup.Consume(sb.ctx, topics, handler); err != nil {
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
	return nil
}
