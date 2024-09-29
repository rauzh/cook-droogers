// Code generated by go-swagger; DO NOT EDIT.

package manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cookdroogers/internal/handlers/models"
)

// GetManagersOKCode is the HTTP code returned for type GetManagersOK
const GetManagersOKCode int = 200

/*
GetManagersOK Success

swagger:response getManagersOK
*/
type GetManagersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ManagerDTO `json:"body,omitempty"`
}

// NewGetManagersOK creates GetManagersOK with default headers values
func NewGetManagersOK() *GetManagersOK {

	return &GetManagersOK{}
}

// WithPayload adds the payload to the get managers o k response
func (o *GetManagersOK) WithPayload(payload []*models.ManagerDTO) *GetManagersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get managers o k response
func (o *GetManagersOK) SetPayload(payload []*models.ManagerDTO) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ManagerDTO, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetManagersUnauthorizedCode is the HTTP code returned for type GetManagersUnauthorized
const GetManagersUnauthorizedCode int = 401

/*
GetManagersUnauthorized Auth error

swagger:response getManagersUnauthorized
*/
type GetManagersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagersUnauthorized creates GetManagersUnauthorized with default headers values
func NewGetManagersUnauthorized() *GetManagersUnauthorized {

	return &GetManagersUnauthorized{}
}

// WithPayload adds the payload to the get managers unauthorized response
func (o *GetManagersUnauthorized) WithPayload(payload *models.LeErrorMessage) *GetManagersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get managers unauthorized response
func (o *GetManagersUnauthorized) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetManagersForbiddenCode is the HTTP code returned for type GetManagersForbidden
const GetManagersForbiddenCode int = 403

/*
GetManagersForbidden Invalid user type

swagger:response getManagersForbidden
*/
type GetManagersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagersForbidden creates GetManagersForbidden with default headers values
func NewGetManagersForbidden() *GetManagersForbidden {

	return &GetManagersForbidden{}
}

// WithPayload adds the payload to the get managers forbidden response
func (o *GetManagersForbidden) WithPayload(payload *models.LeErrorMessage) *GetManagersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get managers forbidden response
func (o *GetManagersForbidden) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetManagersInternalServerErrorCode is the HTTP code returned for type GetManagersInternalServerError
const GetManagersInternalServerErrorCode int = 500

/*
GetManagersInternalServerError Internal error

swagger:response getManagersInternalServerError
*/
type GetManagersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.LeErrorMessage `json:"body,omitempty"`
}

// NewGetManagersInternalServerError creates GetManagersInternalServerError with default headers values
func NewGetManagersInternalServerError() *GetManagersInternalServerError {

	return &GetManagersInternalServerError{}
}

// WithPayload adds the payload to the get managers internal server error response
func (o *GetManagersInternalServerError) WithPayload(payload *models.LeErrorMessage) *GetManagersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get managers internal server error response
func (o *GetManagersInternalServerError) SetPayload(payload *models.LeErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
