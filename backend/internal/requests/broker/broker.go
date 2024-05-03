package broker

import (
	"errors"
	"github.com/IBM/sarama"
)

var (
	ErrNoBroker   error = errors.New("no broker")
	ErrNoConsumer error = errors.New("no consumer")
)

//go:generate mockery --name IBroker --with-expecter
type IBroker interface {
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
	AddHandler(topics []string, handler IConsumerGroupHandler) error
	Close() error
}

type IConsumerGroupHandler interface {
	sarama.ConsumerGroupHandler

	Ready()
	WaitReady()
}
