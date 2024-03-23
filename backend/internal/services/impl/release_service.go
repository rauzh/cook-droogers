package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	service "cookdroogers/internal/services"
	"fmt"
)

type ReleaseService struct {
	trackService service.ITrackService
	repo         repo.ReleaseRepo
}

func (rs *ReleaseService) Create(artist *models.Release) error {
	if err := rs.repo.Create(artist); err != nil {
		return fmt.Errorf("can't create release with err %w", err)
	}
	return nil
}

func (rs *ReleaseService) Get(releaseID uint64) (*models.Release, error) {
	release, err := rs.repo.Get(releaseID)
	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return release, nil
}

func (rs *ReleaseService) Update(artist *models.Release) error {
	if err := rs.repo.Update(artist); err != nil {
		return fmt.Errorf("can't update release with err %w", err)
	}
	return nil
}

func (rs *ReleaseService) GetMainGenre(releaseID uint64) (string, error) {
	release, err := rs.repo.Get(releaseID)
	if err != nil {
		return "", fmt.Errorf("can't get release with err %w", err)
	}

	genres := map[string]int{}
	for _, trackID := range release.Tracks {
		track, err := rs.trackService.Get(trackID)
		if err != nil {
			return "", fmt.Errorf("can't get track %d with err %w", trackID, err)
		}

		genres[track.Genre]++
	}

	var maxAmount int
	var relevantGenre string
	for genre, amount := range genres {
		if amount > maxAmount {
			maxAmount = amount
			relevantGenre = genre
		}
	}

	return relevantGenre, nil
}
