package transactor

import "context"

//go:generate mockery --name Transactor --with-expecter
type Transactor interface {
	IsActive(string) bool
	BeginTransaction() (string, error)
	RollbackTransaction(string)
	CommitTransaction(string) error

	WithinTransaction(context.Context, func(ctx context.Context) error) error
}
