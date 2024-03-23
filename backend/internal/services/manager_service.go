package service

type IManagerService interface {
	GetRandomManagerID() (uint64, error)
}
