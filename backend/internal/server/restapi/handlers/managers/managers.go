package managers

import (
	"cookdroogers/app"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/managers"
	"github.com/go-openapi/runtime/middleware"
)

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
