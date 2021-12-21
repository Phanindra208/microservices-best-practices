// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iAmPlus/microservice/models/apimodels"
)

// UnFollowUserOKCode is the HTTP code returned for type UnFollowUserOK
const UnFollowUserOKCode int = 200

/*UnFollowUserOK No content

swagger:response unFollowUserOK
*/
type UnFollowUserOK struct {
}

// NewUnFollowUserOK creates UnFollowUserOK with default headers values
func NewUnFollowUserOK() *UnFollowUserOK {

	return &UnFollowUserOK{}
}

// WriteResponse to the client
func (o *UnFollowUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UnFollowUserBadRequestCode is the HTTP code returned for type UnFollowUserBadRequest
const UnFollowUserBadRequestCode int = 400

/*UnFollowUserBadRequest Bad request <br>


swagger:response unFollowUserBadRequest
*/
type UnFollowUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewUnFollowUserBadRequest creates UnFollowUserBadRequest with default headers values
func NewUnFollowUserBadRequest() *UnFollowUserBadRequest {

	return &UnFollowUserBadRequest{}
}

// WithPayload adds the payload to the un follow user bad request response
func (o *UnFollowUserBadRequest) WithPayload(payload *apimodels.ErrorResponse) *UnFollowUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the un follow user bad request response
func (o *UnFollowUserBadRequest) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnFollowUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnFollowUserUnauthorizedCode is the HTTP code returned for type UnFollowUserUnauthorized
const UnFollowUserUnauthorizedCode int = 401

/*UnFollowUserUnauthorized Unauthorized

swagger:response unFollowUserUnauthorized
*/
type UnFollowUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewUnFollowUserUnauthorized creates UnFollowUserUnauthorized with default headers values
func NewUnFollowUserUnauthorized() *UnFollowUserUnauthorized {

	return &UnFollowUserUnauthorized{}
}

// WithPayload adds the payload to the un follow user unauthorized response
func (o *UnFollowUserUnauthorized) WithPayload(payload *apimodels.ErrorResponse) *UnFollowUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the un follow user unauthorized response
func (o *UnFollowUserUnauthorized) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnFollowUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnFollowUserInternalServerErrorCode is the HTTP code returned for type UnFollowUserInternalServerError
const UnFollowUserInternalServerErrorCode int = 500

/*UnFollowUserInternalServerError Internal error

swagger:response unFollowUserInternalServerError
*/
type UnFollowUserInternalServerError struct {
}

// NewUnFollowUserInternalServerError creates UnFollowUserInternalServerError with default headers values
func NewUnFollowUserInternalServerError() *UnFollowUserInternalServerError {

	return &UnFollowUserInternalServerError{}
}

// WriteResponse to the client
func (o *UnFollowUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
