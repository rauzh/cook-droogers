package base

import "time"

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

type IRequest interface {
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
	y, m, d := time.Now().UTC().Date()
	req.Date = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
