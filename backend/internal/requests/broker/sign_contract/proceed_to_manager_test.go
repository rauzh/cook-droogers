package sign_contract

import (
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/internal/requests/base"
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
	artistRepo  *mocks.ArtistRepo
	managerRepo *mocks.ManagerRepo
	userRepo    *mocks.UserRepo

	transactor *transacMock.Transactor
	scBroker   *broker_mocks.IBroker

	signReqRepo *signReqRepoMocks.SignContractRequestRepo
}

var dberr = errors.New("db err")

func _newMockSignReqDepFields(t *testing.T) *_depFields {

	transactionMock := transacMock.NewTransactor(t)
	mockBroker := broker_mocks.NewIBroker(t)

	mockArtRepo := mocks.NewArtistRepo(t)
	mockManagerRepo := mocks.NewManagerRepo(t)
	mockUserRepo := mocks.NewUserRepo(t)

	mockSignReqRepo := signReqRepoMocks.NewSignContractRequestRepo(t)

	f := &_depFields{
		artistRepo:  mockArtRepo,
		managerRepo: mockManagerRepo,
		userRepo:    mockUserRepo,
		transactor:  transactionMock,
		scBroker:    mockBroker,
		signReqRepo: mockSignReqRepo,
	}

	return f
}

func TestPublishRequestUseCase_proceedToManager(t *testing.T) {

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
						Status:    base.NewRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 0,
					},
					Nickname:    "skibidi",
					Description: "",
				},
			},
			out: nil,
			dependencies: func(df *_depFields) {

				df.managerRepo.EXPECT().GetRandManagerID(mock.AnythingOfType("context.backgroundCtx")).Return(
					uint64(9), nil).Once()

				df.signReqRepo.EXPECT().Update(mock.AnythingOfType("context.backgroundCtx"), &sign_contract.SignContractRequest{
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
				}).Return(nil).Once()

			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "NoManager",
			in: &args{
				signReq: &sign_contract.SignContractRequest{
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
				},
			},
			out: sctErrors.ErrCantFindManager,
			dependencies: func(df *_depFields) {

				df.managerRepo.EXPECT().GetRandManagerID(mock.AnythingOfType("context.backgroundCtx")).Return(
					0, errors.New("skpdjf")).Once()

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

			signReqHandler := InitSignContractProceedToManagerHandler(f.scBroker, f.signReqRepo, f.managerRepo)

			// act
			err := signReqHandler.(*SignContractProceedToManagerHandler).proceedToManager(tt.in.signReq)

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
