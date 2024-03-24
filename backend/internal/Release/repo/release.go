package repo

import "cookdroogers/models"

type ReleaseRepo interface {
	Create(*models.Release) error
	Get(uint64) (*models.Release, error)
	Update(*models.Release) error
}
