package service

type ISignContractService interface {
	Apply(uint64, string) error
	Accept(uint64) error
	Decline(uint64) error
}
