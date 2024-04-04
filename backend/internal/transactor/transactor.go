package transactor

import "context"

//go:generate mockery --name Transactor --with-expecter
type Transactor interface {
	WithinTransaction(context.Context, func(ctx context.Context) error) error
}
