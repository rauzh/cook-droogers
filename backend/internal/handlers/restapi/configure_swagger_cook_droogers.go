// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"cookdroogers/app"
	"cookdroogers/config"
	"cookdroogers/internal/handlers"
	modelsDTO "cookdroogers/internal/handlers/models"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	usecase2 "cookdroogers/internal/requests/publish/usecase"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/usecase"
	"cookdroogers/models"
	"cookdroogers/pkg/logger"
	"crypto/tls"
	"database/sql"
	goerrors "errors"
	"github.com/go-openapi/strfmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"cookdroogers/internal/handlers/restapi/operations"
	"cookdroogers/internal/handlers/restapi/operations/admin"
	"cookdroogers/internal/handlers/restapi/operations/artist"
	"cookdroogers/internal/handlers/restapi/operations/guest"
	"cookdroogers/internal/handlers/restapi/operations/manager"
	"cookdroogers/internal/handlers/restapi/operations/non_member"
)

//go:generate swagger generate server --target ../../handlers --name SwaggerCookDroogers --spec ../../../swagger-api/swagger.yml --principal interface{}

func configureFlags(api *operations.SwaggerCookDroogersAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerCookDroogersAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	loggerFactory := &logger.LoggerFactorySlog{}

	ctx := context.Background()
	log := loggerFactory.Logger(ctx)

	appConfig := config.ParseConfig()
	if appConfig == nil {
		log.Error("Failed to parse config")
		panic("Failed to parse config")
	}

	cdApp := app.App{Config: appConfig}

	err := cdApp.Init(log)
	if err != nil {
		log.Error("Failed to initialize app: ", slog.Any("error", err))
		panic("Failed to initialize app")
	}

	api.Logger = log.Info

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	handlers.Users = make(map[string]*models.User)

	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = func(username string, pass string) (interface{}, error) {

		user, err := cdApp.Services.UserService.Login(username, pass)
		if err != nil {
			if goerrors.Is(err, sql.ErrNoRows) {
				return nil, errors.New(403, "No such user")
			}
			if strings.Contains(err.Error(), "invalid password") {
				return nil, errors.New(403, "Invalid password")
			}
			return nil, errors.New(500, "Internal error")
		}

		handlers.Users[username] = user

		return user, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.GetHeartbeatHandler = operations.GetHeartbeatHandlerFunc(func(params operations.GetHeartbeatParams) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
			if _, err := rw.Write([]byte("OK")); err != nil {
				errors.New(500, "Internal error")
			}
		})
	})

	api.ManagerAcceptRequestHandler = manager.AcceptRequestHandlerFunc(func(params manager.AcceptRequestParams, principal interface{}) middleware.Responder {

		mngr, err := handlers.LoginManager(params.HTTPRequest.Header.Get("authorization"), &cdApp)
		if err != nil {
			return middleware.Error(400, err.Error())
		}

		reqs, err := cdApp.Services.RequestService.GetAllByManagerID(mngr.ManagerID)
		if err != nil {
			return middleware.Error(500, "Can't get requests")
		}

		reqMap := make(map[uint64]base.Request)

		for _, req := range reqs {
			reqMap[req.RequestID] = req
		}

		switch reqMap[params.ReqID].Type {
		case publish.PubReq:
			pubReqUC := cdApp.UseCases.PublishReqUC.(*usecase2.PublishRequestUseCase)
			pubreq, err := pubReqUC.Get(params.ReqID)
			if err != nil {
				return middleware.Error(500, "Can't get request")
			}
			err = cdApp.UseCases.PublishReqUC.Accept(pubreq)
			if err != nil {
				log.Error("CANT ACCEPT PUBLISH", slog.Any("error", err))
				return middleware.Error(500, "Can't accept request")
			}
		case sign_contract.SignRequest:
			signReqUC := cdApp.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
			signreq, err := signReqUC.Get(params.ReqID)
			if err != nil {
				return middleware.Error(500, "Can't get request")
			}
			err = cdApp.UseCases.SignContractReqUC.Accept(signreq)
			if err != nil {
				log.Error("CANT ACCEPT SIGN CONTRACT", slog.Any("error", err))
				return middleware.Error(500, "Can't accept request")
			}
		}

		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
		})
	})

	if api.AdminAddManagerHandler == nil {
		api.AdminAddManagerHandler = admin.AddManagerHandlerFunc(func(params admin.AddManagerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.AddManager has not yet been implemented")
		})
	}

	api.ArtistAddReleaseHandler = artist.AddReleaseHandlerFunc(func(params artist.AddReleaseParams, principal interface{}) middleware.Responder {

		artistUser, err := handlers.LoginArtist(params.HTTPRequest.Header.Get("authorization"), &cdApp)
		if err != nil {
			return middleware.Error(400, err.Error())
		}

		release := &models.Release{
			Title:        params.Title,
			ArtistID:     artistUser.ArtistID,
			DateCreation: time.Time(params.Date),
		}

		tracksRaw := params.Tracks

		tracks := make([]*models.Track, len(tracksRaw))

		for i, trackRaw := range tracksRaw {
			track := &models.Track{
				Title:    *trackRaw.Title,
				Duration: *trackRaw.Duration,
				Genre:    *trackRaw.Genre,
				Type:     *trackRaw.Type,
				Artists:  []uint64{artistUser.ArtistID},
			}
			tracks[i] = track
		}

		err = cdApp.Services.ReleaseService.Create(release, tracks)
		if err != nil {
			return middleware.Error(500, "can't create release")
		}

		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
		})
	})

	api.ManagerDeclineRequestHandler = manager.DeclineRequestHandlerFunc(func(params manager.DeclineRequestParams, principal interface{}) middleware.Responder {
		mngr, err := handlers.LoginManager(params.HTTPRequest.Header.Get("authorization"), &cdApp)
		if err != nil {
			return middleware.Error(400, err.Error())
		}

		reqs, err := cdApp.Services.RequestService.GetAllByManagerID(mngr.ManagerID)
		if err != nil {
			return middleware.Error(500, "Can't get requests")
		}

		reqMap := make(map[uint64]base.Request)

		for _, req := range reqs {
			reqMap[req.RequestID] = req
		}

		switch reqMap[params.ReqID].Type {
		case publish.PubReq:
			pubReqUC := cdApp.UseCases.PublishReqUC.(*usecase2.PublishRequestUseCase)
			pubreq, err := pubReqUC.Get(params.ReqID)
			if err != nil {
				return middleware.Error(500, "Can't get request")
			}
			err = cdApp.UseCases.PublishReqUC.Decline(pubreq)
			if err != nil {
				log.Error("CANT DECLINE PUBLISH", slog.Any("error", err))
				return middleware.Error(500, "Can't decline request")
			}
		case sign_contract.SignRequest:
			signReqUC := cdApp.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
			signreq, err := signReqUC.Get(params.ReqID)
			if err != nil {
				return middleware.Error(500, "Can't get request")
			}
			err = cdApp.UseCases.SignContractReqUC.Decline(signreq)
			if err != nil {
				log.Error("CANT DECLINE SIGN CONTRACT", slog.Any("error", err))
				return middleware.Error(500, "Can't decline request")
			}
		}

		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
		})
	})

	api.ManagerFetchStatsHandler = manager.FetchStatsHandlerFunc(func(params manager.FetchStatsParams, principal interface{}) middleware.Responder {
		mngr, err := handlers.LoginManager(params.HTTPRequest.Header.Get("authorization"), &cdApp)
		if err != nil {
			return middleware.Error(400, err.Error())
		}

		for _, artistID := range mngr.Artists {
			releases, err := cdApp.Services.ReleaseService.GetAllByArtist(artistID)
			if err != nil {
				return middleware.Error(500, "can't fetch stats")
			}
			for _, release := range releases {
				err = cdApp.Services.StatService.FetchByRelease(&release)
				if err != nil {
					return middleware.Error(500, "can't fetch stats")
				}
			}
		}

		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
		})
	})

	if api.AdminGetManagersHandler == nil {
		api.AdminGetManagersHandler = admin.GetManagersHandlerFunc(func(params admin.GetManagersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.GetManagers has not yet been implemented")
		})
	}

	api.ArtistGetReleaseHandler = artist.GetReleaseHandlerFunc(func(params artist.GetReleaseParams, principal interface{}) middleware.Responder {

		artistUser, err := handlers.LoginArtist(params.HTTPRequest.Header.Get("authorization"), &cdApp)
		if err != nil {
			return middleware.Error(400, err.Error())
		}

		releases, err := cdApp.Services.ReleaseService.GetAllByArtist(artistUser.ArtistID)
		if err != nil {
			return middleware.Error(400, "can't get releases")
		}

		releasesDTO := make([]modelsDTO.ReleaseDTO, len(releases))

		for i, release := range releases {

			releasesDTO[i] = modelsDTO.ReleaseDTO{
				Title:        release.Title,
				Status:       string(release.Status),
				ReleaseID:    release.ReleaseID,
				DateCreation: strfmt.Date(release.DateCreation),
				ArtistID:     artistUser.ArtistID,
				Tracks:       release.Tracks,
			}
		}

		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
			_ = p.Produce(rw, releasesDTO)
		})
	})

	api.NonMemberGetRequestHandler = non_member.GetRequestHandlerFunc(func(params non_member.GetRequestParams, principal interface{}) middleware.Responder {

		user, err := handlers.LoginNonMember(params.HTTPRequest.Header.Get("authorization"), &cdApp)
		if err != nil {
			return middleware.Error(400, err.Error())
		}

		req, err := cdApp.Services.RequestService.GetByID(params.ReqID)
		if err != nil {
			return middleware.Error(500, "Can't get request")
		}

		if req.ApplierID != user.UserID {
			mngr, err := cdApp.Services.ManagerService.GetByUserID(user.UserID)
			if err != nil || mngr.ManagerID != req.ManagerID {
				return middleware.Error(403, "you have no rights to see this request")
			}
		}

		var pubreqDTO *modelsDTO.PublishRequestDTO = nil
		var signreqDTO *modelsDTO.SignRequestDTO = nil

		switch req.Type {
		case publish.PubReq:
			pubReqUC := cdApp.UseCases.PublishReqUC.(*usecase2.PublishRequestUseCase)
			pubreq, err := pubReqUC.Get(params.ReqID)
			if err != nil {
				return middleware.Error(500, "Can't get request")
			}
			pubreqDTO = &modelsDTO.PublishRequestDTO{
				Base: &modelsDTO.RequestDTO{
					ApplierID: pubreq.ApplierID,
					Date:      strfmt.Date(pubreq.Date),
					ManagerID: pubreq.ManagerID,
					RequestID: pubreq.RequestID,
					Status:    string(pubreq.Status),
					Type:      string(pubreq.Type),
				},
				Description:  pubreq.Description,
				ExpectedDate: strfmt.Date(pubreq.ExpectedDate),
				Grade:        int64(pubreq.Grade),
				ReleaseID:    pubreq.ReleaseID,
			}

		case sign_contract.SignRequest:
			signReqUC := cdApp.UseCases.SignContractReqUC.(*usecase.SignContractRequestUseCase)
			signreq, err := signReqUC.Get(params.ReqID)
			if err != nil {
				return middleware.Error(500, "Can't get request")
			}

			signreqDTO = &modelsDTO.SignRequestDTO{
				Base: &modelsDTO.RequestDTO{
					ApplierID: signreq.ApplierID,
					Date:      strfmt.Date(signreq.Date),
					ManagerID: signreq.ManagerID,
					RequestID: signreq.RequestID,
					Status:    string(signreq.Status),
					Type:      string(signreq.Type),
				},
				Description: signreq.Description,
				Nickname:    signreq.Nickname,
			}
		}

		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
			if pubreqDTO != nil {
				_ = p.Produce(rw, pubreqDTO)
			}
			if signreqDTO != nil {
				_ = p.Produce(rw, signreqDTO)
			}
		})
	})

	api.NonMemberGetRequestsHandler = non_member.GetRequestsHandlerFunc(func(params non_member.GetRequestsParams, principal interface{}) middleware.Responder {

		user, err := handlers.LoginNonMember(params.HTTPRequest.Header.Get("authorization"), &cdApp)
		if err != nil {
			return middleware.Error(400, err.Error())
		}

		reqs := make([]base.Request, 0)

		if user.Type == models.ManagerUser {
			mngr, err := cdApp.Services.ManagerService.GetByUserID(user.UserID)
			if err != nil {
				return middleware.Error(500, "can't get manager info")
			}
			reqs, err = cdApp.Services.RequestService.GetAllByManagerID(mngr.ManagerID)
			if err != nil {
				return middleware.Error(500, "can't get manager reqs")
			}
		} else {
			reqs, err = cdApp.Services.RequestService.GetAllByUserID(user.UserID)
			if err != nil {
				return middleware.Error(500, "can't get user reqs")
			}
		}

		reqsDTO := make([]modelsDTO.RequestDTO, len(reqs))

		for i, req := range reqs {
			reqsDTO[i] = modelsDTO.RequestDTO{
				ApplierID: req.ApplierID,
				Date:      strfmt.Date(req.Date),
				ManagerID: req.ManagerID,
				RequestID: req.RequestID,
				Status:    string(req.Status),
				Type:      string(req.Type),
			}
		}

		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
			_ = p.Produce(rw, reqsDTO)
		})
	})

	if api.ArtistGetStatsHandler == nil {
		api.ArtistGetStatsHandler = artist.GetStatsHandlerFunc(func(params artist.GetStatsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation artist.GetStats has not yet been implemented")
		})
	}
	if api.AdminGetUsersHandler == nil {
		api.AdminGetUsersHandler = admin.GetUsersHandlerFunc(func(params admin.GetUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.GetUsers has not yet been implemented")
		})
	}
	if api.ArtistPublishReqHandler == nil {
		api.ArtistPublishReqHandler = artist.PublishReqHandlerFunc(func(params artist.PublishReqParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation artist.PublishReq has not yet been implemented")
		})
	}
	if api.GuestRegisterHandler == nil {
		api.GuestRegisterHandler = guest.RegisterHandlerFunc(func(params guest.RegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation guest.Register has not yet been implemented")
		})
	}
	if api.NonMemberSignContractHandler == nil {
		api.NonMemberSignContractHandler = non_member.SignContractHandlerFunc(func(params non_member.SignContractParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation non_member.SignContract has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
