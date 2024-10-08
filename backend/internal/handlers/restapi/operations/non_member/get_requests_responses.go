// Code generated by go-swagger; DO NOT EDIT.

package non_member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"cookdroogers/internal/handlers/models"
)

// GetRequestsOKCode is the HTTP code returned for type GetRequestsOK
const GetRequestsOKCode int = 200

/*
GetRequestsOK Success

swagger:response getRequestsOK
*/
type GetRequestsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.RequestDTO `json:"body,omitempty"`
}

// NewGetRequestsOK creates GetRequestsOK with default headers values
func NewGetRequestsOK() *GetRequestsOK {

	return &GetRequestsOK{}
}

// WithPayload adds the payload to the get requests o k response
func (o *GetRequestsOK) WithPayload(payload []*models.RequestDTO) *GetRequestsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get requests o k response
func (o *GetRequestsOK) SetPayload(payload []*models.RequestDTO) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRequestsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.RequestDTO, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRequestsInternalServerErrorCode is the HTTP code returned for type GetRequestsInternalServerError
const GetRequestsInternalServerErrorCode int = 500

/*
GetRequestsInternalServerError Internal error

swagger:response getRequestsInternalServerError
*/
type GetRequestsInternalServerError struct {
}

// NewGetRequestsInternalServerError creates GetRequestsInternalServerError with default headers values
func NewGetRequestsInternalServerError() *GetRequestsInternalServerError {

	return &GetRequestsInternalServerError{}
}

// WriteResponse to the client
func (o *GetRequestsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
