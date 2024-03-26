package service

import (
	artistRepo "cookdroogers/internal/Artist/repo"
	s "cookdroogers/internal/Artist/service"
	ms "cookdroogers/internal/Manager/service"
	as "cookdroogers/internal/Request/service"
	us "cookdroogers/internal/User/service"
	"cookdroogers/models"
	"fmt"
	"time"
)

type ArtistService struct {
	requestService as.IRequestService
	managerService ms.IManagerService
	userService    us.IUserService
	repo           artistRepo.ArtistRepo
}

func NewArtistService(
	as as.IRequestService,
	ms ms.IManagerService,
	us us.IUserService,
	r artistRepo.ArtistRepo) s.IArtistService {
	return &ArtistService{
		requestService: as,
		managerService: ms,
		userService:    us,
		repo:           r,
	}
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

func (ars *ArtistService) CreateSignRequest(userID uint64, nickname string) error {

	if nickname == "" {
		return fmt.Errorf("no nickname provided")
	}

	request := models.Request{
		Type: models.SignRequest,
		Meta: map[string]string{
			"nickname": nickname,
			"descr":    "",
		},
		ApplierID: userID,
	}

	if err := ars.requestService.Create(&request); err != nil {
		return fmt.Errorf("can't create request with err %w", err)
	}

	go func(managerService ms.IManagerService, requestService as.IRequestService) {

		request.Status = models.OnApprovalRequest

		managerID, err := managerService.GetRandomManagerID()
		if err == nil {
			request.ManagerID = managerID
		} else {
			request.Status = models.ClosedRequest
			request.Meta["descr"] = "Can't find manager"
		}

		requestService.Update(&request)

	}(ars.managerService, ars.requestService)

	return nil
}

func (ars *ArtistService) ApplySignRequest(requestID uint64) error {

	request, err := ars.requestService.Get(requestID)
	if err != nil {
		return fmt.Errorf("can't get request %d with err %w", requestID, err)
	}

	artist := models.Artist{
		UserID:       request.ApplierID,
		Nickname:     request.Meta["nickname"],
		ContractTerm: time.Now().AddDate(1, 0, 0),
		Activity:     true,
		ManagerID:    request.ManagerID,
	}

	if err := ars.Create(&artist); err != nil {
		return fmt.Errorf("can't create artist %s with err %w", artist.Nickname, err)
	}

	go func() {
		user, err := ars.userService.Get(artist.UserID)
		if err != nil {
			panic("APPLY-SIGN-APPL: Can't get USER")
		}
		user.Type = models.ArtistUser
		ars.userService.Update(user)

		request.Status = models.ClosedRequest
		ars.requestService.Update(request)
	}()

	return nil
}

func (ars *ArtistService) DeclineSignRequest(requestID uint64) error {
	request, err := ars.requestService.Get(requestID)
	if err != nil {
		return fmt.Errorf("can't get request %d with err %w", requestID, err)
	}

	request.Status = models.ClosedRequest
	request.Meta["descr"] = "The request is declined."
	ars.requestService.Update(request)

	return nil
}
