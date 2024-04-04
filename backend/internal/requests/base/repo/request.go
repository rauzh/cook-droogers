package repo

import (
	"context"
	"cookdroogers/internal/requests/base"
)

//go:generate mockery --name RequestRepo --with-expecter
type RequestRepo interface {
	GetAllByManagerID(context.Context, uint64) ([]base.Request, error)
	GetAllByUserID(context.Context, uint64) ([]base.Request, error)
}
