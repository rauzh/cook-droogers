package tracks

import (
	"cookdroogers/app"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/tracks"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureTracksHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.TracksGetTrackByIDHandler = tracks.GetTrackByIDHandlerFunc(func(params tracks.GetTrackByIDParams, principal interface{}) middleware.Responder {
		return getTrackByIDHandlerFunc(params, app)
	})
}
