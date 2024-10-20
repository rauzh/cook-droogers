package requests

import (
	"context"
	"cookdroogers/app"
	serviceErrors "cookdroogers/internal/errors"
	baseReqErrors "cookdroogers/internal/requests/base/errors"
	"cookdroogers/internal/requests/publish"
	pubReqErrors "cookdroogers/internal/requests/publish/errors"
	"cookdroogers/internal/requests/sign_contract"
	sctErrors "cookdroogers/internal/requests/sign_contract/errors"
	"cookdroogers/internal/server/restapi/handlers/common"
	"cookdroogers/internal/server/restapi/operations/requests"
	"cookdroogers/internal/server/restapi/session"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func signContractHandlerFunc(params requests.SignContractParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.NonMemberUserStr {
		return common.ErrorResponse(http.StatusForbidden, "No rights")
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	nickName := *params.SignRequest.Nickname

	signReq := sign_contract.NewSignContractRequest(authUserID, nickName)

	err = app.UseCases.SignContractReqUC.Apply(ctx, signReq)
	if err != nil {
		switch {
		case errors.Is(err, baseReqErrors.ErrAlreadyClosed):
			return common.ErrorResponse(http.StatusConflict, "Already closed")
		case errors.Is(err, baseReqErrors.ErrNoApplierID):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid applier id")
		case errors.Is(err, baseReqErrors.ErrInvalidType):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid request type")
		case errors.Is(err, sctErrors.ErrNickname):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid nickname")
		default:
			return common.ErrorResponse(http.StatusInternalServerError, "Can't apply sign contract request")
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusCreated)
	})
}

func publishHandlerFunc(params requests.PublishReqParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.ArtistUserStr {
		return common.ErrorResponse(http.StatusForbidden, "No rights")
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	releaseID := params.PublicationInfo.ReleaseID
	expectedDate := time.Time(params.PublicationInfo.ExpectedDate)

	pubReq := publish.NewPublishRequest(authUserID, releaseID, expectedDate)

	err = app.UseCases.PublishReqUC.Apply(ctx, pubReq)
	if err != nil {
		switch {
		case errors.Is(err, pubReqErrors.ErrEndContract):
			return common.ErrorResponse(http.StatusBadRequest, "End of the contract")
		case errors.Is(err, pubReqErrors.ErrNotOwner):
			return common.ErrorResponse(http.StatusForbidden, "No rights: not owner")
		case errors.Is(err, serviceErrors.ErrNoSuchInstance):
			return common.ErrorResponse(http.StatusNotFound, "No such release")
		case errors.Is(err, pubReqErrors.ErrReleaseAlreadyPublished):
			return common.ErrorResponse(http.StatusConflict, "Release already published")
		case errors.Is(err, baseReqErrors.ErrAlreadyClosed):
			return common.ErrorResponse(http.StatusConflict, "Already closed")
		case errors.Is(err, pubReqErrors.ErrNoReleaseID):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid releaseID")
		case errors.Is(err, pubReqErrors.ErrInvalidDate):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid date")
		case errors.Is(err, baseReqErrors.ErrNoApplierID):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid applier id")
		case errors.Is(err, baseReqErrors.ErrInvalidType):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid request type")
		default:
			return common.ErrorResponse(http.StatusInternalServerError, "Can't apply publish request")
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusCreated)
	})
}
