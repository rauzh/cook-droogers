// Code generated by go-swagger; DO NOT EDIT.

package guest

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

// NewRegisterParams creates a new RegisterParams object
//
// There are no default values defined in the spec.
func NewRegisterParams() RegisterParams {

	return RegisterParams{}
}

// RegisterParams contains all the bound params for the register operation
// typically these are obtained from a http.Request
//
// swagger:parameters register
type RegisterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Email пользователя
	  Required: true
	  In: query
	*/
	Email strfmt.Email
	/*Пароль пользователя
	  Required: true
	  In: query
	*/
	Password string
	/*Имя пользователя
	  Required: true
	  In: query
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRegisterParams() beforehand.
func (o *RegisterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEmail, qhkEmail, _ := qs.GetOK("email")
	if err := o.bindEmail(qEmail, qhkEmail, route.Formats); err != nil {
		res = append(res, err)
	}

	qPassword, qhkPassword, _ := qs.GetOK("password")
	if err := o.bindPassword(qPassword, qhkPassword, route.Formats); err != nil {
		res = append(res, err)
	}

	qUsername, qhkUsername, _ := qs.GetOK("username")
	if err := o.bindUsername(qUsername, qhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmail binds and validates parameter Email from query.
func (o *RegisterParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("email", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("email", "query", raw); err != nil {
		return err
	}

	// Format: email
	value, err := formats.Parse("email", raw)
	if err != nil {
		return errors.InvalidType("email", "query", "strfmt.Email", raw)
	}
	o.Email = *(value.(*strfmt.Email))

	if err := o.validateEmail(formats); err != nil {
		return err
	}

	return nil
}

// validateEmail carries on validations for parameter Email
func (o *RegisterParams) validateEmail(formats strfmt.Registry) error {

	if err := validate.FormatOf("email", "query", "email", o.Email.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindPassword binds and validates parameter Password from query.
func (o *RegisterParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("password", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("password", "query", raw); err != nil {
		return err
	}
	o.Password = raw

	return nil
}

// bindUsername binds and validates parameter Username from query.
func (o *RegisterParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("username", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("username", "query", raw); err != nil {
		return err
	}
	o.Username = raw

	return nil
}
