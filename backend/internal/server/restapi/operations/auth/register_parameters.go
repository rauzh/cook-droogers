// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"cookdroogers/internal/server/models"
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

	/*Данные для регистрации нового пользователя
	  Required: true
	  In: body
	*/
	UserData *models.RegisterUserDTO
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRegisterParams() beforehand.
func (o *RegisterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.RegisterUserDTO
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("userData", "body", ""))
			} else {
				res = append(res, errors.NewParseError("userData", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.UserData = &body
			}
		}
	} else {
		res = append(res, errors.Required("userData", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
