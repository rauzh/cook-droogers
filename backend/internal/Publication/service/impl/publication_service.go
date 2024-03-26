package service

import (
	ars "cookdroogers/internal/Artist/service"
	ms "cookdroogers/internal/Manager/service"
	"cookdroogers/internal/Publication/repo"
	s "cookdroogers/internal/Publication/service"
	rs "cookdroogers/internal/Release/service"
	as "cookdroogers/internal/Request/service"
	ss "cookdroogers/internal/Statistics/service"
	"cookdroogers/models"
	"fmt"
	"time"
)

type PublicationService struct {
	requestService as.IRequestService
	releaseService rs.IReleaseService
	managerService ms.IManagerService
	artistService  ars.IArtistService
	statService    ss.IStatisticsService
	repo           repo.PublicationRepo
}

func NewPublicationService(
	as as.IRequestService,
	rs rs.IReleaseService,
	ms ms.IManagerService,
	ars ars.IArtistService,
	ss ss.IStatisticsService,
	repo repo.PublicationRepo) s.IPublicationService {
	return &PublicationService{
		requestService: as,
		releaseService: rs,
		managerService: ms,
		artistService:  ars,
		statService:    ss,
		repo:           repo,
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

func (ps *PublicationService) GetAllByArtistSinceDate(date time.Time, artistID uint64) ([]models.Publication, error) {
	publication, err := ps.repo.GetAllByArtistSinceDate(date, artistID)
	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}
