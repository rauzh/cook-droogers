package artists

import (
	"cookdroogers/app"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/artists"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureArtistsHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.ArtistsGetArtistByIDHandler = artists.GetArtistByIDHandlerFunc(func(params artists.GetArtistByIDParams, principal interface{}) middleware.Responder {
		return getArtistByIDHandlerFunc(params, app)
	})
}
