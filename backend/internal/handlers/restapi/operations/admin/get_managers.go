// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetManagersHandlerFunc turns a function with the right signature into a get managers handler
type GetManagersHandlerFunc func(GetManagersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetManagersHandlerFunc) Handle(params GetManagersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetManagersHandler interface for that can handle valid get managers params
type GetManagersHandler interface {
	Handle(GetManagersParams, interface{}) middleware.Responder
}

// NewGetManagers creates a new http.Handler for the get managers operation
func NewGetManagers(ctx *middleware.Context, handler GetManagersHandler) *GetManagers {
	return &GetManagers{Context: ctx, Handler: handler}
}

/*
	GetManagers swagger:route GET /managers admin getManagers

Get list of managers
*/
type GetManagers struct {
	Context *middleware.Context
	Handler GetManagersHandler
}

func (o *GetManagers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetManagersParams()
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
