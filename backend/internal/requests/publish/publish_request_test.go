package publish

import (
	"cookdroogers/internal/requests/base"
	pubReqErrors "cookdroogers/internal/requests/publish/errors"
	cdtime "cookdroogers/pkg/time"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type PublishRequestSuite struct {
	suite.Suite
}

func (s *PublishRequestSuite) TestPublishRequest_ValidateOK(t provider.T) {
	t.Title("Validate: OK")
	t.Tags("PublishRequest")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		req := &PublishRequest{
			Request: base.Request{
				RequestID: 123,
				Type:      PubReq,
				ManagerID: 8,
				ApplierID: 88,
				Status:    base.ProcessingRequest,
				Date:      cdtime.GetToday(),
			},
			ReleaseID:    100,
			Description:  "Test description",
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Grade:        5,
		}

		err := req.Validate(PubReq)

		sCtx.Assert().NoError(err)
	})
}

func (s *PublishRequestSuite) TestBaseRequest_ValidateErrInvalidReleaseID(t provider.T) {
	t.Title("Validate: OK")
	t.Tags("PublishRequest")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		req := &PublishRequest{
			Request: base.Request{
				RequestID: 123,
				Type:      PubReq,
				ManagerID: 8,
				ApplierID: 88,
				Status:    base.ProcessingRequest,
				Date:      cdtime.GetToday(),
			},
			ReleaseID:    0,
			Description:  "Test description",
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Grade:        5,
		}

		err := req.Validate(PubReq)

		sCtx.Assert().ErrorIs(err, pubReqErrors.ErrNoReleaseID)
	})
}

func TestPublishRequestSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(PublishRequestSuite))
}
