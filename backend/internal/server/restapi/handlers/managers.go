package handlers

import (
	"context"
	"cookdroogers/app"
	managerErrors "cookdroogers/internal/manager/errors"
	modelsDTO "cookdroogers/internal/server/models"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/managers"
	"cookdroogers/internal/server/restapi/session"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func getManagerByIDHandlerFunc(params managers.GetManagerByIDParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.ManagerUserStr {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Forbidden",
			})
		})
	}

	ctx := context.Background()

	managerID := params.ManagerID

	if managerID == 0 {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnprocessableEntity)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Invalid managerID",
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

	man, err := app.Services.ManagerService.Get(ctx, managerID)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			switch {
			case errors.Is(err, userErrors.ErrNoUser):
				rw.WriteHeader(http.StatusNotFound)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "No such manager",
				})
			default:
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't get manager",
				})
			}
		})
	}

	if man.UserID != authUserID {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Forbidden",
			})
		})
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
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.AdminUserStr {
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

	managersCD, err := app.Services.ManagerService.GetForAdmin(ctx)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't get manager",
			})
		})
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

func addManagersHandlerFunc(params managers.AddManagersParams, app *app.App) middleware.Responder {
	_, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.AdminUserStr {
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
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			switch {
			case errors.Is(err, userErrors.ErrNoUser):
				rw.WriteHeader(http.StatusNotFound)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "No such user",
				})
			case errors.Is(err, managerErrors.ErrInvalidRole):
				rw.WriteHeader(http.StatusConflict)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Artist user can't be manager",
				})
			case errors.Is(err, managerErrors.ErrExists):
				rw.WriteHeader(http.StatusConflict)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Manager already exists",
				})
			case errors.Is(err, managerErrors.ErrInvalidID):
				rw.WriteHeader(http.StatusUnprocessableEntity)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Invalid id",
				})
			default:
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't create manager",
				})
			}
		})
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusCreated)
	})
}

func ConfigureManagerHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.ManagersGetManagerByIDHandler = managers.GetManagerByIDHandlerFunc(func(params managers.GetManagerByIDParams, principal interface{}) middleware.Responder {
		return getManagerByIDHandlerFunc(params, app)
	})
	api.ManagersGetManagersHandler = managers.GetManagersHandlerFunc(func(params managers.GetManagersParams, principal interface{}) middleware.Responder {
		return getManagersHandlerFunc(params, app)
	})
	api.ManagersAddManagersHandler = managers.AddManagersHandlerFunc(func(params managers.AddManagersParams, principal interface{}) middleware.Responder {
		return addManagersHandlerFunc(params, app)
	})
}
