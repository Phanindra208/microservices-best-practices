// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UsersFeed list of Users structure
//
// swagger:model users_feed
type UsersFeed struct {

	// my feed
	MyFeed []*FeedPost `json:"my_feed"`

	// page state
	PageState *PaginationData `json:"page_state,omitempty"`
}

// Validate validates this users feed
func (m *UsersFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMyFeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsersFeed) validateMyFeed(formats strfmt.Registry) error {

	if swag.IsZero(m.MyFeed) { // not required
		return nil
	}

	for i := 0; i < len(m.MyFeed); i++ {
		if swag.IsZero(m.MyFeed[i]) { // not required
			continue
		}

		if m.MyFeed[i] != nil {
			if err := m.MyFeed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("my_feed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UsersFeed) validatePageState(formats strfmt.Registry) error {

	if swag.IsZero(m.PageState) { // not required
		return nil
	}

	if m.PageState != nil {
		if err := m.PageState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page_state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsersFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsersFeed) UnmarshalBinary(b []byte) error {
	var res UsersFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
