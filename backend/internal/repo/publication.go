package repo

import "cookdroogers/internal/models"

type PublicationRepo interface {
	Create(*models.Publication) error
	Get(uint64) (*models.Publication, error)
	Update(*models.Publication) error
}
