// Code generated by go-swagger; DO NOT EDIT.

package requests

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

	/*Информация для запроса на подписание контракта
	  Required: true
	  In: body
	*/
	SignRequest *models.CreateSignRequestDTO
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSignContractParams() beforehand.
func (o *SignContractParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CreateSignRequestDTO
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("signRequest", "body", ""))
			} else {
				res = append(res, errors.NewParseError("signRequest", "body", "", err))
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
				o.SignRequest = &body
			}
		}
	} else {
		res = append(res, errors.Required("signRequest", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
