package service

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"github.com/pkg/errors"
	"log/slog"
)

type IArtistService interface {
	Create(*models.Artist) error
	Get(uint64) (*models.Artist, error)
	GetByUserID(id uint64) (*models.Artist, error)
	Update(*models.Artist) error
}

var CreateDbError error = errors.New("can't create artist")
var GetDbError error = errors.New("can't get artist")
var UpdateDbError error = errors.New("can't update artist")

type ArtistService struct {
	repo   repo.ArtistRepo
	logger *slog.Logger
}

func NewArtistService(r repo.ArtistRepo, logger *slog.Logger) IArtistService {
	return &ArtistService{repo: r, logger: logger}
}

func (ars *ArtistService) Create(artist *models.Artist) error {
	if err := ars.repo.Create(context.Background(), artist); err != nil {
		return errors.Wrap(CreateDbError, err.Error())
	}
	return nil
}

func (ars *ArtistService) Get(id uint64) (*models.Artist, error) {
	artist, err := ars.repo.Get(context.Background(), id)

	if err != nil {
		return nil, errors.Wrap(GetDbError, err.Error())
	}
	return artist, nil
}

func (ars *ArtistService) GetByUserID(id uint64) (*models.Artist, error) {
	artist, err := ars.repo.GetByUserID(context.Background(), id)

	if err != nil {
		return nil, errors.Wrap(GetDbError, err.Error())
	}
	return artist, nil
}

func (ars *ArtistService) Update(artist *models.Artist) error {
	if err := ars.repo.Update(context.Background(), artist); err != nil {
		return errors.Wrap(UpdateDbError, err.Error())
	}
	return nil
}
