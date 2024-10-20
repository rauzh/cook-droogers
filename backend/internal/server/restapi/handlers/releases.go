package handlers

import (
	"context"
	"cookdroogers/app"
	releaseErrors "cookdroogers/internal/release/errors"
	modelsDTO "cookdroogers/internal/server/models"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/releases"
	"cookdroogers/internal/server/restapi/session"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"net/http"
	"time"
)

func getReleaseByIDHandlerFunc(params releases.GetReleaseByIDParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.ArtistUserStr {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Forbidden",
			})
		})
	}

	ctx := context.Background()

	releaseID := params.ReleaseID

	if releaseID == 0 {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnprocessableEntity)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Invalid releaseID",
			})
		})
	}

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't set role",
			})
		})
	}

	rel, err := app.Services.ReleaseService.Get(ctx, releaseID)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			switch {
			case errors.Is(err, userErrors.ErrNoUser):
				rw.WriteHeader(http.StatusNotFound)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "No such release",
				})
			default:
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't get release",
				})
			}
		})
	}

	artist, err := app.Services.ArtistService.GetByUserID(ctx, authUserID)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't get release",
			})
		})
	}

	if rel.ArtistID != artist.ArtistID {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Forbidden",
			})
		})
	}

	releaseDTO := modelsDTO.ReleaseDTO{
		Title:        rel.Title,
		Status:       string(rel.Status),
		ReleaseID:    rel.ReleaseID,
		DateCreation: strfmt.Date(rel.DateCreation),
		ArtistID:     artist.ArtistID,
		Tracks:       rel.Tracks,
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, releaseDTO)
	})
}

func getReleasesHandlerFunc(params releases.GetReleaseParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.ArtistUserStr {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "No rights",
			})
		})
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't set role",
			})
		})
	}

	artist, err := app.Services.ArtistService.GetByUserID(ctx, authUserID)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't get releases",
			})
		})
	}

	rlss, err := app.Services.ReleaseService.GetAllByArtist(ctx, artist.ArtistID)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't get releases",
			})
		})
	}

	releasesDTO := make([]modelsDTO.ReleaseDTO, len(rlss))

	for i, release := range rlss {
		releasesDTO[i] = modelsDTO.ReleaseDTO{
			Title:        release.Title,
			Status:       string(release.Status),
			ReleaseID:    release.ReleaseID,
			DateCreation: strfmt.Date(release.DateCreation),
			ArtistID:     artist.ArtistID,
			Tracks:       release.Tracks,
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		_ = p.Produce(rw, releasesDTO)
	})
}

func addReleaseHandlerFunc(params releases.AddReleaseParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusUnauthorized)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Auth error",
			})
		})
	}
	if role != models.ArtistUserStr {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusForbidden)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "No rights",
			})
		})
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't set role",
			})
		})
	}

	artist, err := app.Services.ArtistService.GetByUserID(ctx, authUserID)
	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusInternalServerError)
			_ = p.Produce(rw, modelsDTO.LeErrorMessage{
				Error: "Can't get releases",
			})
		})
	}

	release := &models.Release{
		Title:        *params.Release.Title,
		ArtistID:     artist.ArtistID,
		DateCreation: time.Time(*params.Release.Date),
	}

	tracksRaw := params.Release.Tracks
	tracks := make([]*models.Track, len(tracksRaw))

	for i, trackRaw := range tracksRaw {
		track := &models.Track{
			Title:    *trackRaw.Title,
			Duration: *trackRaw.Duration,
			Genre:    *trackRaw.Genre,
			Type:     *trackRaw.Type,
			Artists:  []uint64{artist.ArtistID},
		}
		tracks[i] = track
	}

	err = app.Services.ReleaseService.Create(ctx, release, tracks)

	if err != nil {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			switch {
			case errors.Is(err, releaseErrors.ErrNoTitle):
				rw.WriteHeader(http.StatusUnprocessableEntity)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Invalid title",
				})
			case errors.Is(err, releaseErrors.ErrNoDate):
				rw.WriteHeader(http.StatusUnprocessableEntity)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Invalid date",
				})
			default:
				rw.WriteHeader(http.StatusInternalServerError)
				_ = p.Produce(rw, modelsDTO.LeErrorMessage{
					Error: "Can't create release",
				})
			}
		})
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusCreated)
	})
}

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
