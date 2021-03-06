// Code generated by go-swagger; DO NOT EDIT.

package feed

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iAmPlus/microservice/models/apimodels"
)

// GetRelatedFeedOKCode is the HTTP code returned for type GetRelatedFeedOK
const GetRelatedFeedOKCode int = 200

/*GetRelatedFeedOK Successfully got loyalty points

swagger:response getRelatedFeedOK
*/
type GetRelatedFeedOK struct {

	/*
	  In: Body
	*/
	Payload *apimodels.RelatedFeedResponse `json:"body,omitempty"`
}

// NewGetRelatedFeedOK creates GetRelatedFeedOK with default headers values
func NewGetRelatedFeedOK() *GetRelatedFeedOK {

	return &GetRelatedFeedOK{}
}

// WithPayload adds the payload to the get related feed o k response
func (o *GetRelatedFeedOK) WithPayload(payload *apimodels.RelatedFeedResponse) *GetRelatedFeedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get related feed o k response
func (o *GetRelatedFeedOK) SetPayload(payload *apimodels.RelatedFeedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRelatedFeedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRelatedFeedBadRequestCode is the HTTP code returned for type GetRelatedFeedBadRequest
const GetRelatedFeedBadRequestCode int = 400

/*GetRelatedFeedBadRequest Bad request <br>


swagger:response getRelatedFeedBadRequest
*/
type GetRelatedFeedBadRequest struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewGetRelatedFeedBadRequest creates GetRelatedFeedBadRequest with default headers values
func NewGetRelatedFeedBadRequest() *GetRelatedFeedBadRequest {

	return &GetRelatedFeedBadRequest{}
}

// WithPayload adds the payload to the get related feed bad request response
func (o *GetRelatedFeedBadRequest) WithPayload(payload *apimodels.ErrorResponse) *GetRelatedFeedBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get related feed bad request response
func (o *GetRelatedFeedBadRequest) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRelatedFeedBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRelatedFeedUnauthorizedCode is the HTTP code returned for type GetRelatedFeedUnauthorized
const GetRelatedFeedUnauthorizedCode int = 401

/*GetRelatedFeedUnauthorized Unauthorized

swagger:response getRelatedFeedUnauthorized
*/
type GetRelatedFeedUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewGetRelatedFeedUnauthorized creates GetRelatedFeedUnauthorized with default headers values
func NewGetRelatedFeedUnauthorized() *GetRelatedFeedUnauthorized {

	return &GetRelatedFeedUnauthorized{}
}

// WithPayload adds the payload to the get related feed unauthorized response
func (o *GetRelatedFeedUnauthorized) WithPayload(payload *apimodels.ErrorResponse) *GetRelatedFeedUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get related feed unauthorized response
func (o *GetRelatedFeedUnauthorized) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRelatedFeedUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRelatedFeedInternalServerErrorCode is the HTTP code returned for type GetRelatedFeedInternalServerError
const GetRelatedFeedInternalServerErrorCode int = 500

/*GetRelatedFeedInternalServerError Internal error

swagger:response getRelatedFeedInternalServerError
*/
type GetRelatedFeedInternalServerError struct {
}

// NewGetRelatedFeedInternalServerError creates GetRelatedFeedInternalServerError with default headers values
func NewGetRelatedFeedInternalServerError() *GetRelatedFeedInternalServerError {

	return &GetRelatedFeedInternalServerError{}
}

// WriteResponse to the client
func (o *GetRelatedFeedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
