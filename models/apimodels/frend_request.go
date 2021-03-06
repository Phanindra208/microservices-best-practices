// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FrendRequest User structure
//
// swagger:model frend_request
type FrendRequest struct {

	// source user email id
	SourceUserEmailID string `json:"source_user_email_id,omitempty"`

	// target user email id
	TargetUserEmailID string `json:"target_user_email_id,omitempty"`
}

// Validate validates this frend request
func (m *FrendRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FrendRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrendRequest) UnmarshalBinary(b []byte) error {
	var res FrendRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
