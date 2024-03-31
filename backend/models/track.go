package models

type Track struct {
	TrackID  uint64
	Title    string
	Duration uint64
	Genre    string
	Type     string
	Artists  []uint64
}
