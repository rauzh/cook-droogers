package broker_dto

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/sign_contract"
	"encoding/json"
	"github.com/IBM/sarama"
	"time"
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
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msgJson),
	}, nil
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
