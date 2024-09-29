package sign_contract

import (
	"cookdroogers/internal/requests/base"
	sctErrors "cookdroogers/internal/requests/sign_contract/errors"
	cdtime "cookdroogers/pkg/time"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type SignContractRequestSuite struct {
	suite.Suite
}

func (s *SignContractRequestSuite) TestSignContractRequest_ValidateOK(t provider.T) {
	t.Title("Validate: OK")
	t.Tags("SignContractRequest")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		req := &SignContractRequest{
			Request: base.Request{
				RequestID: 123,
				Type:      SignRequest,
				ManagerID: 8,
				ApplierID: 88,
				Status:    base.ProcessingRequest,
				Date:      cdtime.GetToday(),
			},
			Description: "Test description",
			Nickname:    "leclerc",
		}

		err := req.Validate(SignRequest)

		sCtx.Assert().NoError(err)
	})
}

func (s *SignContractRequestSuite) TestBaseRequest_ValidateErrInvalidNickname(t provider.T) {
	t.Title("Validate: OK")
	t.Tags("SignContractRequest")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		req := &SignContractRequest{
			Request: base.Request{
				RequestID: 123,
				Type:      SignRequest,
				ManagerID: 8,
				ApplierID: 88,
				Status:    base.ProcessingRequest,
				Date:      cdtime.GetToday(),
			},
			Description: "Test description",
			Nickname:    "",
		}

		err := req.Validate(SignRequest)

		sCtx.Assert().ErrorIs(err, sctErrors.ErrNickname)
	})
}

func TestSignContractRequestSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(SignContractRequestSuite))
}
