// Code generated by go-swagger; DO NOT EDIT.

package tracks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cookdroogers/internal/server/models"
)

// GetTrackByIDOKCode is the HTTP code returned for type GetTrackByIDOK
const GetTrackByIDOKCode int = 200

/*
GetTrackByIDOK Success

swagger:response getTrackByIdOK
*/
type GetTrackByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.TrackDTO `json:"body,omitempty"`
}

// NewGetTrackByIDOK creates GetTrackByIDOK with default headers values
func NewGetTrackByIDOK() *GetTrackByIDOK {

	return &GetTrackByIDOK{}
}

// WithPayload adds the payload to the get track by Id o k response
func (o *GetTrackByIDOK) WithPayload(payload *models.TrackDTO) *GetTrackByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get track by Id o k response
func (o *GetTrackByIDOK) SetPayload(payload *models.TrackDTO) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTrackByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTrackByIDUnauthorizedCode is the HTTP code returned for type GetTrackByIDUnauthorized
const GetTrackByIDUnauthorizedCode int = 401

/*
GetTrackByIDUnauthorized Auth error

swagger:response getTrackByIdUnauthorized
*/
type GetTrackByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetTrackByIDUnauthorized creates GetTrackByIDUnauthorized with default headers values
func NewGetTrackByIDUnauthorized() *GetTrackByIDUnauthorized {

	return &GetTrackByIDUnauthorized{}
}

// WithPayload adds the payload to the get track by Id unauthorized response
func (o *GetTrackByIDUnauthorized) WithPayload(payload *models.LeErrorMessage) *GetTrackByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get track by Id unauthorized response
func (o *GetTrackByIDUnauthorized) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTrackByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTrackByIDForbiddenCode is the HTTP code returned for type GetTrackByIDForbidden
const GetTrackByIDForbiddenCode int = 403

/*
GetTrackByIDForbidden Invalid user type

swagger:response getTrackByIdForbidden
*/
type GetTrackByIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetTrackByIDForbidden creates GetTrackByIDForbidden with default headers values
func NewGetTrackByIDForbidden() *GetTrackByIDForbidden {

	return &GetTrackByIDForbidden{}
}

// WithPayload adds the payload to the get track by Id forbidden response
func (o *GetTrackByIDForbidden) WithPayload(payload *models.LeErrorMessage) *GetTrackByIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get track by Id forbidden response
func (o *GetTrackByIDForbidden) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTrackByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTrackByIDNotFoundCode is the HTTP code returned for type GetTrackByIDNotFound
const GetTrackByIDNotFoundCode int = 404

/*
GetTrackByIDNotFound No such track

swagger:response getTrackByIdNotFound
*/
type GetTrackByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetTrackByIDNotFound creates GetTrackByIDNotFound with default headers values
func NewGetTrackByIDNotFound() *GetTrackByIDNotFound {

	return &GetTrackByIDNotFound{}
}

// WithPayload adds the payload to the get track by Id not found response
func (o *GetTrackByIDNotFound) WithPayload(payload *models.LeErrorMessage) *GetTrackByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get track by Id not found response
func (o *GetTrackByIDNotFound) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTrackByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTrackByIDUnprocessableEntityCode is the HTTP code returned for type GetTrackByIDUnprocessableEntity
const GetTrackByIDUnprocessableEntityCode int = 422

/*
GetTrackByIDUnprocessableEntity Invalid params

swagger:response getTrackByIdUnprocessableEntity
*/
type GetTrackByIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetTrackByIDUnprocessableEntity creates GetTrackByIDUnprocessableEntity with default headers values
func NewGetTrackByIDUnprocessableEntity() *GetTrackByIDUnprocessableEntity {

	return &GetTrackByIDUnprocessableEntity{}
}

// WithPayload adds the payload to the get track by Id unprocessable entity response
func (o *GetTrackByIDUnprocessableEntity) WithPayload(payload *models.LeErrorMessage) *GetTrackByIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get track by Id unprocessable entity response
func (o *GetTrackByIDUnprocessableEntity) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTrackByIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTrackByIDInternalServerErrorCode is the HTTP code returned for type GetTrackByIDInternalServerError
const GetTrackByIDInternalServerErrorCode int = 500

/*
GetTrackByIDInternalServerError Internal error

swagger:response getTrackByIdInternalServerError
*/
type GetTrackByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetTrackByIDInternalServerError creates GetTrackByIDInternalServerError with default headers values
func NewGetTrackByIDInternalServerError() *GetTrackByIDInternalServerError {

	return &GetTrackByIDInternalServerError{}
}

// WithPayload adds the payload to the get track by Id internal server error response
func (o *GetTrackByIDInternalServerError) WithPayload(payload *models.LeErrorMessage) *GetTrackByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get track by Id internal server error response
func (o *GetTrackByIDInternalServerError) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTrackByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
