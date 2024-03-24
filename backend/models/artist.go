package models

import "time"

type Artist struct {
	UserID       uint64
	Nickname     string
	ContractTerm time.Time
	Activity     bool
	ManagerID    uint64
}
