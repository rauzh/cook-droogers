// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewSignContractParams creates a new SignContractParams object
//
// There are no default values defined in the spec.
func NewSignContractParams() SignContractParams {

	return SignContractParams{}
}

// SignContractParams contains all the bound params for the sign contract operation
// typically these are obtained from a http.Request
//
// swagger:parameters signContract
type SignContractParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Псевдоним
	  Required: true
	  In: query
	*/
	Nickname string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSignContractParams() beforehand.
func (o *SignContractParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qNickname, qhkNickname, _ := qs.GetOK("nickname")
	if err := o.bindNickname(qNickname, qhkNickname, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNickname binds and validates parameter Nickname from query.
func (o *SignContractParams) bindNickname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("nickname", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("nickname", "query", raw); err != nil {
		return err
	}
	o.Nickname = raw

	return nil
}
