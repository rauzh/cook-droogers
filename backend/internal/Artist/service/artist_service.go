package service

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"fmt"
)

type IArtistService interface {
	Create(*models.Artist) error
	Get(uint64) (*models.Artist, error)
	GetByUserID(id uint64) (*models.Artist, error)
	Update(*models.Artist) error
}

type ArtistService struct {
	repo repo.ArtistRepo
}

func NewArtistService(r repo.ArtistRepo) IArtistService {
	return &ArtistService{repo: r}
}

func (ars *ArtistService) Create(artist *models.Artist) error {
	if err := ars.repo.Create(context.Background(), artist); err != nil {
		return fmt.Errorf("can't create artist with err %w", err)
	}
	return nil
}

func (ars *ArtistService) Get(id uint64) (*models.Artist, error) {
	artist, err := ars.repo.Get(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("can't get artist with err %w", err)
	}
	return artist, nil
}

func (ars *ArtistService) GetByUserID(id uint64) (*models.Artist, error) {
	artist, err := ars.repo.GetByUserID(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("can't get artist with err %w", err)
	}
	return artist, nil
}

func (ars *ArtistService) Update(artist *models.Artist) error {
	if err := ars.repo.Update(context.Background(), artist); err != nil {
		return fmt.Errorf("can't update artist with err %w", err)
	}
	return nil
}
