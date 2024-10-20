package handlers

import (
	"context"
	"cookdroogers/app"
	modelsDTO "cookdroogers/internal/server/models"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/tracks"
	"cookdroogers/internal/server/restapi/session"
	trackErrors "cookdroogers/internal/track/errors"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
)

func getTrackByIDHandlerFunc(params tracks.GetTrackByIDParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.Error(http.StatusUnauthorized, "Auth error")
	}
	if role != models.ArtistUserStr && role != models.ManagerUserStr {
		return middleware.Error(http.StatusForbidden, "No rights")
	}

	ctx := context.Background()

	trackID := params.TrackID

	if trackID == 0 {
		return middleware.Error(http.StatusUnprocessableEntity, "Invalid trackid")
	}

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Can't set role")
	}

	track, err := app.Services.TrackService.Get(ctx, trackID)
	if err != nil {
		switch {
		case errors.Is(err, trackErrors.ErrNoTrack):
			return middleware.Error(http.StatusNotFound, "No such track")
		default:
			return middleware.Error(http.StatusInternalServerError, "Can't get track")
		}
	}

	if len(track.Artists) < 1 {
		return middleware.Error(http.StatusForbidden, "Forbidden")
	}

	if role == models.ArtistUserStr {
		artist, err := app.Services.ArtistService.GetByUserID(ctx, authUserID)
		if err != nil {
			return middleware.Error(http.StatusInternalServerError, "Can't get track")
		}
		flag := false
		for _, trackArtistID := range track.Artists {
			if trackArtistID == artist.ArtistID {
				flag = true
				break
			}
		}
		if !flag {
			return middleware.Error(http.StatusForbidden, "Forbidden")
		}
	}
	if role == models.ManagerUserStr {
		manager, err := app.Services.ManagerService.GetByUserID(ctx, authUserID)
		if err != nil {
			return middleware.Error(http.StatusInternalServerError, "Can't get track")
		}
		flag := false
		for _, artistid := range manager.Artists {
			for _, trackArtistID := range track.Artists {
				if trackArtistID == artistid {
					flag = true
					break
				}
			}
		}
		if !flag {
			return middleware.Error(http.StatusForbidden, "Forbidden")
		}
	}

	trackDTO := modelsDTO.TrackDTO{
		TrackID:  track.TrackID,
		Type:     &track.Type,
		Title:    &track.Title,
		Genre:    &track.Genre,
		Duration: &track.Duration,
		Artists:  track.Artists,
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, trackDTO)
	})
}

func ConfigureTracksHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.TracksGetTrackByIDHandler = tracks.GetTrackByIDHandlerFunc(func(params tracks.GetTrackByIDParams, principal interface{}) middleware.Responder {
		return getTrackByIDHandlerFunc(params, app)
	})
}
