package service

import "time"

type ReleaseStatus string

type Release struct {
	ReleaseID    uint64
	Title        string
	Status       ReleaseStatus
	DateCreation time.Time
	Tracks       []uint64
	ArtistID     uint64
}
