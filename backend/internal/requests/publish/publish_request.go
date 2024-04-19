package publish

import (
	"cookdroogers/internal/requests/base"
	pubReqErrors "cookdroogers/internal/requests/publish/errors"
	cdtime "cookdroogers/pkg/time"
	"time"
)

const (
	EmptyID                  = 0
	PubReq  base.RequestType = "Publish"
)

type PublishRequest struct {
	base.Request
	ReleaseID    uint64
	Grade        int
	ExpectedDate time.Time
	Description  string
}

func NewPublishRequest(applierID uint64, releaseID uint64, expectedDate time.Time) base.IRequest {

	return &PublishRequest{
		Request: base.Request{
			Type:      PubReq,
			ApplierID: applierID,
		},
		ReleaseID:    releaseID,
		ExpectedDate: expectedDate,
	}
}

func (pubReq *PublishRequest) Validate(reqType base.RequestType) error {

	if err := pubReq.Request.Validate(reqType); err != nil {
		return err
	}

	if pubReq.ExpectedDate.IsZero() || !cdtime.CheckDateWeekLater(pubReq.ExpectedDate) {
		return pubReqErrors.ErrInvalidDate
	}

	if pubReq.ReleaseID == base.EmptyID {
		return pubReqErrors.ErrNoReleaseID
	}

	return nil
}

func (pubReq *PublishRequest) GetType() base.RequestType {
	return pubReq.Type
}
