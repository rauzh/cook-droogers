package models

type Manager struct {
	ManagerID uint64
	UserID    uint64
	Artists   []uint64
}
