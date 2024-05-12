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

func userLoop(a *app.App, user *models.User, log *slog.Logger) error {

	err := a.Services.UserService.SetRole(models.NonMemberUser)
	if err != nil {
		log.Error("Can't init user menu: can't set user role: ", slog.Any("error", err))
		return err
	}

	startPosition :=
		`
		0 -- выйти
		1 -- подать заявку на вступление
		2 -- посмотреть список заявок
		3 -- посмотреть информацию о лейбле
	Выберите пункт меню: `

	for {
		fmt.Printf("%s", startPosition)

		var action int
		_, _ = fmt.Scanf("%d", &action)

		switch action {
		case 0:
			return ErrEXIT
		case 1:
			err := applySignRequest(a, user)
			if err != nil {
				fmt.Println("Не удается создать заявку по причине ", err)
				log.Error("Can't apply sign request: ", slog.Any("error", err))
			} else {
				fmt.Println("Заявка на вступление успешно подана.")
			}

		case 2:
			err := lookupReqs(a, user)
			if err != nil {
				fmt.Println("Не удается просмотреть заявки по причине ", err)
				log.Error("Can't look up requests: ", slog.Any("error", err))
			}
		case 3:
			printInfo()
		default:
			fmt.Printf("Неверный пункт меню")
		}
	}
}

func applySignRequest(a *app.App, applier *models.User) error {

	fmt.Printf("%s", "Введите никнейм: ")
	var nickname string
	_, _ = fmt.Scanf("%s", &nickname)

	signReq := sign_contract.NewSignContractRequest(applier.UserID, nickname)

	return a.UseCases.SignContractReqUC.Apply(signReq)
}

func lookupReqs(a *app.App, user *models.User) error {

	reqs, err := a.Services.RequestService.GetAllByUserID(user.UserID)
	if err != nil {
		return err
	}

	reqMap := make(map[uint64]base.Request)

	for _, req := range reqs {
		fmt.Printf("\n\t request id:%d\n\t type:%s\n\t status:%s\n\t date:%s\n\t manager:%d\n\t applier:%d\n",
			req.RequestID, req.Type, req.Status, req.Date, req.ManagerID, req.ApplierID)
		reqMap[req.RequestID] = req
	}

	fmt.Printf("\n%s", "Введите id заявки, которую хотите посмотреть (0, если не хотите): ")
	var reqID uint64
	_, _ = fmt.Scanf("%d", &reqID)
	if reqID == 0 {
		return nil
	}

	switch reqMap[reqID].Type {
	case sign_contract.SignRequest:
		signReqUC := a.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)

		signreq, err := signReqUC.Get(reqID)
		if err != nil {
			return err
		}

		fmt.Printf("\n\t nickname: %s \n\t decription: %s\n", signreq.Nickname, signreq.Description)
	case publish.PubReq:
		pubReqUC := a.UseCases.PublishReqUC.(*usecase2.PublishRequestUseCase)

		pubreq, err := pubReqUC.Get(reqID)
		if err != nil {
			return err
		}

		fmt.Printf("\n\t expected date:%s\n\t decription: %s\n\t grade: %d\n\t release id: %d\n",
			pubreq.ExpectedDate, pubreq.Description, pubreq.Grade, pubreq.ReleaseID)
	}
	return nil
}
