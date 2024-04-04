package repo

import (
	"context"
	"cookdroogers/internal/requests/sign_contract"
)

//go:generate mockery --name SignContractRequestRepo --with-expecter
type SignContractRequestRepo interface {
	Create(context.Context, *sign_contract.SignContractRequest) error
	Get(ctx context.Context, id uint64) (*sign_contract.SignContractRequest, error)
	Update(context.Context, *sign_contract.SignContractRequest) error
	SetMeta(context.Context, *sign_contract.SignContractRequest) error
}
