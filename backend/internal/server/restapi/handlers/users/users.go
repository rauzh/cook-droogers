package users

import (
	"cookdroogers/app"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/users"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureUserHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams, principal interface{}) middleware.Responder {
		return getByUserIDHandlerFunc(params, app)
	})
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams, principal interface{}) middleware.Responder {
		return getUsersHandlerFunc(params, app)
	})
}
