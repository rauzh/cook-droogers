package auth

import (
	"cookdroogers/app"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/auth"
	"github.com/go-openapi/runtime/middleware"
)

var secretKey = []byte("secret-key")

func ConfigureAuthHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.AuthLoginHandler = auth.LoginHandlerFunc(func(params auth.LoginParams) middleware.Responder {
		return loginHandlerFunc(params, app)
	})
	api.AuthRegisterHandler = auth.RegisterHandlerFunc(func(params auth.RegisterParams) middleware.Responder {
		return registerHandlerFunc(params, app)
	})
}
