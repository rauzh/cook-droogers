package models

import "time"

type Publication struct {
	PublicationID uint64
	Date          time.Time
	ReleaseID     uint64
	ManagerID     uint64
}
