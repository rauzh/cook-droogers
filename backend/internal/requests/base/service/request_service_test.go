package service

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/base/repo/mocks"
	"cookdroogers/internal/requests/sign_contract"
	cdtime "cookdroogers/pkg/time"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	reqRepo *mocks.RequestRepo
	logger  *slog.Logger

	res interface{}
}

func _newMockSignReqDepFields(t *testing.T) *_depFields {

	mockReqRepo := mocks.NewRequestRepo(t)

	f := &_depFields{
		reqRepo: mockReqRepo,
		logger:  slog.Default(),
	}

	return f
}

var testDBerr error = errors.New("some db err")

func TestRequestService_GetAllByManagerID(t *testing.T) {

	type args struct {
		id uint64
	}

	tests := []struct {
		name   string
		in     *args
		outErr error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields, []base.Request)
	}{
		{
			name: "OK",
			in: &args{
				id: 7,
			},
			outErr: nil,
			dependencies: func(df *_depFields) {

				df.reqRepo.EXPECT().GetAllByManagerID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
					Return([]base.Request{
						base.Request{
							RequestID: 1,
							Type:      sign_contract.SignRequest,
							Status:    base.OnApprovalRequest,
							Date:      cdtime.GetToday(),
							ApplierID: 12,
							ManagerID: 9,
						},
					}, nil).Once()
			},
			assert: func(t *testing.T, df *_depFields, reqs []base.Request) {
				assert.Equal(t, []base.Request{
					base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
				}, reqs)
			},
		},
		{
			name: "DB err",
			in: &args{
				id: 7,
			},
			outErr: DBerr,
			dependencies: func(df *_depFields) {
				df.reqRepo.EXPECT().GetAllByManagerID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
					Return(nil, testDBerr).Once()
			},
			assert: func(t *testing.T, df *_depFields, reqs []base.Request) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockSignReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			reqService := NewRequestService(f.reqRepo, f.logger)

			// act
			reqs, err := reqService.GetAllByManagerID(tt.in.id)

			// assert
			if !errors.Is(err, tt.outErr) {
				t.Errorf("got %v, want %v", err, tt.outErr)
			}
			if tt.assert != nil {
				tt.assert(t, f, reqs)
			}
		})
	}
}

func TestRequestService_GetAllByUserID(t *testing.T) {

	type args struct {
		id uint64
	}

	tests := []struct {
		name   string
		in     *args
		outErr error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields, []base.Request)
	}{
		{
			name: "OK",
			in: &args{
				id: 7,
			},
			outErr: nil,
			dependencies: func(df *_depFields) {

				df.reqRepo.EXPECT().GetAllByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
					Return([]base.Request{
						base.Request{
							RequestID: 1,
							Type:      sign_contract.SignRequest,
							Status:    base.OnApprovalRequest,
							Date:      cdtime.GetToday(),
							ApplierID: 12,
							ManagerID: 9,
						},
					}, nil).Once()
			},
			assert: func(t *testing.T, df *_depFields, reqs []base.Request) {
				assert.Equal(t, []base.Request{
					base.Request{
						RequestID: 1,
						Type:      sign_contract.SignRequest,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
				}, reqs)
			},
		},
		{
			name: "DB err",
			in: &args{
				id: 7,
			},
			outErr: DBerr,
			dependencies: func(df *_depFields) {
				df.reqRepo.EXPECT().GetAllByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
					Return(nil, testDBerr).Once()
			},
			assert: func(t *testing.T, df *_depFields, reqs []base.Request) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockSignReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			reqService := NewRequestService(f.reqRepo, f.logger)

			// act
			reqs, err := reqService.GetAllByUserID(tt.in.id)

			// assert
			if !errors.Is(err, tt.outErr) {
				t.Errorf("got %v, want %v", err, tt.outErr)
			}
			if tt.assert != nil {
				tt.assert(t, f, reqs)
			}
		})
	}
}

func TestRequestService_GetByID(t *testing.T) {

	type args struct {
		id uint64
	}

	mockRetReq := &base.Request{
		RequestID: 1,
		Type:      sign_contract.SignRequest,
		Status:    base.OnApprovalRequest,
		Date:      cdtime.GetToday(),
		ApplierID: 12,
		ManagerID: 9,
	}

	tests := []struct {
		name   string
		in     *args
		outErr error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields, *base.Request)
	}{
		{
			name: "OK",
			in: &args{
				id: 7,
			},
			outErr: nil,
			dependencies: func(df *_depFields) {
				df.reqRepo.EXPECT().GetByID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
					Return(mockRetReq, nil).Once()
			},
			assert: func(t *testing.T, df *_depFields, req *base.Request) {
				assert.Equal(t, mockRetReq, req)
			},
		},
		{
			name: "DB err",
			in: &args{
				id: 7,
			},
			outErr: DBerr,
			dependencies: func(df *_depFields) {
				df.reqRepo.EXPECT().GetByID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
					Return(nil, testDBerr).Once()
			},
			assert: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockSignReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			reqService := NewRequestService(f.reqRepo, f.logger)

			// act
			req, err := reqService.GetByID(tt.in.id)

			// assert
			if !errors.Is(err, tt.outErr) {
				t.Errorf("got %v, want %v", err, tt.outErr)
			}
			if tt.assert != nil {
				tt.assert(t, f, req)
			}
		})
	}
}
