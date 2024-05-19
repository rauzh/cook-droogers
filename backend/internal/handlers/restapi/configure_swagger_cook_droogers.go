// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

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

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	if api.BasicAuthAuth == nil {
		api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
			return nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.GetHeartbeatHandler == nil {
		api.GetHeartbeatHandler = operations.GetHeartbeatHandlerFunc(func(params operations.GetHeartbeatParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHeartbeat has not yet been implemented")
		})
	}
	if api.ManagerAcceptRequestHandler == nil {
		api.ManagerAcceptRequestHandler = manager.AcceptRequestHandlerFunc(func(params manager.AcceptRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.AcceptRequest has not yet been implemented")
		})
	}
	if api.AdminAddManagerHandler == nil {
		api.AdminAddManagerHandler = admin.AddManagerHandlerFunc(func(params admin.AddManagerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.AddManager has not yet been implemented")
		})
	}
	if api.ArtistAddReleaseHandler == nil {
		api.ArtistAddReleaseHandler = artist.AddReleaseHandlerFunc(func(params artist.AddReleaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation artist.AddRelease has not yet been implemented")
		})
	}
	if api.ManagerDeclineRequestHandler == nil {
		api.ManagerDeclineRequestHandler = manager.DeclineRequestHandlerFunc(func(params manager.DeclineRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.DeclineRequest has not yet been implemented")
		})
	}
	if api.ManagerFetchStatsHandler == nil {
		api.ManagerFetchStatsHandler = manager.FetchStatsHandlerFunc(func(params manager.FetchStatsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.FetchStats has not yet been implemented")
		})
	}
	if api.AdminGetManagersHandler == nil {
		api.AdminGetManagersHandler = admin.GetManagersHandlerFunc(func(params admin.GetManagersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.GetManagers has not yet been implemented")
		})
	}
	if api.ArtistGetReleaseHandler == nil {
		api.ArtistGetReleaseHandler = artist.GetReleaseHandlerFunc(func(params artist.GetReleaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation artist.GetRelease has not yet been implemented")
		})
	}
	if api.NonMemberGetRequestHandler == nil {
		api.NonMemberGetRequestHandler = non_member.GetRequestHandlerFunc(func(params non_member.GetRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation non_member.GetRequest has not yet been implemented")
		})
	}
	if api.NonMemberGetRequestsHandler == nil {
		api.NonMemberGetRequestsHandler = non_member.GetRequestsHandlerFunc(func(params non_member.GetRequestsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation non_member.GetRequests has not yet been implemented")
		})
	}
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
