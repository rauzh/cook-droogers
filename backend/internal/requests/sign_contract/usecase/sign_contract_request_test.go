package usecase

import (
	"context"
	repo_mocks "cookdroogers/internal/repo/mocks"
	"cookdroogers/internal/requests/base"
	base_errors "cookdroogers/internal/requests/base/errors"
	broker_mocks "cookdroogers/internal/requests/broker/mocks"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/data_builder"
	sctErrors "cookdroogers/internal/requests/sign_contract/errors"
	"cookdroogers/internal/requests/sign_contract/repo/mocks"
	transacMock "cookdroogers/internal/transactor/mocks"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	artistRepo *repo_mocks.ArtistRepo
	userRepo   *repo_mocks.UserRepo

	transactor *transacMock.Transactor
	scBroker   *broker_mocks.IBroker

	signReqRepo *mocks.SignContractRequestRepo
	logger      *slog.Logger
}

type SignContractRequestUseCaseSuite struct {
	suite.Suite
}

func _newMockSignReqDepFields(t provider.T) *_depFields {
	mockSignReqRepo := mocks.NewSignContractRequestRepo(t)

	mockArtistRepo := repo_mocks.NewArtistRepo(t)
	mockUserRepo := repo_mocks.NewUserRepo(t)

	mockTransactor := transacMock.NewTransactor(t)

	mockBroker := broker_mocks.NewIBroker(t)

	f := &_depFields{
		artistRepo:  mockArtistRepo,
		userRepo:    mockUserRepo,
		transactor:  mockTransactor,
		scBroker:    mockBroker,
		signReqRepo: mockSignReqRepo,
		logger:      slog.Default(),
	}

	return f
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_DeclineOK(t provider.T) {
	t.Title("Decline: OK")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)

		req := data_builder.NewSignContractRequestBuilder().
			WithRequestID(1).
			WithNickname("pink floyd").
			WithStatus(base.OnApprovalRequest).
			Build()

		df.signReqRepo.EXPECT().Update(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
			Request: base.Request{
				RequestID: 1,
				Type:      sign_contract.SignRequest,
				Status:    base.ClosedRequest,
				ApplierID: req.ApplierID,
				ManagerID: req.ManagerID,
				Date:      req.Date,
			},
			Nickname:    req.Nickname,
			Description: base.DescrDeclinedRequest,
		}).Return(nil).Once()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		err := signReqUseCase.Decline(context.Background(), req)

		sCtx.Assert().NoError(err)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_DeclineInvalidNickName(t provider.T) {
	t.Title("Decline: InvalidNickName")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockSignReqDepFields(t)

		req := data_builder.NewSignContractRequestBuilder().
			WithRequestID(1).
			WithNickname("").
			WithStatus(base.OnApprovalRequest).
			Build()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		err := signReqUseCase.Decline(context.Background(), req)

		sCtx.Assert().ErrorIs(err, sctErrors.ErrNickname)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_ApplyOK(t provider.T) {
	t.Title("Apply: OK")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)

		req := data_builder.NewSignContractRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithDescription("").
			WithNickname("skibidi").
			Build()

		df.scBroker.EXPECT().SendMessage(mock.Anything).Return(0, 0, nil)

		df.signReqRepo.EXPECT().Create(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
			Request: base.Request{
				RequestID: 1,
				Type:      sign_contract.SignRequest,
				Status:    base.NewRequest,
				Date:      cdtime.GetToday(),
				ApplierID: 12,
				ManagerID: 8,
			},
			Nickname:    "skibidi",
			Description: "",
		}).Return(nil).Once()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		err := signReqUseCase.Apply(context.Background(), req)

		sCtx.Assert().NoError(err)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_ApplyFailure(t provider.T) {
	t.Title("Apply: Failure")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)

		req := data_builder.NewSignContractRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithDescription("").
			WithNickname("skibidi").
			Build()

		df.signReqRepo.EXPECT().Create(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
			Request: base.Request{
				RequestID: 1,
				Type:      sign_contract.SignRequest,
				Status:    base.NewRequest,
				Date:      cdtime.GetToday(),
				ApplierID: 12,
				ManagerID: 8,
			},
			Nickname:    "skibidi",
			Description: "",
		}).Return(sql.ErrConnDone).Once()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		err := signReqUseCase.Apply(context.Background(), req)

		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_AcceptOK(t provider.T) {
	t.Title("Accept: OK")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockSignReqDepFields(t)

		req := data_builder.NewSignContractRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.OnApprovalRequest).
			WithNickname("skibidi").
			Build()

		df.transactor.EXPECT().WithinTransaction(mock.AnythingOfType("context.backgroundCtx"), mock.Anything).Return(nil).Once()

		//df.userRepo.EXPECT().UpdateType(mock.AnythingOfType("context.backgroundCtx"), req.ApplierID, models.ArtistUser).Return(nil).Once()

		//df.artistRepo.EXPECT().Create(mock.AnythingOfType("context.backgroundCtx"), &models.Artist{
		//	UserID:       req.ApplierID,
		//	Nickname:     req.Nickname,
		//	ContractTerm: cdtime.GetEndOfContract(),
		//	Activity:     true,
		//	ManagerID:    req.ManagerID,
		//}).Return(nil).Once()

		//df.signReqRepo.EXPECT().Update(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
		//	Request: base.Request{
		//		RequestID: 1,
		//		Type:      sign_contract.SignRequest,
		//		Status:    base.ClosedRequest,
		//		ApplierID: req.ApplierID,
		//		ManagerID: req.ManagerID,
		//		Date:      req.Date,
		//	},
		//	Nickname:    req.Nickname,
		//	Description: req.Description,
		//}).Return(nil).Once()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		err := signReqUseCase.Accept(context.Background(), req)

		sCtx.Assert().NoError(err)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_AcceptInvalidNickName(t provider.T) {
	t.Title("Accept: InvalidNickName")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockSignReqDepFields(t)

		req := data_builder.NewSignContractRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.OnApprovalRequest).
			WithNickname("").
			Build()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		err := signReqUseCase.Accept(context.Background(), req)

		sCtx.Assert().ErrorIs(err, sctErrors.ErrNickname)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_AcceptAlreadyClosed(t provider.T) {
	t.Title("Accept: AlreadyClosed")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockSignReqDepFields(t)

		req := data_builder.NewSignContractRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.ClosedRequest).
			WithNickname("skibidi").
			Build()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		err := signReqUseCase.Accept(context.Background(), req)

		sCtx.Assert().ErrorIs(err, base_errors.ErrAlreadyClosed)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)
		df.signReqRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(1)).
			Return(data_builder.NewSignContractRequestBuilder().WithRequestID(1).Build(), nil).Once()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		signreq, err := signReqUseCase.(*SignContractRequestUseCase).Get(context.Background(), uint64(1))

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(signreq)
		sCtx.Assert().Equal(data_builder.NewSignContractRequestBuilder().WithRequestID(1).Build(), signreq)
	})
}

func (s *SignContractRequestUseCaseSuite) TestSignContractRequestUseCase_GetDbErr(t provider.T) {
	t.Title("Get: DBerr")
	t.Tags("SignContractRequestUseCase")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		df := _newMockSignReqDepFields(t)
		df.signReqRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(1)).
			Return(nil, sql.ErrConnDone).Once()

		signReqUseCase, _ := NewSignContractRequestUseCase(df.userRepo, df.artistRepo, df.transactor, df.scBroker, df.signReqRepo, df.logger)

		signreq, err := signReqUseCase.(*SignContractRequestUseCase).Get(context.Background(), uint64(1))

		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
		sCtx.Assert().Nil(signreq)
	})
}

func TestSignContractRequestUseCaseSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(SignContractRequestUseCaseSuite))
}
