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
	sarama_mocks "github.com/IBM/sarama/mocks"
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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()

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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()

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

			signReqUseCase, err := NewSignContractRequestUseCase(f.managerRepo, f.userRepo, f.artistRepo, f.transactor, f.scBroker, f.signReqRepo)

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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()

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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()

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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()

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

			signReqUseCase, err := NewSignContractRequestUseCase(f.managerRepo, f.userRepo, f.artistRepo, f.transactor, f.scBroker, f.signReqRepo)

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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockSignReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			signReqUseCase, err := NewSignContractRequestUseCase(f.managerRepo, f.userRepo, f.artistRepo, f.transactor, f.scBroker, f.signReqRepo)

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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()

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

				df.scBroker.EXPECT().SignConsumerToTopic(SignRequestProceedToManager).Return(nil).Once()
				df.scBroker.EXPECT().GetConsumerByTopic(SignRequestProceedToManager).Return(&sarama_mocks.PartitionConsumer{}).Once()

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

			signReqUseCase, err := NewSignContractRequestUseCase(f.managerRepo, f.userRepo, f.artistRepo, f.transactor, f.scBroker, f.signReqRepo)

			// act
			err = signReqUseCase.(*SignContractRequestUseCase).proceedToManager(tt.in.signReq)

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
