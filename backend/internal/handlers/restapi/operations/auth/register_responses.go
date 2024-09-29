// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cookdroogers/internal/handlers/models"
)

// RegisterCreatedCode is the HTTP code returned for type RegisterCreated
const RegisterCreatedCode int = 201

/*
RegisterCreated User successfully created

swagger:response registerCreated
*/
type RegisterCreated struct {
}

// NewRegisterCreated creates RegisterCreated with default headers values
func NewRegisterCreated() *RegisterCreated {

	return &RegisterCreated{}
}

// WriteResponse to the client
func (o *RegisterCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// RegisterConflictCode is the HTTP code returned for type RegisterConflict
const RegisterConflictCode int = 409

/*
RegisterConflict User already exists

swagger:response registerConflict
*/
type RegisterConflict struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewRegisterConflict creates RegisterConflict with default headers values
func NewRegisterConflict() *RegisterConflict {

	return &RegisterConflict{}
}

// WithPayload adds the payload to the register conflict response
func (o *RegisterConflict) WithPayload(payload *models.LeErrorMessage) *RegisterConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register conflict response
func (o *RegisterConflict) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUnprocessableEntityCode is the HTTP code returned for type RegisterUnprocessableEntity
const RegisterUnprocessableEntityCode int = 422

/*
RegisterUnprocessableEntity Invalid params

swagger:response registerUnprocessableEntity
*/
type RegisterUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewRegisterUnprocessableEntity creates RegisterUnprocessableEntity with default headers values
func NewRegisterUnprocessableEntity() *RegisterUnprocessableEntity {

	return &RegisterUnprocessableEntity{}
}

// WithPayload adds the payload to the register unprocessable entity response
func (o *RegisterUnprocessableEntity) WithPayload(payload *models.LeErrorMessage) *RegisterUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register unprocessable entity response
func (o *RegisterUnprocessableEntity) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterInternalServerErrorCode is the HTTP code returned for type RegisterInternalServerError
const RegisterInternalServerErrorCode int = 500

/*
RegisterInternalServerError Internal error

swagger:response registerInternalServerError
*/
type RegisterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewRegisterInternalServerError creates RegisterInternalServerError with default headers values
func NewRegisterInternalServerError() *RegisterInternalServerError {

	return &RegisterInternalServerError{}
}

// WithPayload adds the payload to the register internal server error response
func (o *RegisterInternalServerError) WithPayload(payload *models.LeErrorMessage) *RegisterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register internal server error response
func (o *RegisterInternalServerError) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}