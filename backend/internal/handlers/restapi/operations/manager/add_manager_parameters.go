// Code generated by go-swagger; DO NOT EDIT.

package manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAddManagerParams creates a new AddManagerParams object
//
// There are no default values defined in the spec.
func NewAddManagerParams() AddManagerParams {

	return AddManagerParams{}
}

// AddManagerParams contains all the bound params for the add manager operation
// typically these are obtained from a http.Request
//
// swagger:parameters addManager
type AddManagerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID пользователя
	  Required: true
	  In: query
	*/
	UserID []uint64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddManagerParams() beforehand.
func (o *AddManagerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qUserID, qhkUserID, _ := qs.GetOK("user_id")
	if err := o.bindUserID(qUserID, qhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserID binds and validates array parameter UserID from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *AddManagerParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("user_id", "query", rawData)
	}
	var qvUserID string
	if len(rawData) > 0 {
		qvUserID = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	userIDIC := swag.SplitByFormat(qvUserID, "")
	if len(userIDIC) == 0 {
		return errors.Required("user_id", "query", userIDIC)
	}

	var userIDIR []uint64
	for i, userIDIV := range userIDIC {
		// items.Format: "uint64"
		userIDI, err := swag.ConvertUint64(userIDIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "user_id", i), "query", "uint64", userIDI)
		}

		userIDIR = append(userIDIR, userIDI)
	}

	o.UserID = userIDIR

	return nil
}
