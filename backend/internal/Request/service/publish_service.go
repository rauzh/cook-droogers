package service

import "time"

type IPublishService interface {
	Apply(applierID, releaseID uint64, date time.Time) error
	Accept(uint64) error
	Decline(uint64) error
}
