package requests

import (
	"context"
	"cookdroogers/app"
	serviceErrors "cookdroogers/internal/errors"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	pubReqUsecase "cookdroogers/internal/requests/publish/usecase"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/usecase"
	modelsDTO "cookdroogers/internal/server/models"
	"cookdroogers/internal/server/restapi/operations/requests"
	"cookdroogers/internal/server/restapi/session"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"net/http"
)

func getRequestsHandlerFunc(params requests.GetRequestsParams, app *app.App) middleware.Responder {
	var err error

	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.ArtistUserStr && role != models.ManagerUserStr && role != models.NonMemberUserStr {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "No rights",
			})
		})
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't set role",
			})
		})
	}

	reqs := make([]base.Request, 0)
	if role == models.ManagerUserStr {
		man, err := app.Services.ManagerService.GetByUserID(ctx, authUserID)
		if err != nil {
			return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't get requests",
				})
			})
		}
		reqs, err = app.Services.RequestService.GetAllByManagerID(man.ManagerID)
	} else {
		reqs, err = app.Services.RequestService.GetAllByUserID(authUserID)
	}
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't get releases",
			})
		})
	}

	reqsDTO := make([]modelsDTO.RequestDTO, len(reqs))

	for i, req := range reqs {
		reqsDTO[i] = modelsDTO.RequestDTO{
			ApplierID: req.ApplierID,
			Date:      strfmt.Date(req.Date),
			ManagerID: req.ManagerID,
			RequestID: req.RequestID,
			Status:    string(req.Status),
			Type:      string(req.Type),
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, reqsDTO)
	})
}

func getRequestByIDHandlerFunc(params requests.GetRequestParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.NonMemberUserStr && role != models.ArtistUserStr && role != models.ManagerUserStr {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Forbidden",
			})
		})
	}

	ctx := context.Background()

	reqID := params.ReqID

	if reqID == 0 {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnprocessableEntity)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Invalid requestID",
			})
		})
	}

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't set role",
			})
		})
	}

	req, err := app.Services.RequestService.GetByID(reqID)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			switch {
			case errors.Is(err, serviceErrors.ErrNoSuchInstance):
				rw.WriteHeader(http.StatusNotFound)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "No such request",
				})
			default:
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't get request",
				})
			}
		})
	}

	if role == models.ManagerUserStr {
		man, err := app.Services.ManagerService.GetByUserID(ctx, authUserID)
		if err != nil {
			return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't get request",
				})
			})
		}
		if man.ManagerID != req.ManagerID {
			return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
				rw.WriteHeader(http.StatusForbidden)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Forbidden",
				})
			})
		}
	}

	if role == models.ArtistUserStr || role == models.NonMemberUserStr {
		if authUserID != req.ApplierID {
			return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
				rw.WriteHeader(http.StatusForbidden)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Forbidden",
				})
			})
		}
	}

	var pubreqDTO *modelsDTO.PublishRequestDTO = nil
	var signreqDTO *modelsDTO.SignRequestDTO = nil

	switch req.Type {
	case publish.PubReq:
		pubReqUC := app.UseCases.PublishReqUC.(*pubReqUsecase.PublishRequestUseCase)
		pubreq, err := pubReqUC.Get(ctx, reqID)
		if err != nil {
			return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't get publish request",
				})
			})
		}
		pubreqDTO = &modelsDTO.PublishRequestDTO{
			Base: &modelsDTO.RequestDTO{
				ApplierID: pubreq.ApplierID,
				Date:      strfmt.Date(pubreq.Date),
				ManagerID: pubreq.ManagerID,
				RequestID: pubreq.RequestID,
				Status:    string(pubreq.Status),
				Type:      string(pubreq.Type),
			},
			Description:  pubreq.Description,
			ExpectedDate: strfmt.Date(pubreq.ExpectedDate),
			Grade:        int64(pubreq.Grade),
			ReleaseID:    pubreq.ReleaseID,
		}

	case sign_contract.SignRequest:
		signReqUC := app.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
		signreq, err := signReqUC.Get(ctx, reqID)
		if err != nil {
			return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't get sign request",
				})
			})
		}

		signreqDTO = &modelsDTO.SignRequestDTO{
			Base: &modelsDTO.RequestDTO{
				ApplierID: signreq.ApplierID,
				Date:      strfmt.Date(signreq.Date),
				ManagerID: signreq.ManagerID,
				RequestID: signreq.RequestID,
				Status:    string(signreq.Status),
				Type:      string(signreq.Type),
			},
			Description: signreq.Description,
			Nickname:    signreq.Nickname,
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		if pubreqDTO != nil {
			_ = p.Produce(rw, pubreqDTO)
		}
		if signreqDTO != nil {
			_ = p.Produce(rw, signreqDTO)
		}
	})
}
