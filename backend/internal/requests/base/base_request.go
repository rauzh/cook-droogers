package base

import (
	cd_time "cookdroogers/pkg/time"
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

const (
	SignRequest    RequestType = "Sign"
	PublishRequest RequestType = "Publish"
)

const DescrDeclinedRequest = "The request is declined."

type IRequestUseCase interface {
	Apply() error
	Accept() error
	Decline() error
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

func InitDateStatus(req *Request) {
	req.Status = NewRequest
	req.Date = cd_time.GetToday()
}
