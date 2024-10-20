package requests

import (
	"cookdroogers/app"
	"cookdroogers/internal/server/restapi/operations"
	"cookdroogers/internal/server/restapi/operations/requests"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureRequestsHandlers(app *app.App, api *operations.SwaggerCookDroogersAPI) {
	api.RequestsGetRequestsHandler = requests.GetRequestsHandlerFunc(func(params requests.GetRequestsParams, principal interface{}) middleware.Responder {
		return getRequestsHandlerFunc(params, app)
	})
	api.RequestsGetRequestHandler = requests.GetRequestHandlerFunc(func(params requests.GetRequestParams, principal interface{}) middleware.Responder {
		return getRequestByIDHandlerFunc(params, app)
	})
	api.RequestsSignContractHandler = requests.SignContractHandlerFunc(func(params requests.SignContractParams, principal interface{}) middleware.Responder {
		return signContractHandlerFunc(params, app)
	})
	api.RequestsPublishReqHandler = requests.PublishReqHandlerFunc(func(params requests.PublishReqParams, principal interface{}) middleware.Responder {
		return publishHandlerFunc(params, app)
	})
	api.RequestsAcceptRequestHandler = requests.AcceptRequestHandlerFunc(func(params requests.AcceptRequestParams, principal interface{}) middleware.Responder {
		return acceptHandlerFunc(params, app)
	})
	api.RequestsDeclineRequestHandler = requests.DeclineRequestHandlerFunc(func(params requests.DeclineRequestParams, principal interface{}) middleware.Responder {
		return declineHandlerFunc(params, app)
	})
}
