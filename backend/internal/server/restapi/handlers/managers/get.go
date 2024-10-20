package managers

import (
	"context"
	"cookdroogers/app"
	modelsDTO "cookdroogers/internal/server/models"
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

func getManagerByIDHandlerFunc(params managers.GetManagerByIDParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.ManagerUserStr {
		return common.ErrorResponse(http.StatusForbidden, "Forbidden")
	}

	ctx := context.Background()

	managerID := params.ManagerID

	if managerID == 0 {
		return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid managerID")
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

	man, err := app.Services.ManagerService.Get(ctx, managerID)
	if err != nil {
		switch {
		case errors.Is(err, userErrors.ErrNoUser):
			return common.ErrorResponse(http.StatusNotFound, "No such manager")
		default:
			return common.ErrorResponse(http.StatusInternalServerError, "Can't get manager")
		}
	}

	if man.UserID != authUserID {
		return common.ErrorResponse(http.StatusForbidden, "Forbidden")
	}

	managerDTO := modelsDTO.ManagerDTO{
		UserID:    man.UserID,
		ManagerID: man.ManagerID,
		Artists:   man.Artists,
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, managerDTO)
	})
}

func getManagersHandlerFunc(params managers.GetManagersParams, app *app.App) middleware.Responder {
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

	managersCD, err := app.Services.ManagerService.GetForAdmin(ctx)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't get managers")
	}

	managersDTO := make([]modelsDTO.ManagerDTO, len(managersCD))

	for i, man := range managersCD {
		managersDTO[i] = modelsDTO.ManagerDTO{
			UserID:    man.UserID,
			ManagerID: man.ManagerID,
			Artists:   man.Artists,
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, managersDTO)
	})
}
