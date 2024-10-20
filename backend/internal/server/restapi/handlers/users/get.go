package users

import (
	"context"
	"cookdroogers/app"
	modelsDTO "cookdroogers/internal/server/models"
	"cookdroogers/internal/server/restapi/handlers/common"
	"cookdroogers/internal/server/restapi/operations/users"
	"cookdroogers/internal/server/restapi/session"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"net/http"
)

func getByUserIDHandlerFunc(params users.GetUserByIDParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}

	ctx := context.Background()

	userID := params.UserID

	if userID == 0 {
		return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid userid")
	}

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	user, err := app.Services.UserService.Get(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, userErrors.ErrNoUser):
			return common.ErrorResponse(http.StatusNotFound, "No such user.")
		default:
			return common.ErrorResponse(http.StatusInternalServerError, "Can't get user")
		}
	}

	if user.UserID != authUserID {
		return common.ErrorResponse(http.StatusForbidden, "Forbidden")
	}

	userDTO := modelsDTO.UserDTO{
		UserID:   user.UserID,
		Name:     user.Name,
		Password: user.Password,
		Type:     int64(user.Type),
		Email:    strfmt.Email(user.Email),
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, userDTO)
	})
}

func getUsersHandlerFunc(params users.GetUsersParams, app *app.App) middleware.Responder {
	_, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.AdminUserStr {
		return common.ErrorResponse(http.StatusForbidden, "No rights")
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	usersCD, err := app.Services.UserService.GetForAdmin(ctx)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't get users")
	}

	usersDTO := make([]modelsDTO.UserDTO, len(usersCD))

	for i, userCD := range usersCD {
		usersDTO[i] = modelsDTO.UserDTO{
			UserID:   userCD.UserID,
			Name:     userCD.Name,
			Password: userCD.Password,
			Type:     int64(userCD.Type),
			Email:    strfmt.Email(userCD.Email),
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, usersDTO)
	})
}
