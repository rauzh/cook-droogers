package models

import "time"

type ApplicationStatus string

const (
	NewApplication        ApplicationStatus = "New"
	ProcessingApplication ApplicationStatus = "Processing"
	OnApprovalApplication ApplicationStatus = "On approval"
	ClosedApplication     ApplicationStatus = "Closed"
)

type ApplicationType string

const (
	SignApplication    ApplicationType = "Sign"
	PublishApplication ApplicationType = "Publish"
)

type Application struct {
	ApplicationID uint64
	Type          ApplicationType
	Status        ApplicationStatus
	Date          time.Time
	Meta          map[string]string
	ApplierID     uint64
	ManagerID     uint64
}
