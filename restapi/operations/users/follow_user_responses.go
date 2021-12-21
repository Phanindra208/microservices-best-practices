// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iAmPlus/microservice/models/apimodels"
)

// FollowUserOKCode is the HTTP code returned for type FollowUserOK
const FollowUserOKCode int = 200

/*FollowUserOK No content

swagger:response followUserOK
*/
type FollowUserOK struct {
}

// NewFollowUserOK creates FollowUserOK with default headers values
func NewFollowUserOK() *FollowUserOK {

	return &FollowUserOK{}
}

// WriteResponse to the client
func (o *FollowUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// FollowUserBadRequestCode is the HTTP code returned for type FollowUserBadRequest
const FollowUserBadRequestCode int = 400

/*FollowUserBadRequest Bad request <br>


swagger:response followUserBadRequest
*/
type FollowUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewFollowUserBadRequest creates FollowUserBadRequest with default headers values
func NewFollowUserBadRequest() *FollowUserBadRequest {

	return &FollowUserBadRequest{}
}

// WithPayload adds the payload to the follow user bad request response
func (o *FollowUserBadRequest) WithPayload(payload *apimodels.ErrorResponse) *FollowUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the follow user bad request response
func (o *FollowUserBadRequest) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FollowUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FollowUserUnauthorizedCode is the HTTP code returned for type FollowUserUnauthorized
const FollowUserUnauthorizedCode int = 401

/*FollowUserUnauthorized Unauthorized

swagger:response followUserUnauthorized
*/
type FollowUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewFollowUserUnauthorized creates FollowUserUnauthorized with default headers values
func NewFollowUserUnauthorized() *FollowUserUnauthorized {

	return &FollowUserUnauthorized{}
}

// WithPayload adds the payload to the follow user unauthorized response
func (o *FollowUserUnauthorized) WithPayload(payload *apimodels.ErrorResponse) *FollowUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the follow user unauthorized response
func (o *FollowUserUnauthorized) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FollowUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FollowUserInternalServerErrorCode is the HTTP code returned for type FollowUserInternalServerError
const FollowUserInternalServerErrorCode int = 500

/*FollowUserInternalServerError Internal error

swagger:response followUserInternalServerError
*/
type FollowUserInternalServerError struct {
}

// NewFollowUserInternalServerError creates FollowUserInternalServerError with default headers values
func NewFollowUserInternalServerError() *FollowUserInternalServerError {

	return &FollowUserInternalServerError{}
}

// WriteResponse to the client
func (o *FollowUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
