// Code generated by go-swagger; DO NOT EDIT.

package manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// FetchStatsOKCode is the HTTP code returned for type FetchStatsOK
const FetchStatsOKCode int = 200

/*
FetchStatsOK Success

swagger:response fetchStatsOK
*/
type FetchStatsOK struct {
}

// NewFetchStatsOK creates FetchStatsOK with default headers values
func NewFetchStatsOK() *FetchStatsOK {

	return &FetchStatsOK{}
}

// WriteResponse to the client
func (o *FetchStatsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// FetchStatsInternalServerErrorCode is the HTTP code returned for type FetchStatsInternalServerError
const FetchStatsInternalServerErrorCode int = 500

/*
FetchStatsInternalServerError Internal error

swagger:response fetchStatsInternalServerError
*/
type FetchStatsInternalServerError struct {
}

// NewFetchStatsInternalServerError creates FetchStatsInternalServerError with default headers values
func NewFetchStatsInternalServerError() *FetchStatsInternalServerError {

	return &FetchStatsInternalServerError{}
}

// WriteResponse to the client
func (o *FetchStatsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}