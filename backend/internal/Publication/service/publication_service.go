package service

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"fmt"
	"log/slog"
	"time"
)

type IPublicationService interface {
	Create(context.Context, *models.Publication) error
	Update(context.Context, *models.Publication) error
	Get(context.Context, uint64) (*models.Publication, error)
	GetAllByDate(ctx context.Context, date time.Time) ([]models.Publication, error)
	GetAllByManager(ctx context.Context, managerID uint64) ([]models.Publication, error)
	GetAllByArtistSinceDate(ctx context.Context, date time.Time, artistID uint64) ([]models.Publication, error)
}

type PublicationService struct {
	repo   repo.PublicationRepo
	logger *slog.Logger
}

func NewPublicationService(repo repo.PublicationRepo, logger *slog.Logger) IPublicationService {
	return &PublicationService{repo: repo, logger: logger}
}

func (pbcSvc *PublicationService) Create(ctx context.Context, publication *models.Publication) error {
	if err := pbcSvc.repo.Create(ctx, publication); err != nil {
		return fmt.Errorf("can't create publication info with error %w", err)
	}
	return nil
}

func (pbcSvc *PublicationService) Update(ctx context.Context, publication *models.Publication) error {
	if err := pbcSvc.repo.Update(ctx, publication); err != nil {
		return fmt.Errorf("can't update publication info with error %w", err)
	}
	return nil
}

func (pbcSvc *PublicationService) Get(ctx context.Context, id uint64) (*models.Publication, error) {
	publication, err := pbcSvc.repo.Get(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}

func (pbcSvc *PublicationService) GetAllByDate(ctx context.Context, date time.Time) ([]models.Publication, error) {
	publications, err := pbcSvc.repo.GetAllByDate(ctx, date)

	if err != nil {
		return nil, fmt.Errorf("can't get publications info with error %w", err)
	}
	return publications, nil
}

func (pbcSvc *PublicationService) GetAllByManager(ctx context.Context, mng uint64) ([]models.Publication, error) {
	publications, err := pbcSvc.repo.GetAllByManager(ctx, mng)

	if err != nil {
		return nil, fmt.Errorf("can't get publications info with error %w", err)
	}
	return publications, nil
}

func (pbcSvc *PublicationService) GetAllByArtistSinceDate(ctx context.Context, date time.Time, artistID uint64) ([]models.Publication, error) {
	publications, err := pbcSvc.repo.GetAllByArtistSinceDate(ctx, date, artistID)

	if err != nil {
		return nil, fmt.Errorf("can't get publications info with error %w", err)
	}
	return publications, nil
}
