// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cookdroogers/internal/server/models"
)

// DeclineRequestOKCode is the HTTP code returned for type DeclineRequestOK
const DeclineRequestOKCode int = 200

/*
DeclineRequestOK Success

swagger:response declineRequestOK
*/
type DeclineRequestOK struct {
}

// NewDeclineRequestOK creates DeclineRequestOK with default headers values
func NewDeclineRequestOK() *DeclineRequestOK {

	return &DeclineRequestOK{}
}

// WriteResponse to the client
func (o *DeclineRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeclineRequestUnauthorizedCode is the HTTP code returned for type DeclineRequestUnauthorized
const DeclineRequestUnauthorizedCode int = 401

/*
DeclineRequestUnauthorized Auth error

swagger:response declineRequestUnauthorized
*/
type DeclineRequestUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewDeclineRequestUnauthorized creates DeclineRequestUnauthorized with default headers values
func NewDeclineRequestUnauthorized() *DeclineRequestUnauthorized {

	return &DeclineRequestUnauthorized{}
}

// WithPayload adds the payload to the decline request unauthorized response
func (o *DeclineRequestUnauthorized) WithPayload(payload *models.LeErrorMessage) *DeclineRequestUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the decline request unauthorized response
func (o *DeclineRequestUnauthorized) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeclineRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeclineRequestForbiddenCode is the HTTP code returned for type DeclineRequestForbidden
const DeclineRequestForbiddenCode int = 403

/*
DeclineRequestForbidden Invalid user type (or this manager is not maintainer of the request)

swagger:response declineRequestForbidden
*/
type DeclineRequestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewDeclineRequestForbidden creates DeclineRequestForbidden with default headers values
func NewDeclineRequestForbidden() *DeclineRequestForbidden {

	return &DeclineRequestForbidden{}
}

// WithPayload adds the payload to the decline request forbidden response
func (o *DeclineRequestForbidden) WithPayload(payload *models.LeErrorMessage) *DeclineRequestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the decline request forbidden response
func (o *DeclineRequestForbidden) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeclineRequestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeclineRequestNotFoundCode is the HTTP code returned for type DeclineRequestNotFound
const DeclineRequestNotFoundCode int = 404

/*
DeclineRequestNotFound No such request

swagger:response declineRequestNotFound
*/
type DeclineRequestNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewDeclineRequestNotFound creates DeclineRequestNotFound with default headers values
func NewDeclineRequestNotFound() *DeclineRequestNotFound {

	return &DeclineRequestNotFound{}
}

// WithPayload adds the payload to the decline request not found response
func (o *DeclineRequestNotFound) WithPayload(payload *models.LeErrorMessage) *DeclineRequestNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the decline request not found response
func (o *DeclineRequestNotFound) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeclineRequestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeclineRequestUnprocessableEntityCode is the HTTP code returned for type DeclineRequestUnprocessableEntity
const DeclineRequestUnprocessableEntityCode int = 422

/*
DeclineRequestUnprocessableEntity Invalid params

swagger:response declineRequestUnprocessableEntity
*/
type DeclineRequestUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewDeclineRequestUnprocessableEntity creates DeclineRequestUnprocessableEntity with default headers values
func NewDeclineRequestUnprocessableEntity() *DeclineRequestUnprocessableEntity {

	return &DeclineRequestUnprocessableEntity{}
}

// WithPayload adds the payload to the decline request unprocessable entity response
func (o *DeclineRequestUnprocessableEntity) WithPayload(payload *models.LeErrorMessage) *DeclineRequestUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the decline request unprocessable entity response
func (o *DeclineRequestUnprocessableEntity) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeclineRequestUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeclineRequestInternalServerErrorCode is the HTTP code returned for type DeclineRequestInternalServerError
const DeclineRequestInternalServerErrorCode int = 500

/*
DeclineRequestInternalServerError Internal error

swagger:response declineRequestInternalServerError
*/
type DeclineRequestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewDeclineRequestInternalServerError creates DeclineRequestInternalServerError with default headers values
func NewDeclineRequestInternalServerError() *DeclineRequestInternalServerError {

	return &DeclineRequestInternalServerError{}
}

// WithPayload adds the payload to the decline request internal server error response
func (o *DeclineRequestInternalServerError) WithPayload(payload *models.LeErrorMessage) *DeclineRequestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the decline request internal server error response
func (o *DeclineRequestInternalServerError) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeclineRequestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}