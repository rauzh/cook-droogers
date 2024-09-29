// Code generated by go-swagger; DO NOT EDIT.

package artist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetArtistByIDParams creates a new GetArtistByIDParams object
//
// There are no default values defined in the spec.
func NewGetArtistByIDParams() GetArtistByIDParams {

	return GetArtistByIDParams{}
}

// GetArtistByIDParams contains all the bound params for the get artist by ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters getArtistByID
type GetArtistByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID артиста
	  Required: true
	  In: path
	*/
	ArtistID uint64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetArtistByIDParams() beforehand.
func (o *GetArtistByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rArtistID, rhkArtistID, _ := route.Params.GetOK("artist_id")
	if err := o.bindArtistID(rArtistID, rhkArtistID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindArtistID binds and validates parameter ArtistID from path.
func (o *GetArtistByIDParams) bindArtistID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("artist_id", "path", "uint64", raw)
	}
	o.ArtistID = value

	return nil
}
