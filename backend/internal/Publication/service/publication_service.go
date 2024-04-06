package service

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"fmt"
	"time"
)

type IPublicationService interface {
	Create(*models.Publication) error
	Update(*models.Publication) error
	Get(uint64) (*models.Publication, error)
	GetAllByDate(date time.Time) ([]models.Publication, error)
	GetAllByManager(managerID uint64) ([]models.Publication, error)
	GetAllByArtistSinceDate(date time.Time, artistID uint64) ([]models.Publication, error)
}

type PublicationService struct {
	repo repo.PublicationRepo
}

func NewPublicationService(repo repo.PublicationRepo) IPublicationService {
	return &PublicationService{repo: repo}
}

func (pbcSvc *PublicationService) Create(publication *models.Publication) error {
	if err := pbcSvc.repo.Create(context.Background(), publication); err != nil {
		return fmt.Errorf("can't create publication info with error %w", err)
	}
	return nil
}

func (pbcSvc *PublicationService) Update(publication *models.Publication) error {
	if err := pbcSvc.repo.Update(context.Background(), publication); err != nil {
		return fmt.Errorf("can't update publication info with error %w", err)
	}
	return nil
}

func (pbcSvc *PublicationService) Get(id uint64) (*models.Publication, error) {
	publication, err := pbcSvc.repo.Get(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}

func (pbcSvc *PublicationService) GetAllByDate(date time.Time) ([]models.Publication, error) {
	publications, err := pbcSvc.repo.GetAllByDate(context.Background(), date)

	if err != nil {
		return nil, fmt.Errorf("can't get publications info with error %w", err)
	}
	return publications, nil
}

func (pbcSvc *PublicationService) GetAllByManager(mng uint64) ([]models.Publication, error) {
	publications, err := pbcSvc.repo.GetAllByManager(context.Background(), mng)

	if err != nil {
		return nil, fmt.Errorf("can't get publications info with error %w", err)
	}
	return publications, nil
}

func (pbcSvc *PublicationService) GetAllByArtistSinceDate(date time.Time, artistID uint64) ([]models.Publication, error) {
	publications, err := pbcSvc.repo.GetAllByArtistSinceDate(context.Background(), date, artistID)

	if err != nil {
		return nil, fmt.Errorf("can't get publications info with error %w", err)
	}
	return publications, nil
}
