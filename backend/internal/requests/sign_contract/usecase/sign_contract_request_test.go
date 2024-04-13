package usecase

import (
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/internal/requests/base"
	base_errors "cookdroogers/internal/requests/base/errors"
	broker_mocks "cookdroogers/internal/requests/broker/mocks"
	"cookdroogers/internal/requests/sign_contract"
	sctErrors "cookdroogers/internal/requests/sign_contract/errors"
	signReqRepoMocks "cookdroogers/internal/requests/sign_contract/repo/mocks"
	transacMock "cookdroogers/internal/transactor/mocks"
	cdtime "cookdroogers/pkg/time"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
)

type _depFields struct {
	artistRepo *mocks.ArtistRepo
	userRepo   *mocks.UserRepo

	transactor *transacMock.Transactor
	scBroker   *broker_mocks.IBroker

	signReqRepo *signReqRepoMocks.SignContractRequestRepo
}

var dberr = errors.New("db err")

func _newMockSignReqDepFields(t *testing.T) *_depFields {

	transactionMock := transacMock.NewTransactor(t)
	mockBroker := broker_mocks.NewIBroker(t)

	mockArtRepo := mocks.NewArtistRepo(t)
	mockUserRepo := mocks.NewUserRepo(t)

	mockSignReqRepo := signReqRepoMocks.NewSignContractRequestRepo(t)

	f := &_depFields{
		artistRepo:  mockArtRepo,
		userRepo:    mockUserRepo,
		transactor:  transactionMock,
		scBroker:    mockBroker,
		signReqRepo: mockSignReqRepo,
	}

	return f
}

func TestPublishRequestUseCase_Decline(t *testing.T) {

	type args struct {
		signReq *sign_contract.SignContractRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields)
	}{
		{
			name: "OK",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					Nickname:    "pink floyd",
					Description: "",
				},
			},
			out: nil,
			dependencies: func(df *_depFields) {

				df.signReqRepo.EXPECT().Update(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.ClosedRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					Nickname:    "pink floyd",
					Description: base.DescrDeclinedRequest,
				}).Return(nil).Once()
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "InvalidNickName",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					Nickname:    "",
					Description: "",
				},
			},
			out: sctErrors.ErrNickname,
			dependencies: func(df *_depFields) {
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockSignReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			signReqUseCase, err := NewSignContractRequestUseCase(f.userRepo, f.artistRepo, f.transactor, f.scBroker, f.signReqRepo)

			// act
			err = signReqUseCase.Decline(tt.in.signReq)

			// assert
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
			if tt.assert != nil {
				tt.assert(t, f)
			}
		})
	}
}

func TestPublishRequestUseCase_Accept(t *testing.T) {

	type args struct {
		signReq *sign_contract.SignContractRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields)
	}{
		{
			name: "OK",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					Nickname:    "skibidi",
					Description: "",
				},
			},
			out: nil,
			dependencies: func(df *_depFields) {
				df.transactor.EXPECT().WithinTransaction(mock.AnythingOfType("context.backgroundCtx"),
					mock.Anything).Return(nil).Once()
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "InvalidNickName",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					Nickname:    "",
					Description: "",
				},
			},
			out: sctErrors.ErrNickname,
			dependencies: func(df *_depFields) {
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "AlreadyClosed",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.ClosedRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					Nickname:    "skibidi",
					Description: "",
				},
			},
			out: base_errors.ErrAlreadyClosed,
			dependencies: func(df *_depFields) {
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockSignReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			signReqUseCase, err := NewSignContractRequestUseCase(f.userRepo, f.artistRepo, f.transactor, f.scBroker, f.signReqRepo)

			// act
			err = signReqUseCase.Accept(tt.in.signReq)

			// assert
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
			if tt.assert != nil {
				tt.assert(t, f)
			}
		})
	}
}

func TestPublishRequestUseCase_Apply(t *testing.T) {

	type args struct {
		signReq *sign_contract.SignContractRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields)
	}{
		{
			name: "OK",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    "",
						ApplierID: 12,
						ManagerID: 0,
					},
					Nickname:    "skibidi",
					Description: "",
				},
			},
			out: nil,
			dependencies: func(df *_depFields) {

				df.scBroker.EXPECT().SendMessage(mock.Anything).Return(0, 0, nil)

				df.signReqRepo.EXPECT().Create(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.NewRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 0,
					},
					Nickname:    "skibidi",
					Description: "",
				}).Return(nil).Once()
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "CantCreate",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    "",
						ApplierID: 12,
						ManagerID: 0,
					},
					Nickname:    "skibidi",
					Description: "",
				},
			},
			out: dberr,
			dependencies: func(df *_depFields) {

				df.signReqRepo.EXPECT().Create(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.NewRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 0,
					},
					Nickname:    "skibidi",
					Description: "",
				}).Return(dberr).Once()
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockSignReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			signReqUseCase, err := NewSignContractRequestUseCase(f.userRepo, f.artistRepo, f.transactor, f.scBroker, f.signReqRepo)

			// act
			err = signReqUseCase.Apply(tt.in.signReq)

			// assert
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
			if tt.assert != nil {
				tt.assert(t, f)
			}
		})
	}
}
