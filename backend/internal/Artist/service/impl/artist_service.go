package service

import (
	artistRepo "cookdroogers/internal/Artist/repo"
	s "cookdroogers/internal/Artist/service"
	"cookdroogers/models"
	"fmt"
)

type ArtistService struct {
	repo artistRepo.ArtistRepo
}

func NewArtistService(r artistRepo.ArtistRepo) s.IArtistService {
	return &ArtistService{repo: r}
}

func (ars *ArtistService) Create(artist *models.Artist) error {
	if err := ars.repo.Create(artist); err != nil {
		return fmt.Errorf("can't create artist with err %w", err)
	}
	return nil
}

func (ars *ArtistService) Get(id uint64) (*models.Artist, error) {
	artist, err := ars.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get artist with err %w", err)
	}
	return artist, nil
}

func (ars *ArtistService) GetAll() ([]models.Artist, error) {
	artists, err := ars.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("can't get artists with err %w", err)
	}
	return artists, nil
}

func (ars *ArtistService) Update(artist *models.Artist) error {
	if err := ars.repo.Update(artist); err != nil {
		return fmt.Errorf("can't update artist with err %w", err)
	}
	return nil
}