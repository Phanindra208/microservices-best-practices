// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorResponse Error response body for 4xx responses
//
// swagger:model ErrorResponse
type ErrorResponse struct {

	// 6xx if validation error.
	// 1000+ for special errors.
	// Otherwise, normal HTTP status codes apply.
	//
	Code int64 `json:"code,omitempty"`

	// Informative text - not for error checking
	Message string `json:"message,omitempty"`

	// validation
	Validation *ErrorResponseValidation `json:"validation,omitempty"`
}

// Validate validates this error response
func (m *ErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponse) validateValidation(formats strfmt.Registry) error {

	if swag.IsZero(m.Validation) { // not required
		return nil
	}

	if m.Validation != nil {
		if err := m.Validation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponse) UnmarshalBinary(b []byte) error {
	var res ErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ErrorResponseValidation Details about validation error
//
// swagger:model ErrorResponseValidation
type ErrorResponseValidation struct {

	// The parameter name which caused the error
	Param string `json:"param,omitempty"`

	// Error code for validation error - like error_code
	Code string `json:"code,omitempty"`
}

// Validate validates this error response validation
func (m *ErrorResponseValidation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseValidation) UnmarshalBinary(b []byte) error {
	var res ErrorResponseValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
