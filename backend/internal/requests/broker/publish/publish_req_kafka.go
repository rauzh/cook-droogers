package publish

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/broker"
	"cookdroogers/internal/requests/broker/broker_dto"
	criteria "cookdroogers/internal/requests/criteria_controller"
	"cookdroogers/internal/requests/publish"
	publishReqRepo "cookdroogers/internal/requests/publish/repo"
	cdtime "cookdroogers/pkg/time"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

const (
	PublishRequestProceedToManager = "publish_request_proceed_to_manager"
	RequestTimeOutExplanation      = "the request is no longer relevant"
)

type PublishProceedToManagerConsumerHandler struct {
	broker broker.IBroker

	publishRepo publishReqRepo.PublishRequestRepo
	artistRepo  repo.ArtistRepo

	criterias criteria.ICriteriaCollection

	ready chan bool
}

func InitPublishProceedToManagerConsumerHandler(
	broker broker.IBroker,
	publishRepo publishReqRepo.PublishRequestRepo,
	artistRepo repo.ArtistRepo,
	criterias criteria.ICriteriaCollection,
) broker.IConsumerGroupHandler {
	return &PublishProceedToManagerConsumerHandler{
		broker:      broker,
		publishRepo: publishRepo,
		artistRepo:  artistRepo,
		criterias:   criterias,
		ready:       make(chan bool),
	}
}

func (handler *PublishProceedToManagerConsumerHandler) Ready() {
	handler.ready = make(chan bool)
	handler.ready <- true
}

func (handler *PublishProceedToManagerConsumerHandler) WaitReady() {
	<-handler.ready
}

func (handler *PublishProceedToManagerConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	close(handler.ready)
	return nil
}

func (handler *PublishProceedToManagerConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (handler *PublishProceedToManagerConsumerHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for {
		select {
		case message := <-claim.Messages():

			if message.Topic == PublishRequestProceedToManager {
				err := handler.processProceedToManagerMsg(message)
				if err != nil {
					// don't mark message as consumed and return
				}
			}

			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

func (handler *PublishProceedToManagerConsumerHandler) processProceedToManagerMsg(msg *sarama.ConsumerMessage) error {
	var err error

	pubReqMessage := broker_dto.PublishReqMessage{}
	if err := json.Unmarshal(msg.Value, &pubReqMessage); err != nil {
		return err
	}

	pubReq := pubReqMessage.ToPublishReq()

	if err := pubReq.Validate(publish.PubReq); err != nil {
		return handler.closeProceedToManagerReq(pubReq, err.Error())
	}

	if msg.Timestamp.Before(cdtime.RelevantPeriod()) {
		return handler.closeProceedToManagerReq(pubReq, RequestTimeOutExplanation)
	}

	if err := handler.proceedToManager(pubReq); err != nil {

		retryProducerMsg := &sarama.ProducerMessage{
			Topic:     PublishRequestProceedToManager,
			Value:     sarama.StringEncoder(msg.Value),
			Timestamp: msg.Timestamp, // setting OLD timestamp (first one) for TIMEOUT mechanism
		}

		_, _, err = handler.broker.SendMessage(retryProducerMsg)
	}

	return err
}

func (handler *PublishProceedToManagerConsumerHandler) sendProceedToManagerMSG(pubReq *publish.PublishRequest) error {

	msg, err := broker_dto.NewPublishRequestProducerMsg(PublishRequestProceedToManager, pubReq)
	if err != nil {
		return fmt.Errorf("can't apply publish request: can't proceed to manager with err %w", err)
	}

	_, _, err = handler.broker.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("can't apply publish request: can't proceed to manager with err %w", err)
	}

	return nil
}

func (handler *PublishProceedToManagerConsumerHandler) closeProceedToManagerReq(
	pubReq *publish.PublishRequest, explanation string) error {

	pubReq.Description = base.DescrDeclinedRequest + ".\n" + explanation
	pubReq.Status = base.ClosedRequest

	if err := handler.publishRepo.Update(context.Background(), pubReq); err != nil {
		return handler.sendProceedToManagerMSG(pubReq) // if db can't update, resend msg
	}

	return nil
}
