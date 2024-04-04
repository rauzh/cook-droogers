package sign_contract

import (
	"cookdroogers/internal/requests/base"
)

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
	EmptyID        = 0
)
