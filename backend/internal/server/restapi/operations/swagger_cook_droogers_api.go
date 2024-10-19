// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"cookdroogers/internal/server/restapi/operations/artist"
	"cookdroogers/internal/server/restapi/operations/auth"
	"cookdroogers/internal/server/restapi/operations/manager"
	"cookdroogers/internal/server/restapi/operations/releases"
	"cookdroogers/internal/server/restapi/operations/requests"
	"cookdroogers/internal/server/restapi/operations/tracks"
	"cookdroogers/internal/server/restapi/operations/user"
	"cookdroogers/internal/server/restapi/operations/users"
)

// NewSwaggerCookDroogersAPI creates a new SwaggerCookDroogers instance
func NewSwaggerCookDroogersAPI(spec *loads.Document) *SwaggerCookDroogersAPI {
	return &SwaggerCookDroogersAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		GetHeartbeatHandler: GetHeartbeatHandlerFunc(func(params GetHeartbeatParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHeartbeat has not yet been implemented")
		}),
		RequestsAcceptRequestHandler: requests.AcceptRequestHandlerFunc(func(params requests.AcceptRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.AcceptRequest has not yet been implemented")
		}),
		ManagerAddManagerHandler: manager.AddManagerHandlerFunc(func(params manager.AddManagerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.AddManager has not yet been implemented")
		}),
		ReleasesAddReleaseHandler: releases.AddReleaseHandlerFunc(func(params releases.AddReleaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation releases.AddRelease has not yet been implemented")
		}),
		RequestsDeclineRequestHandler: requests.DeclineRequestHandlerFunc(func(params requests.DeclineRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.DeclineRequest has not yet been implemented")
		}),
		ArtistGetArtistByIDHandler: artist.GetArtistByIDHandlerFunc(func(params artist.GetArtistByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation artist.GetArtistByID has not yet been implemented")
		}),
		ManagerGetManagerByIDHandler: manager.GetManagerByIDHandlerFunc(func(params manager.GetManagerByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.GetManagerByID has not yet been implemented")
		}),
		ManagerGetManagersHandler: manager.GetManagersHandlerFunc(func(params manager.GetManagersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manager.GetManagers has not yet been implemented")
		}),
		ReleasesGetReleaseHandler: releases.GetReleaseHandlerFunc(func(params releases.GetReleaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation releases.GetRelease has not yet been implemented")
		}),
		ReleasesGetReleaseByIDHandler: releases.GetReleaseByIDHandlerFunc(func(params releases.GetReleaseByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation releases.GetReleaseByID has not yet been implemented")
		}),
		RequestsGetRequestHandler: requests.GetRequestHandlerFunc(func(params requests.GetRequestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.GetRequest has not yet been implemented")
		}),
		RequestsGetRequestsHandler: requests.GetRequestsHandlerFunc(func(params requests.GetRequestsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.GetRequests has not yet been implemented")
		}),
		TracksGetTrackByIDHandler: tracks.GetTrackByIDHandlerFunc(func(params tracks.GetTrackByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tracks.GetTrackByID has not yet been implemented")
		}),
		UsersGetUserByIDHandler: users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUserByID has not yet been implemented")
		}),
		UserGetUsersHandler: user.GetUsersHandlerFunc(func(params user.GetUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUsers has not yet been implemented")
		}),
		AuthLoginHandler: auth.LoginHandlerFunc(func(params auth.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.Login has not yet been implemented")
		}),
		RequestsPublishReqHandler: requests.PublishReqHandlerFunc(func(params requests.PublishReqParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.PublishReq has not yet been implemented")
		}),
		AuthRegisterHandler: auth.RegisterHandlerFunc(func(params auth.RegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.Register has not yet been implemented")
		}),
		RequestsSignContractHandler: requests.SignContractHandlerFunc(func(params requests.SignContractParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation requests.SignContract has not yet been implemented")
		}),

		// Applies when the "access_token" header is set
		JWTAuthAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (JWTAuth) access_token from header param [access_token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*SwaggerCookDroogersAPI the swagger cook droogers API */
type SwaggerCookDroogersAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// JWTAuthAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key access_token provided in the header
	JWTAuthAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// GetHeartbeatHandler sets the operation handler for the get heartbeat operation
	GetHeartbeatHandler GetHeartbeatHandler
	// RequestsAcceptRequestHandler sets the operation handler for the accept request operation
	RequestsAcceptRequestHandler requests.AcceptRequestHandler
	// ManagerAddManagerHandler sets the operation handler for the add manager operation
	ManagerAddManagerHandler manager.AddManagerHandler
	// ReleasesAddReleaseHandler sets the operation handler for the add release operation
	ReleasesAddReleaseHandler releases.AddReleaseHandler
	// RequestsDeclineRequestHandler sets the operation handler for the decline request operation
	RequestsDeclineRequestHandler requests.DeclineRequestHandler
	// ArtistGetArtistByIDHandler sets the operation handler for the get artist by ID operation
	ArtistGetArtistByIDHandler artist.GetArtistByIDHandler
	// ManagerGetManagerByIDHandler sets the operation handler for the get manager by ID operation
	ManagerGetManagerByIDHandler manager.GetManagerByIDHandler
	// ManagerGetManagersHandler sets the operation handler for the get managers operation
	ManagerGetManagersHandler manager.GetManagersHandler
	// ReleasesGetReleaseHandler sets the operation handler for the get release operation
	ReleasesGetReleaseHandler releases.GetReleaseHandler
	// ReleasesGetReleaseByIDHandler sets the operation handler for the get release by ID operation
	ReleasesGetReleaseByIDHandler releases.GetReleaseByIDHandler
	// RequestsGetRequestHandler sets the operation handler for the get request operation
	RequestsGetRequestHandler requests.GetRequestHandler
	// RequestsGetRequestsHandler sets the operation handler for the get requests operation
	RequestsGetRequestsHandler requests.GetRequestsHandler
	// TracksGetTrackByIDHandler sets the operation handler for the get track by ID operation
	TracksGetTrackByIDHandler tracks.GetTrackByIDHandler
	// UsersGetUserByIDHandler sets the operation handler for the get user by ID operation
	UsersGetUserByIDHandler users.GetUserByIDHandler
	// UserGetUsersHandler sets the operation handler for the get users operation
	UserGetUsersHandler user.GetUsersHandler
	// AuthLoginHandler sets the operation handler for the login operation
	AuthLoginHandler auth.LoginHandler
	// RequestsPublishReqHandler sets the operation handler for the publish req operation
	RequestsPublishReqHandler requests.PublishReqHandler
	// AuthRegisterHandler sets the operation handler for the register operation
	AuthRegisterHandler auth.RegisterHandler
	// RequestsSignContractHandler sets the operation handler for the sign contract operation
	RequestsSignContractHandler requests.SignContractHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *SwaggerCookDroogersAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *SwaggerCookDroogersAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *SwaggerCookDroogersAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SwaggerCookDroogersAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SwaggerCookDroogersAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SwaggerCookDroogersAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SwaggerCookDroogersAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SwaggerCookDroogersAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SwaggerCookDroogersAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SwaggerCookDroogersAPI
func (o *SwaggerCookDroogersAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.JWTAuthAuth == nil {
		unregistered = append(unregistered, "AccessTokenAuth")
	}

	if o.GetHeartbeatHandler == nil {
		unregistered = append(unregistered, "GetHeartbeatHandler")
	}
	if o.RequestsAcceptRequestHandler == nil {
		unregistered = append(unregistered, "requests.AcceptRequestHandler")
	}
	if o.ManagerAddManagerHandler == nil {
		unregistered = append(unregistered, "manager.AddManagerHandler")
	}
	if o.ReleasesAddReleaseHandler == nil {
		unregistered = append(unregistered, "releases.AddReleaseHandler")
	}
	if o.RequestsDeclineRequestHandler == nil {
		unregistered = append(unregistered, "requests.DeclineRequestHandler")
	}
	if o.ArtistGetArtistByIDHandler == nil {
		unregistered = append(unregistered, "artist.GetArtistByIDHandler")
	}
	if o.ManagerGetManagerByIDHandler == nil {
		unregistered = append(unregistered, "manager.GetManagerByIDHandler")
	}
	if o.ManagerGetManagersHandler == nil {
		unregistered = append(unregistered, "manager.GetManagersHandler")
	}
	if o.ReleasesGetReleaseHandler == nil {
		unregistered = append(unregistered, "releases.GetReleaseHandler")
	}
	if o.ReleasesGetReleaseByIDHandler == nil {
		unregistered = append(unregistered, "releases.GetReleaseByIDHandler")
	}
	if o.RequestsGetRequestHandler == nil {
		unregistered = append(unregistered, "requests.GetRequestHandler")
	}
	if o.RequestsGetRequestsHandler == nil {
		unregistered = append(unregistered, "requests.GetRequestsHandler")
	}
	if o.TracksGetTrackByIDHandler == nil {
		unregistered = append(unregistered, "tracks.GetTrackByIDHandler")
	}
	if o.UsersGetUserByIDHandler == nil {
		unregistered = append(unregistered, "users.GetUserByIDHandler")
	}
	if o.UserGetUsersHandler == nil {
		unregistered = append(unregistered, "user.GetUsersHandler")
	}
	if o.AuthLoginHandler == nil {
		unregistered = append(unregistered, "auth.LoginHandler")
	}
	if o.RequestsPublishReqHandler == nil {
		unregistered = append(unregistered, "requests.PublishReqHandler")
	}
	if o.AuthRegisterHandler == nil {
		unregistered = append(unregistered, "auth.RegisterHandler")
	}
	if o.RequestsSignContractHandler == nil {
		unregistered = append(unregistered, "requests.SignContractHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SwaggerCookDroogersAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SwaggerCookDroogersAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "JWTAuth":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.JWTAuthAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *SwaggerCookDroogersAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerCookDroogersAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerCookDroogersAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SwaggerCookDroogersAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the swagger cook droogers API
func (o *SwaggerCookDroogersAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SwaggerCookDroogersAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/heartbeat"] = NewGetHeartbeat(o.context, o.GetHeartbeatHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/requests/{req_id}/accept"] = requests.NewAcceptRequest(o.context, o.RequestsAcceptRequestHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/managers"] = manager.NewAddManager(o.context, o.ManagerAddManagerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/releases"] = releases.NewAddRelease(o.context, o.ReleasesAddReleaseHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/requests/{req_id}/decline"] = requests.NewDeclineRequest(o.context, o.RequestsDeclineRequestHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/artists/{artist_id}"] = artist.NewGetArtistByID(o.context, o.ArtistGetArtistByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/managers/{manager_id}"] = manager.NewGetManagerByID(o.context, o.ManagerGetManagerByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/managers"] = manager.NewGetManagers(o.context, o.ManagerGetManagersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/releases"] = releases.NewGetRelease(o.context, o.ReleasesGetReleaseHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/releases/{release_id}"] = releases.NewGetReleaseByID(o.context, o.ReleasesGetReleaseByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/requests/{req_id}"] = requests.NewGetRequest(o.context, o.RequestsGetRequestHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/requests"] = requests.NewGetRequests(o.context, o.RequestsGetRequestsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tracks/{track_id}"] = tracks.NewGetTrackByID(o.context, o.TracksGetTrackByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{user_id}"] = users.NewGetUserByID(o.context, o.UsersGetUserByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = user.NewGetUsers(o.context, o.UserGetUsersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = auth.NewLogin(o.context, o.AuthLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/requests/publications"] = requests.NewPublishReq(o.context, o.RequestsPublishReqHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/register"] = auth.NewRegister(o.context, o.AuthRegisterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/requests/contracts"] = requests.NewSignContract(o.context, o.RequestsSignContractHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SwaggerCookDroogersAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SwaggerCookDroogersAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SwaggerCookDroogersAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SwaggerCookDroogersAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SwaggerCookDroogersAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
