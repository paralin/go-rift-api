// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlayerInfoDTO player info d t o
// swagger:model PlayerInfoDTO
type PlayerInfoDTO struct {

	// player Id
	PlayerID int64 `json:"playerId,omitempty"`

	// position
	Position string `json:"position,omitempty"`

	// role
	Role Role `json:"role,omitempty"`

	// team Id
	TeamID string `json:"teamId,omitempty"`
}

// Validate validates this player info d t o
func (m *PlayerInfoDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerInfoDTO) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if err := m.Role.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("role")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayerInfoDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerInfoDTO) UnmarshalBinary(b []byte) error {
	var res PlayerInfoDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
