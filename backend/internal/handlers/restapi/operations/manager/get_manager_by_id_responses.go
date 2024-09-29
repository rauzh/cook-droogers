// Code generated by go-swagger; DO NOT EDIT.

package manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cookdroogers/internal/handlers/models"
)

// GetManagerByIDOKCode is the HTTP code returned for type GetManagerByIDOK
const GetManagerByIDOKCode int = 200

/*
GetManagerByIDOK Success

swagger:response getManagerByIdOK
*/
type GetManagerByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ManagerDTO `json:"body,omitempty"`
}

// NewGetManagerByIDOK creates GetManagerByIDOK with default headers values
func NewGetManagerByIDOK() *GetManagerByIDOK {

	return &GetManagerByIDOK{}
}

// WithPayload adds the payload to the get manager by Id o k response
func (o *GetManagerByIDOK) WithPayload(payload *models.ManagerDTO) *GetManagerByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get manager by Id o k response
func (o *GetManagerByIDOK) SetPayload(payload *models.ManagerDTO) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagerByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetManagerByIDUnauthorizedCode is the HTTP code returned for type GetManagerByIDUnauthorized
const GetManagerByIDUnauthorizedCode int = 401

/*
GetManagerByIDUnauthorized Auth error

swagger:response getManagerByIdUnauthorized
*/
type GetManagerByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagerByIDUnauthorized creates GetManagerByIDUnauthorized with default headers values
func NewGetManagerByIDUnauthorized() *GetManagerByIDUnauthorized {

	return &GetManagerByIDUnauthorized{}
}

// WithPayload adds the payload to the get manager by Id unauthorized response
func (o *GetManagerByIDUnauthorized) WithPayload(payload *models.LeErrorMessage) *GetManagerByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get manager by Id unauthorized response
func (o *GetManagerByIDUnauthorized) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagerByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetManagerByIDForbiddenCode is the HTTP code returned for type GetManagerByIDForbidden
const GetManagerByIDForbiddenCode int = 403

/*
GetManagerByIDForbidden Invalid user type

swagger:response getManagerByIdForbidden
*/
type GetManagerByIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagerByIDForbidden creates GetManagerByIDForbidden with default headers values
func NewGetManagerByIDForbidden() *GetManagerByIDForbidden {

	return &GetManagerByIDForbidden{}
}

// WithPayload adds the payload to the get manager by Id forbidden response
func (o *GetManagerByIDForbidden) WithPayload(payload *models.LeErrorMessage) *GetManagerByIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get manager by Id forbidden response
func (o *GetManagerByIDForbidden) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagerByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetManagerByIDNotFoundCode is the HTTP code returned for type GetManagerByIDNotFound
const GetManagerByIDNotFoundCode int = 404

/*
GetManagerByIDNotFound No such manager

swagger:response getManagerByIdNotFound
*/
type GetManagerByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagerByIDNotFound creates GetManagerByIDNotFound with default headers values
func NewGetManagerByIDNotFound() *GetManagerByIDNotFound {

	return &GetManagerByIDNotFound{}
}

// WithPayload adds the payload to the get manager by Id not found response
func (o *GetManagerByIDNotFound) WithPayload(payload *models.LeErrorMessage) *GetManagerByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get manager by Id not found response
func (o *GetManagerByIDNotFound) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagerByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetManagerByIDUnprocessableEntityCode is the HTTP code returned for type GetManagerByIDUnprocessableEntity
const GetManagerByIDUnprocessableEntityCode int = 422

/*
GetManagerByIDUnprocessableEntity Invalid params

swagger:response getManagerByIdUnprocessableEntity
*/
type GetManagerByIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagerByIDUnprocessableEntity creates GetManagerByIDUnprocessableEntity with default headers values
func NewGetManagerByIDUnprocessableEntity() *GetManagerByIDUnprocessableEntity {

	return &GetManagerByIDUnprocessableEntity{}
}

// WithPayload adds the payload to the get manager by Id unprocessable entity response
func (o *GetManagerByIDUnprocessableEntity) WithPayload(payload *models.LeErrorMessage) *GetManagerByIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get manager by Id unprocessable entity response
func (o *GetManagerByIDUnprocessableEntity) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagerByIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetManagerByIDInternalServerErrorCode is the HTTP code returned for type GetManagerByIDInternalServerError
const GetManagerByIDInternalServerErrorCode int = 500

/*
GetManagerByIDInternalServerError Internal error

swagger:response getManagerByIdInternalServerError
*/
type GetManagerByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagerByIDInternalServerError creates GetManagerByIDInternalServerError with default headers values
func NewGetManagerByIDInternalServerError() *GetManagerByIDInternalServerError {

	return &GetManagerByIDInternalServerError{}
}

// WithPayload adds the payload to the get manager by Id internal server error response
func (o *GetManagerByIDInternalServerError) WithPayload(payload *models.LeErrorMessage) *GetManagerByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get manager by Id internal server error response
func (o *GetManagerByIDInternalServerError) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagerByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}