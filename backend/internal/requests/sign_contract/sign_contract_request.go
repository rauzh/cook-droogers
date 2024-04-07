package sign_contract

import (
	"cookdroogers/internal/requests/base"
	sctErrors "cookdroogers/internal/requests/sign_contract/errors"
)

const SignRequest base.RequestType = "Sign"

type SignContractRequest struct {
	base.Request
	Nickname    string
	Description string
}

const (
	YearsContract  = 1
	MonthsContract = 0
	DaysContract   = 0
	MaxNicknameLen = 128
)

func (scReq *SignContractRequest) Validate(reqType base.RequestType) error {

	if err := scReq.Request.Validate(reqType); err != nil {
		return nil
	}

	if scReq.Nickname == "" || len(scReq.Nickname) > MaxNicknameLen {
		return sctErrors.ErrNickname
	}

	return nil
}

func (scReq *SignContractRequest) GetType() base.RequestType {
	return scReq.Type
}
