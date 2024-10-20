package handlers

import (
	"context"
	"cookdroogers/app"
	modelsDTO "cookdroogers/internal/server/models"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/artists"
	"cookdroogers/internal/server/restapi/session"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"net/http"
)

func getArtistByIDHandlerFunc(params artists.GetArtistByIDParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.Error(http.StatusUnauthorized, "Auth error")
	}
	if role != models.ArtistUserStr && role != models.ManagerUserStr {
		return middleware.Error(http.StatusForbidden, "No rights")
	}

	ctx := context.Background()

	id := params.ID

	var userIdFlag bool
	if params.ByUserID == nil {
		userIdFlag = false
	} else {
		userIdFlag = *params.ByUserID
	}

	if id == 0 {
		return middleware.Error(http.StatusUnprocessableEntity, "Invalid id")
	}

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Can't set role")
	}

	var artist *models.Artist
	if userIdFlag {
		artist, err = app.Services.ArtistService.GetByUserID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, userErrors.ErrNoUser):
				return middleware.Error(http.StatusNotFound, "No such artist")
			default:
				return middleware.Error(http.StatusInternalServerError, "Can't get artist")
			}
		}
	} else {
		artist, err = app.Services.ArtistService.Get(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, userErrors.ErrNoUser):
				return middleware.Error(http.StatusNotFound, "No such artist")
			default:
				return middleware.Error(http.StatusInternalServerError, "Can't get artist")
			}
		}
	}

	if role == models.ArtistUserStr {
		if artist.UserID != authUserID {
			return middleware.Error(http.StatusForbidden, "Forbidden")
		}
	}

	if role == models.ManagerUserStr {
		manager, err := app.Services.ManagerService.GetByUserID(ctx, authUserID)
		if err != nil {
			return middleware.Error(http.StatusInternalServerError, "Can't get artist")
		}
		flag := false
		for _, artistID := range manager.Artists {
			if artistID == artist.ArtistID {
				flag = true
				break
			}
		}
		if !flag {
			return middleware.Error(http.StatusForbidden, "Forbidden")
		}
	}

	artistDTO := modelsDTO.ArtistDTO{
		ArtistID:     artist.ArtistID,
		Activity:     artist.Activity,
		ContractTerm: strfmt.Date(artist.ContractTerm),
		ManagerID:    artist.ManagerID,
		Nickname:     artist.Nickname,
		UserID:       artist.UserID,
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, artistDTO)
	})
}

func ConfigureArtistsHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.ArtistsGetArtistByIDHandler = artists.GetArtistByIDHandlerFunc(func(params artists.GetArtistByIDParams, principal interface{}) middleware.Responder {
		return getArtistByIDHandlerFunc(params, app)
	})
}
