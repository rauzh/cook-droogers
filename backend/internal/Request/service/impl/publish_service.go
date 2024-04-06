package service

import (
	artistService "cookdroogers/internal/Artist/service"
	publicationService "cookdroogers/internal/Publication/service"
	releaseService "cookdroogers/internal/Release/service"
	requestErrors "cookdroogers/internal/Request/errors"
	requestService "cookdroogers/internal/Request/service"
	statisticsServive "cookdroogers/internal/Statistics/service"
	transactionManager "cookdroogers/internal/TransactionManager"
	"cookdroogers/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type PublishService struct {
	reqSvc             requestService.IRequestService
	artSvc             artistService.IArtistService
	statSvc            statisticsServive.IStatisticsService
	rlsSvc             releaseService.IReleaseService
	pblcSvc            publicationService.IPublicationService
	transactionManager transactionManager.TransactionManager
}

const DefaultGrade uint8 = 5

func NewPublishService(
	reqSvc requestService.IRequestService,
	artSvc artistService.IArtistService,
	statSvc statisticsServive.IStatisticsService,
	rlsSvc releaseService.IReleaseService,
	pblcSvc publicationService.IPublicationService,
	transactionMngr transactionManager.TransactionManager) requestService.IPublishService {
	return &PublishService{
		reqSvc:             reqSvc,
		artSvc:             artSvc,
		statSvc:            statSvc,
		rlsSvc:             rlsSvc,
		pblcSvc:            pblcSvc,
		transactionManager: transactionMngr,
	}
}

// сделать нормальный ооп с наследованием от общей базовой заявки
func (pblSvc *PublishService) Apply(applierID, releaseID uint64, date time.Time) error {

	// Publication date must be at least week later than application date
	if time.Now().UTC().Add(Week).After(date) {
		return requestErrors.ErrLessThanWeek
	}

	request := models.Request{
		Type: models.PublishRequest,
		Meta: map[string]string{
			"release": fmt.Sprintf("%d", releaseID),
			"grade":   "5",
			"date":    date.String(),
			"descr":   ""},
		ApplierID: applierID,
	}

	// Create request.go instance in db
	if err := pblSvc.reqSvc.Create(&request); err != nil {
		return err
	}

	// Async process request.go and proceed it to manager ДОКРУТИТЬ НАДЕЖНОСТЬ!!!
	go pblSvc.proceedToManager(request, releaseID, date)

	return nil
}

func (pblSvc *PublishService) proceedToManager(request models.Request, releaseID uint64, date time.Time) {

	// Set request.go status to "Processing" & UPD db
	request.Status = models.ProcessingRequest

	err := pblSvc.reqSvc.Update(&request)
	if err != nil {
		panic("CREATE-PUBL-APPL: Can't update APPL")
	}

	// Analyze request.go and set degree
	pblSvc.computeDegree(&request, releaseID, date)

	// Get applier's data
	artist, err := pblSvc.artSvc.Get(request.ApplierID)
	if err != nil {
		panic("CREATE-PUBL-APPL: Can't get ARTIST")
	}

	// Proceed request.go to the artist's manager & change its status to "On Approval"
	request.ManagerID = artist.ManagerID
	request.Status = models.OnApprovalRequest

	err = pblSvc.reqSvc.Update(&request)
	if err != nil {
		panic("CREATE-PUBL-APPL: Can't update APPL")
	}
}

// пробег по проверкам которые в итоге возрвращаютт структуру с диффом оценки и описанием с каким-нибудь реквест валидатор (чтобы можно бвло норм расширять)
func (pblSvc *PublishService) computeDegree(request *models.Request, releaseID uint64, date time.Time) {
	var grade uint8 = DefaultGrade

	pubsThatDay, errPubsThatDay := pblSvc.pblcSvc.GetAllByDate(date)
	// If more than 1 release that day, decrement the grade
	if errPubsThatDay != nil && len(pubsThatDay) > 1 {
		grade--
		request.Meta["descr"] += "\n Too many releases that day"
	}

	pubsFromThatArtistLastSeason, errPubsFromArtist := pblSvc.pblcSvc.GetAllByArtistSinceDate(
		date.AddDate(0, -3, 1),
		request.ApplierID,
	)
	// If more than 2 releases from that artist in last 3 months, decrement the grade
	if errPubsFromArtist != nil && len(pubsFromThatArtistLastSeason) > 2 {
		grade--
		request.Meta["descr"] += "\n Too many releases from this artist"
	}

	relevantGenre, errGetRelevantGenre := pblSvc.statSvc.GetRelevantGenre()
	currentGenre, errGetCurGenre := pblSvc.rlsSvc.GetMainGenre(releaseID)
	// If pub's genre is not relevant, decrement the grade
	if errGetRelevantGenre != nil && errGetCurGenre != nil && currentGenre != relevantGenre {
		grade--
		request.Meta["descr"] += "\n Not relevant genre"
	}

	request.Meta["grade"] = fmt.Sprintf("%d", grade)
}

func (pblSvc *PublishService) Accept(requestID uint64) error {

	request, err := pblSvc.reqSvc.Get(requestID)
	if err != nil {
		return fmt.Errorf("can't get request.go %d with err %w", requestID, err)
	}

	// Get releaseID from meta
	releaseID, err := strconv.ParseUint(request.Meta["release"], 10, 64)
	if err != nil {
		return requestErrors.ErrInvalidMetaReleaseID
	}

	// Get date from meta время в константу
	date, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", request.Meta["date"])
	if err != nil {
		return errors.New("can't get date from publication request")
	}

	publication := models.Publication{
		ReleaseID: releaseID,
		Date:      date,
		ManagerID: request.ManagerID,
	}

	transactionHash, err := pblSvc.transactionManager.BeginTransaction()
	if err != nil {
		return err
	}
	// Create publication
	if err := pblSvc.pblcSvc.Create(&publication); err != nil {
		pblSvc.transactionManager.RollbackTransaction(transactionHash)
		return fmt.Errorf("can't create publication with err %w", err)
	}

	// Update associated release and request.go
	if err := pblSvc.rlsSvc.UpdateStatus(releaseID, models.PublishedRelease); err != nil {
		pblSvc.transactionManager.RollbackTransaction(transactionHash)
		return fmt.Errorf("can't update publication with err %w", err)
	}

	request.Status = models.ClosedRequest
	if err := pblSvc.reqSvc.Update(request); err != nil {
		pblSvc.transactionManager.RollbackTransaction(transactionHash)
		return fmt.Errorf("can't update request.go with err %w", err)
	}

	err = pblSvc.transactionManager.CommitTransaction(transactionHash)
	return err
}

func (pblSvc *PublishService) Decline(requestID uint64) error {
	request, err := pblSvc.reqSvc.Get(requestID)
	if err != nil {
		return fmt.Errorf("can't get request.go %d with err %w", requestID, err)
	}

	request.Status = models.ClosedRequest
	request.Meta["descr"] = "The request.go is declined."

	return pblSvc.reqSvc.Update(request)
}
