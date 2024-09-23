package base

import (
	baseReqErrors "cookdroogers/internal/requests/base/errors"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

// TEST_HW: mock-less test
type BaseRequestSuite struct {
	suite.Suite
}

func (s *BaseRequestSuite) TestBaseRequest_ValidateOK(t provider.T) {
	t.Title("Validate: OK")
	t.Tags("BaseRequest")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		req := GetBaseRequestObject()

		err := req.Validate("Sign")

		sCtx.Assert().NoError(err)
	})
}

func (s *BaseRequestSuite) TestBaseRequest_ValidateErrInvalidType(t provider.T) {
	t.Title("Validate: OK")
	t.Tags("BaseRequest")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		req := GetBaseRequestObject()

		err := req.Validate("Publish")

		sCtx.Assert().ErrorIs(err, baseReqErrors.ErrInvalidType)
	})
}

func TestBaseRequestSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(BaseRequestSuite))
}
