package managers

import (
	"context"
	"cookdroogers/app"
	managerErrors "cookdroogers/internal/manager/errors"
	"cookdroogers/internal/server/restapi/handlers/common"
	"cookdroogers/internal/server/restapi/operations/managers"
	"cookdroogers/internal/server/restapi/session"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
)

func addManagersHandlerFunc(params managers.AddManagersParams, app *app.App) middleware.Responder {
	_, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.AdminUserStr {
		return common.ErrorResponse(http.StatusForbidden, "Forbidden")
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	err = app.Transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, userid := range params.ManagersData.UserID {

			if userid == 0 {
				return managerErrors.ErrInvalidID
			}

			user, err := app.Services.UserService.Get(ctx, userid)
			if err != nil {
				return err
			}

			if user.Type == models.ManagerUser {
				return managerErrors.ErrExists
			}
			if user.Type == models.ArtistUser {
				return managerErrors.ErrInvalidRole
			}

			err = app.Services.UserService.UpdateType(ctx, userid, models.ManagerUser)
			if err != nil {
				return err
			}
			man := models.Manager{UserID: userid}

			err = app.Services.ManagerService.Create(ctx, &man)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		switch {
		case errors.Is(err, userErrors.ErrNoUser):
			return common.ErrorResponse(http.StatusNotFound, "No such user")
		case errors.Is(err, managerErrors.ErrInvalidRole):
			return common.ErrorResponse(http.StatusConflict, "Artist user can't be manager")
		case errors.Is(err, managerErrors.ErrExists):
			return common.ErrorResponse(http.StatusConflict, "Manager already exists")
		case errors.Is(err, managerErrors.ErrInvalidID):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid id")
		default:
			common.ErrorResponse(http.StatusInternalServerError, "Can't create manager")
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusCreated)
	})
}
