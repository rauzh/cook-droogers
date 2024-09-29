// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRequestHandlerFunc turns a function with the right signature into a get request handler
type GetRequestHandlerFunc func(GetRequestParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRequestHandlerFunc) Handle(params GetRequestParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRequestHandler interface for that can handle valid get request params
type GetRequestHandler interface {
	Handle(GetRequestParams, interface{}) middleware.Responder
}

// NewGetRequest creates a new http.Handler for the get request operation
func NewGetRequest(ctx *middleware.Context, handler GetRequestHandler) *GetRequest {
	return &GetRequest{Context: ctx, Handler: handler}
}

/*
	GetRequest swagger:route GET /requests/{req_id} requests getRequest

Get specified request
*/
type GetRequest struct {
	Context *middleware.Context
	Handler GetRequestHandler
}

func (o *GetRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRequestParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}