package models

type Manager struct {
	UserID  uint64   `json:"user_id"`
	Artists []uint64 `json:"artists"`
}
