package repo

type TransactionManager interface {
	BeginTransaction()
	RollbackTransaction()
	CommitTransaction()
}
