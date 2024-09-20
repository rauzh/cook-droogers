package base

import (
	baseReqErrors "cookdroogers/internal/requests/base/errors"
	cdtime "cookdroogers/pkg/time"
	"errors"
	"testing"
)

var dberr = errors.New("db err")

// TEST WITHOUT MOCK / STUB
func TestBaseRequest_Validate(t *testing.T) {

	type args struct {
		reqType RequestType
	}

	tests := []struct {
		name   string
		actor  *Request
		in     *args
		outErr error

		assert func(*testing.T)
	}{
		{
			name: "OK",
			actor: &Request{
				RequestID: 1,
				Type:      "Sign",
				Status:    OnApprovalRequest,
				Date:      cdtime.GetToday(),
				ApplierID: 12,
				ManagerID: 9,
			},
			in: &args{
				reqType: "Sign",
			},
			outErr: nil,

			assert: func(t *testing.T) {

			},
		},
		{
			name: "ErrInvalidType",
			actor: &Request{
				RequestID: 1,
				Type:      "Sign",
				Status:    OnApprovalRequest,
				Date:      cdtime.GetToday(),
				ApplierID: 12,
				ManagerID: 9,
			},
			in: &args{
				reqType: "Publish",
			},
			outErr: baseReqErrors.ErrInvalidType,

			assert: func(t *testing.T) {

			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// act
			err := tt.actor.Validate(tt.in.reqType)

			// assert
			if !errors.Is(err, tt.outErr) {
				t.Errorf("got %v, want %v", err, tt.outErr)
			}
			if tt.assert != nil {
				tt.assert(t)
			}
		})
	}
}
