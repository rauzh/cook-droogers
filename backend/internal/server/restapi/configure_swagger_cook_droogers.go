// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"cookdroogers/app"
	"cookdroogers/config"
	artistsHandlers "cookdroogers/internal/server/restapi/handlers/artists"
	authHandlers "cookdroogers/internal/server/restapi/handlers/auth"
	managersHandlers "cookdroogers/internal/server/restapi/handlers/managers"
	releasesHandlers "cookdroogers/internal/server/restapi/handlers/releases"
	requestsHandlers "cookdroogers/internal/server/restapi/handlers/requests"
	tracksHandlers "cookdroogers/internal/server/restapi/handlers/tracks"
	usersHandlers "cookdroogers/internal/server/restapi/handlers/users"
	"cookdroogers/internal/server/restapi/session"
	"cookdroogers/pkg/logger"
	"crypto/tls"
	"log/slog"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"cookdroogers/internal/server/restapi/operations"
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
		_, _, role, err := session.VerifyToken(token)
		if err != nil {
			return nil, errors.Unauthenticated("invalid token")
		}
		return role, nil
	}

	authHandlers.ConfigureAuthHandlers(&cdApp, api)
	usersHandlers.ConfigureUserHandlers(&cdApp, api)
	tracksHandlers.ConfigureTracksHandlers(&cdApp, api)
	artistsHandlers.ConfigureArtistsHandlers(&cdApp, api)
	managersHandlers.ConfigureManagerHandlers(&cdApp, api)
	releasesHandlers.ConfigureReleasesHandlers(&cdApp, api)
	requestsHandlers.ConfigureRequestsHandlers(&cdApp, api)

	api.GetHeartbeatHandler = operations.GetHeartbeatHandlerFunc(func(params operations.GetHeartbeatParams) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.WriteHeader(http.StatusOK)
			if _, err := rw.Write([]byte("OK")); err != nil {
				_ = errors.New(500, "internal error")
			}
		})
	})

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
	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"HEAD", "PATCH", "PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type", "X-Requested-With", "access_token"},
	})

	return corsHandler.Handler(handler)
}
