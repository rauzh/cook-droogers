package techUI

import (
	"cookdroogers/app"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	usecase2 "cookdroogers/internal/requests/publish/usecase"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/usecase"
	"cookdroogers/models"
	"fmt"
	"log/slog"
)

type managerMenu struct {
	a       *app.App
	user    *models.User
	manager *models.Manager
	log     *slog.Logger
}

func initManagerMenu(a *app.App, user *models.User, log *slog.Logger) (*managerMenu, error) {

	err := a.Services.UserService.SetRole(models.ManagerUser)
	if err != nil {
		log.Error("Can't init manager menu: can't set manager role: ", slog.Any("error", err))
		return nil, err
	}

	manager, err := a.Services.ManagerService.GetByUserID(user.UserID)
	if err != nil {
		log.Error("Can't init manager menu: can't get manager: ", slog.Any("error", err))
		return nil, err
	}

	return &managerMenu{
		a:       a,
		user:    user,
		manager: manager,
		log:     log,
	}, nil
}

func (menu *managerMenu) Loop() error {

	startPosition :=
		`
		0 -- выйти
		1 -- открыть список заявок
		2 -- получить текущую статистику своих артистов
		3 -- подгрузить статистику своих артистов

	Выберите пункт меню: `

	for {
		fmt.Printf("%s", startPosition)

		var action int
		_, _ = fmt.Scanf("%d", &action)

		switch action {
		case 0:
			return ErrEXIT
		case 1:
			err := menu.lookupReqs()
			if err != nil {
				fmt.Println("Не удается обработать заявки по причине ", err)
				menu.log.Error("Can't look up requests: ", slog.Any("error", err))
			}
		case 2:
			menu.stats()
		case 3:
			err := menu.fetchStats()
			if err != nil {
				fmt.Println("Не удается подгрузить статистику по причине ", err)
				menu.log.Error("Can't fetch stats", slog.Any("error", err))
			} else {
				fmt.Println("Статистика успешно подгружена.")
			}
		default:
			fmt.Printf("Неверный пункт меню")
		}
	}
}

func (menu *managerMenu) stats() {

	report, err := menu.a.Services.ReportService.GetReportForManager(menu.manager.ManagerID)
	if err != nil {
		menu.log.Error("Can't get stats: ", slog.Any("error", err))
	}

	fmt.Printf("%s: %s\n\n", "relevant_genre", string(report["relevant_genre"]))
	fmt.Printf("%s: %s\n\n", "artists_stats", string(report["artists_stats"]))
}

func (menu *managerMenu) fetchStats() error {
	var errr error
	for _, artistID := range menu.manager.Artists {
		releases, err := menu.a.Services.ReleaseService.GetAllByArtist(artistID)
		if err != nil {
			return err
		}
		for _, release := range releases {
			err = menu.a.Services.StatService.FetchByRelease(&release)
			if err != nil {
				fmt.Println(err, " for release ", release.ReleaseID)
				errr = err
			}
		}
	}

	return errr
}

func (menu *managerMenu) lookupReqs() error {

	reqs, err := menu.a.Services.RequestService.GetAllByManagerID(menu.manager.ManagerID)
	if err != nil {
		return err
	}

	reqMap := make(map[uint64]base.Request)

	for _, req := range reqs {
		fmt.Printf("\n\t request id:%d\n\t type:%s\n\t status:%s\n\t date:%s\n\t manager:%d\n\t applier:%d\n",
			req.RequestID, req.Type, req.Status, req.Date, req.ManagerID, req.ApplierID)
		reqMap[req.RequestID] = req
	}

	fmt.Printf("\n%s", "Введите id заявки, на которую хотите ответить (0, если не хотите): ")
	var reqID uint64
	_, _ = fmt.Scanf("%d", &reqID)
	if reqID == 0 {
		return nil
	}

	switch reqMap[reqID].Type {
	case sign_contract.SignRequest:
		signReqUC := menu.a.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)

		signreq, err := signReqUC.Get(reqID)
		if err != nil {
			return err
		}

		fmt.Printf("\n\t nickname: %s \n\t decription: %s\n", signreq.Nickname, signreq.Description)
	case publish.PubReq:
		pubReqUC := menu.a.UseCases.PublishReqUC.(*usecase2.PublishRequestUseCase)

		pubreq, err := pubReqUC.Get(reqID)
		if err != nil {
			return err
		}

		fmt.Printf("\n\t expected date:%s\n\t decription: %s\n\t grade: %d\n\t release id: %d\n",
			pubreq.ExpectedDate, pubreq.Description, pubreq.Grade, pubreq.ReleaseID)
	}

	chooseReqAction :=
		`
		0 -- выйти
		1 -- принять
		2 -- отклонить

	Выберите действие: `
	fmt.Printf("%s", chooseReqAction)
	var action int
	_, _ = fmt.Scanf("%d", &action)

	switch action {
	case 0:
		return nil
	case 1:
		switch reqMap[reqID].Type {
		case sign_contract.SignRequest:
			signReqUC := menu.a.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
			signreq, err := signReqUC.Get(reqID)
			if err != nil {
				return err
			}
			err = menu.a.UseCases.SignContractReqUC.Accept(signreq)
			if err != nil {
				menu.log.Error("CANT ACCEPT SIGN CONTRACT", slog.Any("error", err))
				return err
			}
		case publish.PubReq:
			pubReqUC := menu.a.UseCases.PublishReqUC.(*usecase2.PublishRequestUseCase)
			pubreq, err := pubReqUC.Get(reqID)
			if err != nil {
				return err
			}
			err = menu.a.UseCases.PublishReqUC.Accept(pubreq)
			if err != nil {
				menu.log.Error("CANT ACCEPT PUBLISH", slog.Any("error", err))
				return err
			}
		}
	case 2:
		switch reqMap[reqID].Type {
		case sign_contract.SignRequest:
			signReqUC := menu.a.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
			signreq, err := signReqUC.Get(reqID)
			if err != nil {
				return err
			}
			err = menu.a.UseCases.SignContractReqUC.Decline(signreq)
			if err != nil {
				return err
			}
		case publish.PubReq:
			pubReqUC := menu.a.UseCases.PublishReqUC.(*usecase2.PublishRequestUseCase)
			pubreq, err := pubReqUC.Get(reqID)
			if err != nil {
				return err
			}
			err = menu.a.UseCases.PublishReqUC.Decline(pubreq)
			if err != nil {
				return err
			}
		}
	default:
		fmt.Printf("Неверный пункт меню")
	}

	return err
}
