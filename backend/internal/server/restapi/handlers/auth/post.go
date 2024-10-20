package auth

import (
	"context"
	"cookdroogers/app"
	models2 "cookdroogers/internal/server/models"
	"cookdroogers/internal/server/restapi/handlers/common"
	"cookdroogers/internal/server/restapi/operations/auth"
	"cookdroogers/internal/server/restapi/session"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
)

func registerHandlerFunc(params auth.RegisterParams, cdApp *app.App) middleware.Responder {
	password := params.UserData.Password
	name := params.UserData.Username
	email := params.UserData.Email

	user := models.User{
		Password: *password,
		Name:     *name,
		Email:    email.String(),
		Type:     models.NonMemberUser,
	}

	ctx := context.Background()
	err := cdApp.Services.UserService.SetRole(ctx, user.Type)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	err = cdApp.Services.UserService.Create(ctx, &user)
	if err != nil {
		switch {
		case errors.Is(err, userErrors.ErrInvalidEmail):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid email.")
		case errors.Is(err, userErrors.ErrInvalidName):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid username. It can't be empty string.")
		case errors.Is(err, userErrors.ErrInvalidPassword):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid password. It can't be empty string.")
		case errors.Is(err, userErrors.ErrAlreadyTaken):
			return common.ErrorResponse(http.StatusConflict, "User already exists")
		default:
			return common.ErrorResponse(http.StatusInternalServerError, "Can't create user")
		}
	}

	tokenString, err := session.CreateToken(user.UserID, user.Email, models.NonMemberUserStr)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Could not create token")
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusCreated)
		_ = p.Produce(rw, models2.AccessTokenDTO{
			AccessKey: tokenString,
		})
	})
}

func loginHandlerFunc(params auth.LoginParams, cdApp *app.App) middleware.Responder {
	email := params.LoginData.Email.String()
	password := params.LoginData.Password.String()

	ctx := context.Background()

	if session.CheckAdmin(email, password, cdApp) {
		err := session.LoginAdmin(ctx, cdApp)
		if err != nil {
			return common.ErrorResponse(http.StatusInternalServerError, "Can't authorize")
		}
		tokenString, err := session.CreateToken(0, email, models.AdminUserStr)
		if err != nil {
			return common.ErrorResponse(http.StatusInternalServerError, "Could not create token")
		}
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
			_ = p.Produce(rw, models2.AccessTokenDTO{
				AccessKey: tokenString,
			})
		})
	}

	user, err := cdApp.Services.UserService.Login(ctx, email, password)
	if err != nil {
		switch {
		case errors.Is(err, userErrors.ErrInvalidEmail):
			return common.ErrorResponse(http.StatusNotFound, "No such user.")
		case errors.Is(err, userErrors.ErrInvalidPassword):
			return common.ErrorResponse(http.StatusUnauthorized, "Auth error: Invalid password.")
		default:
			return common.ErrorResponse(http.StatusInternalServerError, "Can't authorize")
		}
	}

	err = cdApp.Services.UserService.SetRole(ctx, user.Type)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	var role string
	switch user.Type {
	case models.ArtistUser:
		role = models.ArtistUserStr
	case models.ManagerUser:
		role = models.ManagerUserStr
	case models.NonMemberUser:
		role = models.NonMemberUserStr
	case models.AdminUser:
		role = models.AdminUserStr
	}

	tokenString, err := session.CreateToken(user.UserID, email, role)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Could not create token")
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, models2.AccessTokenDTO{
			AccessKey: tokenString,
		})
	})
}
