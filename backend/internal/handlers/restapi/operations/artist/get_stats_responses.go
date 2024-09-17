// Code generated by go-swagger; DO NOT EDIT.

package artist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetStatsOKCode is the HTTP code returned for type GetStatsOK
const GetStatsOKCode int = 200

/*
GetStatsOK Success

swagger:response getStatsOK
*/
type GetStatsOK struct {
}

// NewGetStatsOK creates GetStatsOK with default headers values
func NewGetStatsOK() *GetStatsOK {

	return &GetStatsOK{}
}

// WriteResponse to the client
func (o *GetStatsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetStatsInternalServerErrorCode is the HTTP code returned for type GetStatsInternalServerError
const GetStatsInternalServerErrorCode int = 500

/*
GetStatsInternalServerError Internal error

swagger:response getStatsInternalServerError
*/
type GetStatsInternalServerError struct {
}

// NewGetStatsInternalServerError creates GetStatsInternalServerError with default headers values
func NewGetStatsInternalServerError() *GetStatsInternalServerError {

	return &GetStatsInternalServerError{}
}

// WriteResponse to the client
func (o *GetStatsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}