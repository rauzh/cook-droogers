package models

import "time"

type Statistics struct {
	StatID  uint64
	Date    time.Time
	Streams uint64
	Likes   uint64
	TrackID uint64
}
