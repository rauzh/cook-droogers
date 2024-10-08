// Code generated by go-swagger; DO NOT EDIT.

package artist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddReleaseHandlerFunc turns a function with the right signature into a add release handler
type AddReleaseHandlerFunc func(AddReleaseParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddReleaseHandlerFunc) Handle(params AddReleaseParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddReleaseHandler interface for that can handle valid add release params
type AddReleaseHandler interface {
	Handle(AddReleaseParams, interface{}) middleware.Responder
}

// NewAddRelease creates a new http.Handler for the add release operation
func NewAddRelease(ctx *middleware.Context, handler AddReleaseHandler) *AddRelease {
	return &AddRelease{Context: ctx, Handler: handler}
}

/*
	AddRelease swagger:route POST /releases artist addRelease

Upload release
*/
type AddRelease struct {
	Context *middleware.Context
	Handler AddReleaseHandler
}

func (o *AddRelease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddReleaseParams()
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
