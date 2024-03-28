package repo

//go:generate mockery --name TransactionManager --with-expecter
type TransactionManager interface {
	IsActive(string) bool
	BeginTransaction() (string, error)
	RollbackTransaction(string)
	CommitTransaction(string) error
}
