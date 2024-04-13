package broker_dto

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	"encoding/json"
	"github.com/IBM/sarama"
	"time"
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

func NewMsgValue(msg *PublishReqMessage) ([]byte, error) {
	msgjson, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return msgjson, nil
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
