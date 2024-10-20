package requests

import (
	"context"
	"cookdroogers/app"
	serviceErrors "cookdroogers/internal/errors"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	pubReqUseCase "cookdroogers/internal/requests/publish/usecase"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/usecase"
	"cookdroogers/internal/server/restapi/handlers/common"
	"cookdroogers/internal/server/restapi/operations/requests"
	"cookdroogers/internal/server/restapi/session"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
)

func acceptHandlerFunc(params requests.AcceptRequestParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.ManagerUserStr {
		return common.ErrorResponse(http.StatusForbidden, "Forbidden")
	}

	ctx := context.Background()

	reqID := params.ReqID

	if reqID == 0 {
		return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid requestID")
	}

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	man, err := app.Services.ManagerService.GetByUserID(ctx, authUserID)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't accept request")
	}

	reqs, err := app.Services.RequestService.GetAllByManagerID(man.ManagerID)
	idx := len(reqs) + 1
	for i, req := range reqs {
		if req.RequestID == reqID {
			idx = i
			break
		}
	}
	if idx == len(reqs)+1 {
		return common.ErrorResponse(http.StatusNotFound, "No such request for this manager")
	}
	if reqs[idx].Status == base.ClosedRequest {
		return common.ErrorResponse(http.StatusConflict, "Already closed")
	}

	switch reqs[idx].Type {
	case publish.PubReq:
		pubReqUC := app.UseCases.PublishReqUC.(*pubReqUseCase.PublishRequestUseCase)
		pubreq, err := pubReqUC.Get(ctx, reqID)
		if err != nil {
			switch {
			case errors.Is(err, serviceErrors.ErrNoSuchInstance):
				return common.ErrorResponse(http.StatusNotFound, "No such publish request")
			}
			return common.ErrorResponse(http.StatusInternalServerError, "Can't accept request")
		}
		err = app.UseCases.PublishReqUC.Accept(ctx, pubreq)
		if err != nil {
			return common.ErrorResponse(http.StatusInternalServerError, "Can't accept request")
		}
	case sign_contract.SignRequest:
		signReqUC := app.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
		signreq, err := signReqUC.Get(ctx, reqID)
		if err != nil {
			switch {
			case errors.Is(err, serviceErrors.ErrNoSuchInstance):
				return common.ErrorResponse(http.StatusNotFound, "No such sign request")
			}
			return common.ErrorResponse(http.StatusInternalServerError, "Can't accept request")
		}
		err = app.UseCases.SignContractReqUC.Accept(ctx, signreq)
		if err != nil {
			return common.ErrorResponse(http.StatusInternalServerError, "Can't accept request")
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
	})
}

func declineHandlerFunc(params requests.DeclineRequestParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.ManagerUserStr {
		return common.ErrorResponse(http.StatusForbidden, "Forbidden")
	}

	ctx := context.Background()

	reqID := params.ReqID

	if reqID == 0 {
		return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid requestID")
	}

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	man, err := app.Services.ManagerService.GetByUserID(ctx, authUserID)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't decline request")
	}

	reqs, err := app.Services.RequestService.GetAllByManagerID(man.ManagerID)
	idx := len(reqs) + 1
	for i, req := range reqs {
		if req.RequestID == reqID {
			idx = i
			break
		}
	}
	if idx == len(reqs)+1 {
		return common.ErrorResponse(http.StatusNotFound, "No such request for this manager")
	}
	if reqs[idx].Status == base.ClosedRequest {
		return common.ErrorResponse(http.StatusConflict, "Already closed")
	}

	switch reqs[idx].Type {
	case publish.PubReq:
		pubReqUC := app.UseCases.PublishReqUC.(*pubReqUseCase.PublishRequestUseCase)
		pubreq, err := pubReqUC.Get(ctx, reqID)
		if err != nil {
			switch {
			case errors.Is(err, serviceErrors.ErrNoSuchInstance):
				return common.ErrorResponse(http.StatusNotFound, "No such publish request")
			}
			return common.ErrorResponse(http.StatusInternalServerError, "Can't decline request")
		}
		err = app.UseCases.PublishReqUC.Decline(ctx, pubreq)
		if err != nil {
			return common.ErrorResponse(http.StatusInternalServerError, "Can't decline request")
		}
	case sign_contract.SignRequest:
		signReqUC := app.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
		signreq, err := signReqUC.Get(ctx, reqID)
		if err != nil {
			switch {
			case errors.Is(err, serviceErrors.ErrNoSuchInstance):
				return common.ErrorResponse(http.StatusNotFound, "No such sign request")
			}
			return common.ErrorResponse(http.StatusInternalServerError, "Can't decline request")
		}
		err = app.UseCases.SignContractReqUC.Decline(ctx, signreq)
		if err != nil {
			return common.ErrorResponse(http.StatusInternalServerError, "Can't decline request")
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
	})
}
