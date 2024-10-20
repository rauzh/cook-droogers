package releases

import (
	"cookdroogers/app"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/releases"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureReleasesHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.ReleasesGetReleaseHandler = releases.GetReleaseHandlerFunc(func(params releases.GetReleaseParams, principal interface{}) middleware.Responder {
		return getReleasesHandlerFunc(params, app)
	})
	api.ReleasesGetReleaseByIDHandler = releases.GetReleaseByIDHandlerFunc(func(params releases.GetReleaseByIDParams, principal interface{}) middleware.Responder {
		return getReleaseByIDHandlerFunc(params, app)
	})
	api.ReleasesAddReleaseHandler = releases.AddReleaseHandlerFunc(func(params releases.AddReleaseParams, principal interface{}) middleware.Responder {
		return addReleaseHandlerFunc(params, app)
	})
}
