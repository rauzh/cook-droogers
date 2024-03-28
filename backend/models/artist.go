package models

import "time"

type Artist struct {
	UserID       uint64    `json:"user_id"`
	Nickname     string    `json:"nickname"`
	ContractTerm time.Time `json:"contract_term"`
	Activity     bool      `json:"activity"`
	ManagerID    uint64    `json:"manager_id"`
}
