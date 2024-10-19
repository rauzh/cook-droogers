// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"cookdroogers/app"
	"cookdroogers/config"
	"cookdroogers/internal/server/restapi/handlers"
	"cookdroogers/internal/server/restapi/session"
	"cookdroogers/pkg/logger"
	"crypto/tls"
	"log/slog"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/artist"
	"cookdroogers/internal/server/restapi/operations/manager"
	"cookdroogers/internal/server/restapi/operations/releases"
	"cookdroogers/internal/server/restapi/operations/requests"
)

//go:generate swagger generate server --target ../../server --name SwaggerCookDroogers --spec ../../../swagger-api/swagger.yml --principal interface{}

func configureFlags(api *operations.SwaggerCookDroogersAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerCookDroogersAPI) http.Handler {

	api.ServeError = errors.ServeError
	api.UseSwaggerUI()
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	ctx := context.Background()

	loggerFactory := &logger.LoggerFactorySlog{}
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

	// Applies when the "access_token" header is set
	api.JWTAuthAuth = func(token string) (interface{}, error) {
		_, role, err := session.VerifyToken(token)
		if err != nil {
			return nil, errors.Unauthenticated("invalid token")
		}
		return role, nil
	}

	handlers.ConfigureAuthHandlers(&cdApp, api)
	handlers.ConfigureUserHandlers(&cdApp, api)
	handlers.ConfigureTracksHandlers(&cdApp, api)

	api.GetHeartbeatHandler = operations.GetHeartbeatHandlerFunc(func(params operations.GetHeartbeatParams) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
			if _, err := rw.Write([]byte("OK")); err != nil {
				_ = errors.New(500, "internal error")
			}
		})
	})

	if api.RequestsAcceptRequestHandler == nil {
		api.RequestsAcceptRequestHandler = requests.AcceptRequestHandlerFunc(func(params requests.AcceptRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.AcceptRequest has not yet been implemented")
		})
	}
	if api.ManagerAddManagerHandler == nil {
		api.ManagerAddManagerHandler = manager.AddManagerHandlerFunc(func(params manager.AddManagerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.AddManager has not yet been implemented")
		})
	}
	if api.ReleasesAddReleaseHandler == nil {
		api.ReleasesAddReleaseHandler = releases.AddReleaseHandlerFunc(func(params releases.AddReleaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation releases.AddRelease has not yet been implemented")
		})
	}
	if api.RequestsDeclineRequestHandler == nil {
		api.RequestsDeclineRequestHandler = requests.DeclineRequestHandlerFunc(func(params requests.DeclineRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.DeclineRequest has not yet been implemented")
		})
	}
	if api.ArtistGetArtistByIDHandler == nil {
		api.ArtistGetArtistByIDHandler = artist.GetArtistByIDHandlerFunc(func(params artist.GetArtistByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation artist.GetArtistByID has not yet been implemented")
		})
	}
	if api.ManagerGetManagerByIDHandler == nil {
		api.ManagerGetManagerByIDHandler = manager.GetManagerByIDHandlerFunc(func(params manager.GetManagerByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.GetManagerByID has not yet been implemented")
		})
	}
	if api.ManagerGetManagersHandler == nil {
		api.ManagerGetManagersHandler = manager.GetManagersHandlerFunc(func(params manager.GetManagersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.GetManagers has not yet been implemented")
		})
	}
	if api.ReleasesGetReleaseHandler == nil {
		api.ReleasesGetReleaseHandler = releases.GetReleaseHandlerFunc(func(params releases.GetReleaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation releases.GetRelease has not yet been implemented")
		})
	}
	if api.ReleasesGetReleaseByIDHandler == nil {
		api.ReleasesGetReleaseByIDHandler = releases.GetReleaseByIDHandlerFunc(func(params releases.GetReleaseByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation releases.GetReleaseByID has not yet been implemented")
		})
	}
	if api.RequestsGetRequestHandler == nil {
		api.RequestsGetRequestHandler = requests.GetRequestHandlerFunc(func(params requests.GetRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.GetRequest has not yet been implemented")
		})
	}
	if api.RequestsGetRequestsHandler == nil {
		api.RequestsGetRequestsHandler = requests.GetRequestsHandlerFunc(func(params requests.GetRequestsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.GetRequests has not yet been implemented")
		})
	}

	if api.RequestsPublishReqHandler == nil {
		api.RequestsPublishReqHandler = requests.PublishReqHandlerFunc(func(params requests.PublishReqParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.PublishReq has not yet been implemented")
		})
	}

	if api.RequestsSignContractHandler == nil {
		api.RequestsSignContractHandler = requests.SignContractHandlerFunc(func(params requests.SignContractParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.SignContract has not yet been implemented")
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
