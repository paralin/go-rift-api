// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLootCollectionsOwnership lol loot collections ownership
// swagger:model LolLootCollectionsOwnership
type LolLootCollectionsOwnership struct {

	// free to play reward
	FreeToPlayReward bool `json:"freeToPlayReward,omitempty"`

	// owned
	Owned bool `json:"owned,omitempty"`

	// rental
	Rental *LolLootCollectionsRental `json:"rental,omitempty"`
}

// Validate validates this lol loot collections ownership
func (m *LolLootCollectionsOwnership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRental(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLootCollectionsOwnership) validateRental(formats strfmt.Registry) error {

	if swag.IsZero(m.Rental) { // not required
		return nil
	}

	if m.Rental != nil {
		if err := m.Rental.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rental")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLootCollectionsOwnership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLootCollectionsOwnership) UnmarshalBinary(b []byte) error {
	var res LolLootCollectionsOwnership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
