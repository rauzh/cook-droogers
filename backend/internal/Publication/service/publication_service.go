package service

import (
	"cookdroogers/models"
)

type IPublicationService interface {
	Create(*models.Publication) error
	Update(*models.Publication) error
	Get(uint64) (*models.Publication, error)
}
