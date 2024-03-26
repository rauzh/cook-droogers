package service

import (
	artistService "cookdroogers/internal/Artist/service"
	managerService "cookdroogers/internal/Manager/service"
	requestErrors "cookdroogers/internal/Request/errors"
	requestService "cookdroogers/internal/Request/service"
	userService "cookdroogers/internal/User/service"
	"cookdroogers/models"
	"fmt"
	"time"
)

type SignContractService struct {
	reqSvc requestService.IRequestService
	mngSvc managerService.IManagerService
	usrSvc userService.IUserService
	artSvc artistService.IArtistService
}

func NewSignContractService(
	reqSvc requestService.IRequestService,
	mngSvc managerService.IManagerService,
	usrSvc userService.IUserService,
	artSvc artistService.IArtistService,
) requestService.ISignContractService {
	return &SignContractService{
		reqSvc: reqSvc,
		mngSvc: mngSvc,
		usrSvc: usrSvc,
		artSvc: artSvc,
	}
}

func (sctSvc *SignContractService) Apply(userID uint64, nickname string) error {

	if nickname == "" {
		return requestErrors.ErrNoNickname
	}

	request := models.Request{
		Type: models.SignRequest,
		Meta: map[string]string{
			"nickname": nickname,
			"descr":    "",
		},
		ApplierID: userID,
	}

	if err := sctSvc.reqSvc.Create(&request); err != nil {
		return fmt.Errorf("can't create request with err %w", err)
	}

	go sctSvc.proceedToManager(request)

	return nil
}

func (sctSvc *SignContractService) proceedToManager(request models.Request) {
	request.Status = models.OnApprovalRequest

	managerID, err := sctSvc.mngSvc.GetRandomManagerID()
	if err == nil {
		request.ManagerID = managerID
	} else {
		request.Status = models.ClosedRequest
		request.Meta["descr"] = "Can't find manager"
	}

	sctSvc.reqSvc.Update(&request)
}

func (sctSvc *SignContractService) Accept(requestID uint64) error {

	request, err := sctSvc.reqSvc.Get(requestID)
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

	if err := sctSvc.artSvc.Create(&artist); err != nil {
		return fmt.Errorf("can't create artist %s with err %w", artist.Nickname, err)
	}

	if err := sctSvc.usrSvc.UpdateType(artist.UserID, models.ArtistUser); err != nil {
		return fmt.Errorf("can't update user with err %w", err)
	}

	request.Status = models.ClosedRequest
	if err := sctSvc.reqSvc.Update(request); err != nil {
		return fmt.Errorf("can't update reqiest with err %w", err)
	}

	return nil
}

func (sctSvc *SignContractService) Decline(requestID uint64) error {
	request, err := sctSvc.reqSvc.Get(requestID)
	if err != nil {
		return fmt.Errorf("can't get request %d with err %w", requestID, err)
	}

	request.Status = models.ClosedRequest
	request.Meta["descr"] = "The request is declined."

	return sctSvc.reqSvc.Update(request)
}
