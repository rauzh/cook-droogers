// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AcceptRequestHandlerFunc turns a function with the right signature into a accept request handler
type AcceptRequestHandlerFunc func(AcceptRequestParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AcceptRequestHandlerFunc) Handle(params AcceptRequestParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AcceptRequestHandler interface for that can handle valid accept request params
type AcceptRequestHandler interface {
	Handle(AcceptRequestParams, interface{}) middleware.Responder
}

// NewAcceptRequest creates a new http.Handler for the accept request operation
func NewAcceptRequest(ctx *middleware.Context, handler AcceptRequestHandler) *AcceptRequest {
	return &AcceptRequest{Context: ctx, Handler: handler}
}

/*
	AcceptRequest swagger:route PATCH /requests/{req_id}/accept requests acceptRequest

Accept specified request
*/
type AcceptRequest struct {
	Context *middleware.Context
	Handler AcceptRequestHandler
}

func (o *AcceptRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAcceptRequestParams()
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