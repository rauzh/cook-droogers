package publish

import (
	"cookdroogers/internal/requests/base"
	"time"
)

const (
	EmptyID = 0
)

type PublishRequest struct {
	base.Request
	ReleaseID    uint64
	Grade        int
	ExpectedDate time.Time
	Description  string
}
