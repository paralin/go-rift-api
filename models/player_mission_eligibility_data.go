// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlayerMissionEligibilityData player mission eligibility data
// swagger:model PlayerMissionEligibilityData
type PlayerMissionEligibilityData struct {

	// level
	Level int32 `json:"level,omitempty"`

	// loyalty enabled
	LoyaltyEnabled bool `json:"loyaltyEnabled,omitempty"`

	// player inventory
	PlayerInventory *PlayerInventory `json:"playerInventory,omitempty"`

	// user info token
	UserInfoToken string `json:"userInfoToken,omitempty"`
}

// Validate validates this player mission eligibility data
func (m *PlayerMissionEligibilityData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayerInventory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerMissionEligibilityData) validatePlayerInventory(formats strfmt.Registry) error {

	if swag.IsZero(m.PlayerInventory) { // not required
		return nil
	}

	if m.PlayerInventory != nil {
		if err := m.PlayerInventory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("playerInventory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayerMissionEligibilityData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerMissionEligibilityData) UnmarshalBinary(b []byte) error {
	var res PlayerMissionEligibilityData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
