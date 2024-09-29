package data_builder

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/sign_contract"
	cdtime "cookdroogers/pkg/time"
	"time"
)

type SignContractRequestBuilder struct {
	SignContractRequest *sign_contract.SignContractRequest
}

func NewSignContractRequestBuilder() *SignContractRequestBuilder {
	return &SignContractRequestBuilder{
		SignContractRequest: &sign_contract.SignContractRequest{
			Request: base.Request{
				RequestID: 123,
				Type:      sign_contract.SignRequest,
				ManagerID: 8,
				ApplierID: 88,
				Status:    base.ProcessingRequest,
				Date:      cdtime.GetToday(),
			},
			Description: "Test description",
			Nickname:    "leclerc",
		},
	}
}

func (b *SignContractRequestBuilder) WithRequestID(Requestid uint64) *SignContractRequestBuilder {
	b.SignContractRequest.RequestID = Requestid
	return b
}

func (b *SignContractRequestBuilder) WithApplierID(id uint64) *SignContractRequestBuilder {
	b.SignContractRequest.ApplierID = id
	return b
}

func (b *SignContractRequestBuilder) WithManagerID(id uint64) *SignContractRequestBuilder {
	b.SignContractRequest.ManagerID = id
	return b
}

func (b *SignContractRequestBuilder) WithType(t base.RequestType) *SignContractRequestBuilder {
	b.SignContractRequest.Type = t
	return b
}

func (b *SignContractRequestBuilder) WithStatus(s base.RequestStatus) *SignContractRequestBuilder {
	b.SignContractRequest.Status = s
	return b
}

func (b *SignContractRequestBuilder) WithDate(date time.Time) *SignContractRequestBuilder {
	b.SignContractRequest.Date = date
	return b
}

func (b *SignContractRequestBuilder) WithDescription(descr string) *SignContractRequestBuilder {
	b.SignContractRequest.Description = descr
	return b
}

func (b *SignContractRequestBuilder) WithNickname(nick string) *SignContractRequestBuilder {
	b.SignContractRequest.Nickname = nick
	return b
}

func (b *SignContractRequestBuilder) Build() *sign_contract.SignContractRequest {
	return b.SignContractRequest
}
