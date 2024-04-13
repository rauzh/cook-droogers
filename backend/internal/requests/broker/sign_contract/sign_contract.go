package sign_contract

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/broker"
	"cookdroogers/internal/requests/broker/broker_dto"
	"cookdroogers/internal/requests/sign_contract"
	signRepo "cookdroogers/internal/requests/sign_contract/repo"
	cdtime "cookdroogers/pkg/time"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

const (
	SignRequestProceedToManager = "sign_request_proceed_to_manager"
	RequestTimeOutExplanation   = "the request is no longer relevant"
)

type SignContractProceedToManagerHandler struct {
	broker broker.IBroker

	signReqRepo signRepo.SignContractRequestRepo
	mngRepo     repo.ManagerRepo

	ready chan bool
}

func InitSignContractProceedToManagerHandler(
	broker broker.IBroker,
	signReqRepo signRepo.SignContractRequestRepo,
	mngRepo repo.ManagerRepo,
) broker.IConsumerGroupHandler {
	return &SignContractProceedToManagerHandler{
		broker:      broker,
		signReqRepo: signReqRepo,
		mngRepo:     mngRepo,
	}
}

func (handler *SignContractProceedToManagerHandler) Ready() {
	handler.ready = make(chan bool)
}

func (handler *SignContractProceedToManagerHandler) WaitReady() {
	<-handler.ready
}

func (handler *SignContractProceedToManagerHandler) Setup(session sarama.ConsumerGroupSession) error {
	close(handler.ready)
	return nil
}

func (handler *SignContractProceedToManagerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (handler *SignContractProceedToManagerHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for {
		select {
		case message := <-claim.Messages():

			if message.Topic == SignRequestProceedToManager {
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

func (handler *SignContractProceedToManagerHandler) processProceedToManagerMsg(msg *sarama.ConsumerMessage) error {
	var err error

	signContractReqMsg := broker_dto.SignContractReqMessage{}
	if err := json.Unmarshal(msg.Value, &signContractReqMsg); err != nil {
		return err
	}

	signReq := signContractReqMsg.ToSignContractReq()

	if err := signReq.Validate(sign_contract.SignRequest); err != nil {
		return handler.closeProceedToManagerReq(signReq, err.Error())
	}

	if msg.Timestamp.Before(cdtime.RelevantPeriod()) {
		return handler.closeProceedToManagerReq(signReq, RequestTimeOutExplanation)
	}

	if err := handler.proceedToManager(signReq); err != nil {

		retryProducerMsg := &sarama.ProducerMessage{
			Topic:     SignRequestProceedToManager,
			Value:     sarama.StringEncoder(msg.Value),
			Timestamp: msg.Timestamp, // setting OLD timestamp (first one) for TIMEOUT mechanism
		}

		_, _, err = handler.broker.SendMessage(retryProducerMsg)
	}

	return err
}

func (handler *SignContractProceedToManagerHandler) sendProceedToManagerMSG(signReq *sign_contract.SignContractRequest) error {

	msg, err := broker_dto.NewSignRequestProducerMsg(SignRequestProceedToManager, signReq)
	if err != nil {
		return fmt.Errorf("can't apply sign contract request: can't proceed to manager with err %w", err)
	}

	_, _, err = handler.broker.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("can't apply sign contract request: can't proceed to manager with err %w", err)
	}

	return nil
}

func (handler *SignContractProceedToManagerHandler) closeProceedToManagerReq(
	signReq *sign_contract.SignContractRequest, explanation string) error {

	signReq.Description = base.DescrDeclinedRequest + ".\n" + explanation
	signReq.Status = base.ClosedRequest

	if err := handler.signReqRepo.Update(context.Background(), signReq); err != nil {
		return handler.sendProceedToManagerMSG(signReq) // if db can't update, resend msg
	}

	return nil
}
