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
	RequestID uint64            `json:"request_id"`
	Type      RequestType       `json:"type"`
	Status    RequestStatus     `json:"status"`
	Date      time.Time         `json:"date"`
	Meta      map[string]string `json:"meta"`
	ApplierID uint64            `json:"applier_id"`
	ManagerID uint64            `json:"manager_id"`
}
