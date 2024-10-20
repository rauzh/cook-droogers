package service

import (
	"context"
	"cookdroogers/internal/repo"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"database/sql"
	"github.com/pkg/errors"
	"log/slog"
	"strings"
)

type IArtistService interface {
	Create(context.Context, *models.Artist) error
	Get(context.Context, uint64) (*models.Artist, error)
	GetByUserID(ctx context.Context, id uint64) (*models.Artist, error)
	Update(context.Context, *models.Artist) error
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

func (ars *ArtistService) Create(ctx context.Context, artist *models.Artist) error {
	if err := ars.repo.Create(ctx, artist); err != nil {
		return errors.Wrap(CreateDbError, err.Error())
	}
	return nil
}

func (ars *ArtistService) Get(ctx context.Context, id uint64) (*models.Artist, error) {
	artist, err := ars.repo.Get(ctx, id)
	if err != nil && strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		return nil, userErrors.ErrNoUser
	}
	if err != nil {
		return nil, errors.Wrap(GetDbError, err.Error())
	}
	return artist, nil
}

func (ars *ArtistService) GetByUserID(ctx context.Context, id uint64) (*models.Artist, error) {
	artist, err := ars.repo.GetByUserID(ctx, id)
	if err != nil && strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		return nil, userErrors.ErrNoUser
	}
	if err != nil {
		return nil, errors.Wrap(GetDbError, err.Error())
	}
	return artist, nil
}

func (ars *ArtistService) Update(ctx context.Context, artist *models.Artist) error {
	if err := ars.repo.Update(ctx, artist); err != nil {
		return errors.Wrap(UpdateDbError, err.Error())
	}
	return nil
}
