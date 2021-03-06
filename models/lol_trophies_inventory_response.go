// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolTrophiesInventoryResponse lol trophies inventory response
// swagger:model LolTrophiesInventoryResponse
type LolTrophiesInventoryResponse struct {

	// items
	Items *LolTrophiesInventoryItemsByType `json:"items,omitempty"`
}

// Validate validates this lol trophies inventory response
func (m *LolTrophiesInventoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolTrophiesInventoryResponse) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	if m.Items != nil {
		if err := m.Items.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolTrophiesInventoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolTrophiesInventoryResponse) UnmarshalBinary(b []byte) error {
	var res LolTrophiesInventoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
