package usecase

import (
	"context"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/sign_contract"
	cdtime "cookdroogers/pkg/time"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

const (
	SignRequestProceedToManager = "sign_request_proceed_to_manager"
	RequestTimeOutExplanation   = "the request is no longer relevant"
)

type SignContractReqMessage struct {
	RequestID   uint64             `json:"request_id"`
	Type        base.RequestType   `json:"type"`
	Status      base.RequestStatus `json:"status"`
	Date        time.Time          `json:"date"`
	ApplierID   uint64             `json:"applier_id"`
	ManagerID   uint64             `json:"manager_id"`
	Nickname    string             `json:"nickname"`
	Description string             `json:"description"`
}

func NewSignRequestProducerMsg(topic string, req *sign_contract.SignContractRequest) (*sarama.ProducerMessage, error) {
	msg := NewSignContractReqMessage(req)
	msgJson, err := NewMsgValue(msg)
	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msgJson),
	}, nil
}

func (sctUseCase *SignContractRequestUseCase) runProceedToManagerConsumer() error {
	if sctUseCase.scBroker == nil {
		return errors.New("no broker")
	}

	go func() {
		for {
			select {
			case msg := <-sctUseCase.scBroker.Consumers[SignRequestProceedToManager].Messages():
				sctUseCase.processProceedToManagerMsg(msg)
			}
		}
	}()

	return nil
}

func (sctUseCase *SignContractRequestUseCase) processProceedToManagerMsg(msg *sarama.ConsumerMessage) {

	signContractReqMsg := SignContractReqMessage{}
	if err := json.Unmarshal(msg.Value, &signContractReqMsg); err != nil {
		return // LOG
	}

	signReq := signContractReqMsg.ToSignContractReq()

	if err := signReq.Validate(sign_contract.SignRequest); err != nil {
		sctUseCase.closeProceedToManagerReq(signReq, err.Error())
		return
	}

	if msg.Timestamp.Before(cdtime.RelevantPeriod()) {
		sctUseCase.closeProceedToManagerReq(signReq, RequestTimeOutExplanation)
		return
	}

	if err := sctUseCase.proceedToManager(signReq); err != nil {

		retryProducerMsg := &sarama.ProducerMessage{
			Topic:     SignRequestProceedToManager,
			Value:     sarama.StringEncoder(msg.Value),
			Timestamp: msg.Timestamp, // реализация протухания начиная с **первой отправки**
		}

		_, _, _ = sctUseCase.scBroker.SendMessage(retryProducerMsg)
	}
}

func (sctUseCase *SignContractRequestUseCase) closeProceedToManagerReq(signReq *sign_contract.SignContractRequest, explanation string) {

	signReq.Description = base.DescrDeclinedRequest + ".\n" + explanation
	signReq.Status = base.ClosedRequest

	if err := sctUseCase.repo.Update(context.Background(), signReq); err != nil {
		_ = sctUseCase.sendProceedToManagerMSG(signReq) // resend VALIDATED message
	}
}

func (sctUseCase *SignContractRequestUseCase) sendProceedToManagerMSG(signReq *sign_contract.SignContractRequest) error {

	msg, err := NewSignRequestProducerMsg(SignRequestProceedToManager, signReq)
	if err != nil {
		return fmt.Errorf("can't apply sign contract request: can't proceed to manager with err %w", err)
	}

	_, _, err = sctUseCase.scBroker.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("can't apply sign contract request: can't proceed to manager with err %w", err)
	}

	return nil
}

func NewSignContractReqMessage(req *sign_contract.SignContractRequest) *SignContractReqMessage {
	return &SignContractReqMessage{
		RequestID:   req.RequestID,
		Type:        req.Type,
		Status:      req.Status,
		Date:        req.Date,
		ApplierID:   req.ApplierID,
		ManagerID:   req.ManagerID,
		Nickname:    req.Nickname,
		Description: req.Description,
	}
}

func (msg *SignContractReqMessage) ToSignContractReq() *sign_contract.SignContractRequest {
	return &sign_contract.SignContractRequest{
		Request: base.Request{
			RequestID: msg.RequestID,
			Type:      msg.Type,
			Status:    msg.Status,
			Date:      msg.Date,
			ApplierID: msg.ApplierID,
			ManagerID: msg.ManagerID,
		},
		Nickname:    msg.Nickname,
		Description: msg.Description,
	}
}

func NewMsgValue(msg *SignContractReqMessage) ([]byte, error) {
	msgjson, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return msgjson, nil
}
