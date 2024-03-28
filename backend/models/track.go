package models

type Track struct {
	TrackID  uint64   `json:"track_id"`
	Title    string   `json:"title"`
	Duration uint64   `json:"duration"`
	Genre    string   `json:"genre"`
	Type     string   `json:"type"`
	Artists  []uint64 `json:"artists"`
}
