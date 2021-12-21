// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iAmPlus/microservice/models/apimodels"
)

// CreateUserOKCode is the HTTP code returned for type CreateUserOK
const CreateUserOKCode int = 200

/*CreateUserOK No content

swagger:response createUserOK
*/
type CreateUserOK struct {
}

// NewCreateUserOK creates CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {

	return &CreateUserOK{}
}

// WriteResponse to the client
func (o *CreateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CreateUserBadRequestCode is the HTTP code returned for type CreateUserBadRequest
const CreateUserBadRequestCode int = 400

/*CreateUserBadRequest Bad request <br>


swagger:response createUserBadRequest
*/
type CreateUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewCreateUserBadRequest creates CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {

	return &CreateUserBadRequest{}
}

// WithPayload adds the payload to the create user bad request response
func (o *CreateUserBadRequest) WithPayload(payload *apimodels.ErrorResponse) *CreateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user bad request response
func (o *CreateUserBadRequest) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUserUnauthorizedCode is the HTTP code returned for type CreateUserUnauthorized
const CreateUserUnauthorizedCode int = 401

/*CreateUserUnauthorized Unauthorized

swagger:response createUserUnauthorized
*/
type CreateUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewCreateUserUnauthorized creates CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {

	return &CreateUserUnauthorized{}
}

// WithPayload adds the payload to the create user unauthorized response
func (o *CreateUserUnauthorized) WithPayload(payload *apimodels.ErrorResponse) *CreateUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user unauthorized response
func (o *CreateUserUnauthorized) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUserInternalServerErrorCode is the HTTP code returned for type CreateUserInternalServerError
const CreateUserInternalServerErrorCode int = 500

/*CreateUserInternalServerError Internal error

swagger:response createUserInternalServerError
*/
type CreateUserInternalServerError struct {
}

// NewCreateUserInternalServerError creates CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {

	return &CreateUserInternalServerError{}
}

// WriteResponse to the client
func (o *CreateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}