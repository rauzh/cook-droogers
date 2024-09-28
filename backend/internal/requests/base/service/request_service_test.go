package service

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/base/repo/mocks"
	"errors"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	reqRepo *mocks.RequestRepo
	logger  *slog.Logger
}

// TEST_HW: default test
type RequestServiceSuite struct {
	suite.Suite
}

func _newMockSignReqDepFields(t provider.T) *_depFields {

	mockReqRepo := mocks.NewRequestRepo(t)

	f := &_depFields{
		reqRepo: mockReqRepo,
		logger:  slog.Default(),
	}

	return f
}

func (s *RequestServiceSuite) TestRequestService_GetAllByManagerIDOK(t provider.T) {
	t.Title("GetAllByManagerID: OK")
	t.Tags("RequestService")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)
		df.reqRepo.EXPECT().GetAllByManagerID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
			Return([]base.Request{
				*base.GetBaseRequestObject(),
			}, nil).Once()

		reqService := NewRequestService(df.reqRepo, df.logger)

		reqs, err := reqService.GetAllByManagerID(uint64(7))

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(reqs)
		sCtx.Assert().Equal([]base.Request{
			*base.GetBaseRequestObject(),
		}, reqs)
	})
}

func (s *RequestServiceSuite) TestRequestService_GetAllByManagerIDDbErr(t provider.T) {
	t.Title("GetAllByManagerID: DBerr")
	t.Tags("RequestService")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		someDBerr := errors.New("some db err")

		df := _newMockSignReqDepFields(t)
		df.reqRepo.EXPECT().GetAllByManagerID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
			Return(nil, someDBerr).Once()

		reqService := NewRequestService(df.reqRepo, df.logger)

		reqs, err := reqService.GetAllByManagerID(uint64(7))

		sCtx.Assert().Nil(reqs)
		sCtx.Assert().ErrorIs(err, DBerr)
	})
}

func (s *RequestServiceSuite) TestRequestService_GetAllByUserIDOK(t provider.T) {
	t.Title("GetAllByUserID: OK")
	t.Tags("RequestService")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)
		df.reqRepo.EXPECT().GetAllByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
			Return([]base.Request{
				*base.GetBaseRequestObject(),
			}, nil).Once()

		reqService := NewRequestService(df.reqRepo, df.logger)

		reqs, err := reqService.GetAllByUserID(uint64(7))

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(reqs)
		sCtx.Assert().Equal([]base.Request{
			*base.GetBaseRequestObject(),
		}, reqs)
	})
}

func (s *RequestServiceSuite) TestRequestService_GetAllByUserIDDbErr(t provider.T) {
	t.Title("GetAllByUserID: DBerr")
	t.Tags("RequestService")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		someDBerr := errors.New("some db err")

		df := _newMockSignReqDepFields(t)
		df.reqRepo.EXPECT().GetAllByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
			Return(nil, someDBerr).Once()

		reqService := NewRequestService(df.reqRepo, df.logger)

		reqs, err := reqService.GetAllByUserID(uint64(7))

		sCtx.Assert().Nil(reqs)
		sCtx.Assert().ErrorIs(err, DBerr)
	})
}

func (s *RequestServiceSuite) TestRequestService_GetByIDOK(t provider.T) {
	t.Title("GetByID: OK")
	t.Tags("RequestService")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)
		df.reqRepo.EXPECT().GetByID(mock.AnythingOfType("context.backgroundCtx"), uint64(1)).
			Return(base.GetBaseRequestObject(), nil).Once()

		reqService := NewRequestService(df.reqRepo, df.logger)

		req, err := reqService.GetByID(uint64(1))

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(req)
		sCtx.Assert().Equal(base.GetBaseRequestObject(), req)
	})
}

func (s *RequestServiceSuite) TestRequestService_GetByIDDbErr(t provider.T) {
	t.Title("GetByID: DBerr")
	t.Tags("RequestService")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		someDBerr := errors.New("some db err")

		df := _newMockSignReqDepFields(t)
		df.reqRepo.EXPECT().GetByID(mock.AnythingOfType("context.backgroundCtx"), uint64(1)).
			Return(nil, someDBerr).Once()

		reqService := NewRequestService(df.reqRepo, df.logger)

		reqs, err := reqService.GetByID(uint64(1))

		sCtx.Assert().Nil(reqs)
		sCtx.Assert().ErrorIs(err, DBerr)
	})
}

func TestRequestServiceSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(RequestServiceSuite))
}
