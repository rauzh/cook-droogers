package base

import (
	baseReqErrors "cookdroogers/internal/requests/base/errors"
	cdtime "cookdroogers/pkg/time"
	"time"
)

type RequestStatus string

const (
	NewRequest        RequestStatus = "New"
	ProcessingRequest RequestStatus = "Processing"
	OnApprovalRequest RequestStatus = "On approval"
	ClosedRequest     RequestStatus = "Closed"
)

type RequestType string

const DescrDeclinedRequest = "The request is declined."
const EmptyID = 0

type IRequestUseCase interface {
	Apply(request IRequest) error
	Accept(request IRequest) error
	Decline(request IRequest) error
}

type IRequest interface {
	Validate(RequestType) error
	GetType() RequestType
}

type Request struct {
	RequestID uint64
	Type      RequestType
	Status    RequestStatus
	Date      time.Time
	ApplierID uint64
	ManagerID uint64
}

func (req *Request) Validate(requestType RequestType) error {
	if req.ApplierID == EmptyID {
		return baseReqErrors.ErrNoApplierID
	}
	if req.Type != requestType {
		return baseReqErrors.ErrInvalidType
	}
	if req.Status == ClosedRequest {
		return baseReqErrors.ErrAlreadyClosed
	}

	return nil
}

func (req *Request) GetType() RequestType {
	return req.Type
}

func InitDateStatus(req *Request) {
	req.Status = NewRequest
	req.Date = cdtime.GetToday()
}
