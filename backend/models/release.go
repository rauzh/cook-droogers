package models

import "time"

type ReleaseStatus string

const (
	UnpublishedRelease ReleaseStatus = "Unpublished"
	PublishedRelease   ReleaseStatus = "Published"
)

type Release struct {
	ReleaseID    uint64        `json:"release_id"`
	Title        string        `json:"title"`
	Status       ReleaseStatus `json:"status"`
	DateCreation time.Time     `json:"date_creation"`
	Tracks       []uint64      `json:"tracks"`
	ArtistID     uint64        `json:"artist_id"`
}
