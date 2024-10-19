// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cookdroogers/internal/server/models"
)

// LoginOKCode is the HTTP code returned for type LoginOK
const LoginOKCode int = 200

/*
LoginOK User successfully logged in

swagger:response loginOK
*/
type LoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccessTokenDTO `json:"body,omitempty"`
}

// NewLoginOK creates LoginOK with default headers values
func NewLoginOK() *LoginOK {

	return &LoginOK{}
}

// WithPayload adds the payload to the login o k response
func (o *LoginOK) WithPayload(payload *models.AccessTokenDTO) *LoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login o k response
func (o *LoginOK) SetPayload(payload *models.AccessTokenDTO) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUnauthorizedCode is the HTTP code returned for type LoginUnauthorized
const LoginUnauthorizedCode int = 401

/*
LoginUnauthorized Auth error

swagger:response loginUnauthorized
*/
type LoginUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewLoginUnauthorized creates LoginUnauthorized with default headers values
func NewLoginUnauthorized() *LoginUnauthorized {

	return &LoginUnauthorized{}
}

// WithPayload adds the payload to the login unauthorized response
func (o *LoginUnauthorized) WithPayload(payload *models.LeErrorMessage) *LoginUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login unauthorized response
func (o *LoginUnauthorized) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginNotFoundCode is the HTTP code returned for type LoginNotFound
const LoginNotFoundCode int = 404

/*
LoginNotFound No such user

swagger:response loginNotFound
*/
type LoginNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewLoginNotFound creates LoginNotFound with default headers values
func NewLoginNotFound() *LoginNotFound {

	return &LoginNotFound{}
}

// WithPayload adds the payload to the login not found response
func (o *LoginNotFound) WithPayload(payload *models.LeErrorMessage) *LoginNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login not found response
func (o *LoginNotFound) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUnprocessableEntityCode is the HTTP code returned for type LoginUnprocessableEntity
const LoginUnprocessableEntityCode int = 422

/*
LoginUnprocessableEntity Invalid params

swagger:response loginUnprocessableEntity
*/
type LoginUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewLoginUnprocessableEntity creates LoginUnprocessableEntity with default headers values
func NewLoginUnprocessableEntity() *LoginUnprocessableEntity {

	return &LoginUnprocessableEntity{}
}

// WithPayload adds the payload to the login unprocessable entity response
func (o *LoginUnprocessableEntity) WithPayload(payload *models.LeErrorMessage) *LoginUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login unprocessable entity response
func (o *LoginUnprocessableEntity) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginInternalServerErrorCode is the HTTP code returned for type LoginInternalServerError
const LoginInternalServerErrorCode int = 500

/*
LoginInternalServerError Internal error

swagger:response loginInternalServerError
*/
type LoginInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewLoginInternalServerError creates LoginInternalServerError with default headers values
func NewLoginInternalServerError() *LoginInternalServerError {

	return &LoginInternalServerError{}
}

// WithPayload adds the payload to the login internal server error response
func (o *LoginInternalServerError) WithPayload(payload *models.LeErrorMessage) *LoginInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login internal server error response
func (o *LoginInternalServerError) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}