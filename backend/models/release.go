package models

import "time"

type ReleaseStatus string

const (
	UnpublishedRelease ReleaseStatus = "Unpublished"
	PublishedRelease   ReleaseStatus = "Published"
)

type Release struct {
	ReleaseID    uint64
	Title        string
	Status       ReleaseStatus
	DateCreation time.Time
	Tracks       []uint64
	ArtistID     uint64
}
