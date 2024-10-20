package releases

import (
	"context"
	"cookdroogers/app"
	releaseErrors "cookdroogers/internal/release/errors"
	"cookdroogers/internal/server/restapi/handlers/common"
	"cookdroogers/internal/server/restapi/operations/releases"
	"cookdroogers/internal/server/restapi/session"
	"cookdroogers/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func addReleaseHandlerFunc(params releases.AddReleaseParams, app *app.App) middleware.Responder {
	authUserID, _, role, err := session.GetAuthenticatedUser(params.HTTPRequest)
	if err != nil {
		return common.ErrorResponse(http.StatusUnauthorized, "Auth error")
	}
	if role != models.ArtistUserStr {
		return common.ErrorResponse(http.StatusForbidden, "Forbidden")
	}

	ctx := context.Background()

	err = app.Services.UserService.SetRole(ctx, models.UserTypeStrToEnum(role))
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't set role")
	}

	artist, err := app.Services.ArtistService.GetByUserID(ctx, authUserID)
	if err != nil {
		return common.ErrorResponse(http.StatusInternalServerError, "Can't create release")
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
		switch {
		case errors.Is(err, releaseErrors.ErrNoTitle):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid title")
		case errors.Is(err, releaseErrors.ErrNoDate):
			return common.ErrorResponse(http.StatusUnprocessableEntity, "Invalid date")
		default:
			return common.ErrorResponse(http.StatusInternalServerError, "Can't create release")
		}
	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		rw.WriteHeader(http.StatusCreated)
	})
}
