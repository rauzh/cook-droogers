package service

import (
	"cookdroogers/models"
	"time"
)

type IPublicationService interface {
	Create(*models.Publication) error
	Update(*models.Publication) error
	Get(uint64) (*models.Publication, error)
	CreatePublApplication(applierID, releaseID uint64, date time.Time) error
	ApplyPublApplication(uint64) error
	DeclinePublApplication(uint64) error
}
