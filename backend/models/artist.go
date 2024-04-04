package models

import "time"

type Artist struct {
	ArtistID     uint64
	UserID       uint64
	Nickname     string
	ContractTerm time.Time
	Activity     bool
	ManagerID    uint64
}
