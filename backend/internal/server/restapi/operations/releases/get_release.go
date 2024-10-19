// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetReleaseHandlerFunc turns a function with the right signature into a get release handler
type GetReleaseHandlerFunc func(GetReleaseParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReleaseHandlerFunc) Handle(params GetReleaseParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetReleaseHandler interface for that can handle valid get release params
type GetReleaseHandler interface {
	Handle(GetReleaseParams, interface{}) middleware.Responder
}

// NewGetRelease creates a new http.Handler for the get release operation
func NewGetRelease(ctx *middleware.Context, handler GetReleaseHandler) *GetRelease {
	return &GetRelease{Context: ctx, Handler: handler}
}

/*
	GetRelease swagger:route GET /releases releases getRelease

Get releases
*/
type GetRelease struct {
	Context *middleware.Context
	Handler GetReleaseHandler
}

func (o *GetRelease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetReleaseParams()
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