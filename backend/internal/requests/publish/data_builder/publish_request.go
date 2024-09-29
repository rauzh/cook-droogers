package data_builder

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	cdtime "cookdroogers/pkg/time"
	"time"
)

type PublishRequestBuilder struct {
	PublishRequest *publish.PublishRequest
}

func NewPublishRequestBuilder() *PublishRequestBuilder {
	return &PublishRequestBuilder{
		PublishRequest: &publish.PublishRequest{
			Request: base.Request{
				RequestID: 123,
				Type:      publish.PubReq,
				ManagerID: 8,
				ApplierID: 88,
				Status:    base.ProcessingRequest,
				Date:      cdtime.GetToday(),
			},
			ReleaseID:    100,
			Description:  "Test description",
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Grade:        5,
		},
	}
}

func (b *PublishRequestBuilder) WithRequestID(Requestid uint64) *PublishRequestBuilder {
	b.PublishRequest.RequestID = Requestid
	return b
}

func (b *PublishRequestBuilder) WithApplierID(id uint64) *PublishRequestBuilder {
	b.PublishRequest.ApplierID = id
	return b
}

func (b *PublishRequestBuilder) WithManagerID(id uint64) *PublishRequestBuilder {
	b.PublishRequest.ManagerID = id
	return b
}

func (b *PublishRequestBuilder) WithType(t base.RequestType) *PublishRequestBuilder {
	b.PublishRequest.Type = t
	return b
}

func (b *PublishRequestBuilder) WithStatus(s base.RequestStatus) *PublishRequestBuilder {
	b.PublishRequest.Status = s
	return b
}

func (b *PublishRequestBuilder) WithDate(date time.Time) *PublishRequestBuilder {
	b.PublishRequest.Date = date
	return b
}

func (b *PublishRequestBuilder) WithExpexctedDate(date time.Time) *PublishRequestBuilder {
	b.PublishRequest.ExpectedDate = date
	return b
}

func (b *PublishRequestBuilder) WithReleaseID(id uint64) *PublishRequestBuilder {
	b.PublishRequest.ReleaseID = id
	return b
}

func (b *PublishRequestBuilder) WithGrade(gr int) *PublishRequestBuilder {
	b.PublishRequest.Grade = gr
	return b
}

func (b *PublishRequestBuilder) WithDescription(descr string) *PublishRequestBuilder {
	b.PublishRequest.Description = descr
	return b
}

func (b *PublishRequestBuilder) Build() *publish.PublishRequest {
	return b.PublishRequest
}
