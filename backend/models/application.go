package models

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

type Request struct {
	RequestID uint64
	Type      RequestType
	Status    RequestStatus
	Date      time.Time
	Meta      map[string]string
	ApplierID uint64
	ManagerID uint64
}
