// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaginationData pagination data
//
// swagger:model pagination_data
type PaginationData struct {

	// page
	Page int64 `json:"page,omitempty"`

	// per page
	PerPage int64 `json:"perPage,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`

	// total page
	TotalPage int64 `json:"totalPage,omitempty"`
}

// Validate validates this pagination data
func (m *PaginationData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationData) UnmarshalBinary(b []byte) error {
	var res PaginationData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
