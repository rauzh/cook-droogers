package models

import "time"

type Statistics struct {
	StatID  uint64    `json:"stat_id"`
	Date    time.Time `json:"date"`
	Streams uint64    `json:"streams"`
	Likes   uint64    `json:"likes"`
	TrackID uint64    `json:"track_id"`
}
