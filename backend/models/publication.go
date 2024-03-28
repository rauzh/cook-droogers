package models

import "time"

type Publication struct {
	PublicationID uint64    `json:"publication_id"`
	Date          time.Time `json:"date"`
	ReleaseID     uint64    `json:"release_id"`
	ManagerID     uint64    `json:"manager_id"`
}
