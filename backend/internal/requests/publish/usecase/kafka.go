package usecase

import (
	"context"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	cdtime "cookdroogers/pkg/time"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

const (
	PublishRequestProceedToManager = "publish_request_proceed_to_manager"
	RequestTimeOutExplanation      = "the request is no longer relevant"
)

type PublishReqMessage struct {
	RequestID    uint64             `json:"request_id"`
	Type         base.RequestType   `json:"type"`
	Status       base.RequestStatus `json:"status"`
	Date         time.Time          `json:"date"`
	ApplierID    uint64             `json:"applier_id"`
	ManagerID    uint64             `json:"manager_id"`
	ReleaseID    uint64             `json:"release_id"`
	Grade        int                `json:"grade"`
	ExpectedDate time.Time          `json:"expected_date"`
	Description  string             `json:"description"`
}

func NewPublishRequestProducerMsg(topic string, req *publish.PublishRequest) (*sarama.ProducerMessage, error) {
	msg := NewPublishReqMessage(req)
	msgJson, err := NewMsgValue(msg)
	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msgJson),
	}, nil
}

func (publishUseCase *PublishRequestUseCase) runProceedToManagerConsumer() error {
	if publishUseCase.pbBroker == nil {
		return errors.New("no broker")
	}

	go func() {
		for {
			select {
			case msg := <-publishUseCase.pbBroker.Consumers[PublishRequestProceedToManager].Messages():
				publishUseCase.processProceedToManagerMsg(msg)
			}
		}
	}()

	return nil
}

func (publishUseCase *PublishRequestUseCase) processProceedToManagerMsg(msg *sarama.ConsumerMessage) {

	pubReqMessage := PublishReqMessage{}
	if err := json.Unmarshal(msg.Value, &pubReqMessage); err != nil {
		return
	}

	pubReq := pubReqMessage.ToPublishReq()

	if err := pubReq.Validate(publish.PubReq); err != nil {
		publishUseCase.closeProceedToManagerReq(pubReq, err.Error())
		return
	}

	if msg.Timestamp.Before(cdtime.RelevantPeriod()) {
		publishUseCase.closeProceedToManagerReq(pubReq, RequestTimeOutExplanation)
		return
	}

	if err := publishUseCase.proceedToManager(pubReq); err != nil {

		retryProducerMsg := &sarama.ProducerMessage{
			Topic:     PublishRequestProceedToManager,
			Value:     sarama.StringEncoder(msg.Value),
			Timestamp: msg.Timestamp, // setting OLD timestamp (first one) for TIMEOUT mechanism
		}

		_, _, _ = publishUseCase.pbBroker.SendMessage(retryProducerMsg)
	}
}

func (publishUseCase *PublishRequestUseCase) sendProceedToManagerMSG(pubReq *publish.PublishRequest) error {

	msg, err := NewPublishRequestProducerMsg(PublishRequestProceedToManager, pubReq)
	if err != nil {
		return fmt.Errorf("can't apply publish request: can't proceed to manager with err %w", err)
	}

	_, _, err = publishUseCase.pbBroker.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("can't apply publish request: can't proceed to manager with err %w", err)
	}

	return nil
}

func (publishUseCase *PublishRequestUseCase) closeProceedToManagerReq(pubReq *publish.PublishRequest, explanation string) {

	pubReq.Description = base.DescrDeclinedRequest + ".\n" + explanation
	pubReq.Status = base.ClosedRequest

	if err := publishUseCase.repo.Update(context.Background(), pubReq); err != nil {
		_ = publishUseCase.sendProceedToManagerMSG(pubReq) // resend VALIDATED message
	}
}

func NewMsgValue(msg *PublishReqMessage) ([]byte, error) {
	msgjson, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return msgjson, nil
}

func NewPublishReqMessage(req *publish.PublishRequest) *PublishReqMessage {
	return &PublishReqMessage{
		RequestID:    req.RequestID,
		Type:         req.Type,
		Status:       req.Status,
		Date:         req.Date,
		ApplierID:    req.ApplierID,
		ManagerID:    req.ManagerID,
		ReleaseID:    req.ReleaseID,
		Grade:        req.Grade,
		ExpectedDate: req.ExpectedDate,
		Description:  req.Description,
	}
}

func (msg *PublishReqMessage) ToPublishReq() *publish.PublishRequest {
	return &publish.PublishRequest{
		Request: base.Request{
			RequestID: msg.RequestID,
			Type:      msg.Type,
			Status:    msg.Status,
			Date:      msg.Date,
			ApplierID: msg.ApplierID,
			ManagerID: msg.ManagerID,
		},
		ReleaseID:    msg.ReleaseID,
		Grade:        msg.Grade,
		ExpectedDate: msg.ExpectedDate,
		Description:  msg.Description,
	}
}
