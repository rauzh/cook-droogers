// Code generated by go-swagger; DO NOT EDIT.

package non_member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRequestsHandlerFunc turns a function with the right signature into a get requests handler
type GetRequestsHandlerFunc func(GetRequestsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRequestsHandlerFunc) Handle(params GetRequestsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRequestsHandler interface for that can handle valid get requests params
type GetRequestsHandler interface {
	Handle(GetRequestsParams, interface{}) middleware.Responder
}

// NewGetRequests creates a new http.Handler for the get requests operation
func NewGetRequests(ctx *middleware.Context, handler GetRequestsHandler) *GetRequests {
	return &GetRequests{Context: ctx, Handler: handler}
}

/*
	GetRequests swagger:route GET /requests non-member manager artist getRequests

Get requests
*/
type GetRequests struct {
	Context *middleware.Context
	Handler GetRequestsHandler
}

func (o *GetRequests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRequestsParams()
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
