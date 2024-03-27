package service

import (
	"cookdroogers/internal/Publication/repo"
	s "cookdroogers/internal/Publication/service"
	"cookdroogers/models"
	"fmt"
	"time"
)

type PublicationService struct {
	repo repo.PublicationRepo
}

func NewPublicationService(
	repo repo.PublicationRepo) s.IPublicationService {
	return &PublicationService{
		repo: repo,
	}
}

func (ps *PublicationService) Create(publication *models.Publication) error {
	if err := ps.repo.Create(publication); err != nil {
		return fmt.Errorf("can't create publication info with error %w", err)
	}
	return nil
}

func (ps *PublicationService) Update(publication *models.Publication) error {
	if err := ps.repo.Update(publication); err != nil {
		return fmt.Errorf("can't update publication info with error %w", err)
	}
	return nil
}

func (ps *PublicationService) Get(id uint64) (*models.Publication, error) {
	publication, err := ps.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}

func (ps *PublicationService) GetAllByDate(date time.Time) ([]models.Publication, error) {
	publication, err := ps.repo.GetAllByDate(date)
	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}

func (ps *PublicationService) GetAllByManager(mng uint64) ([]models.Publication, error) {
	publication, err := ps.repo.GetAllByManager(mng)
	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}

func (ps *PublicationService) GetAllByArtistSinceDate(date time.Time, artistID uint64) ([]models.Publication, error) {
	publication, err := ps.repo.GetAllByArtistSinceDate(date, artistID)
	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}
