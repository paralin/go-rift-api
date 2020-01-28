// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLoadoutsCreateLoadoutRequestDTO lol loadouts create loadout request d t o
// swagger:model LolLoadoutsCreateLoadoutRequestDTO
type LolLoadoutsCreateLoadoutRequestDTO struct {

	// loadout
	Loadout *LolLoadoutsCreateLoadoutDTO `json:"loadout,omitempty"`

	// service to jwts map
	ServiceToJwtsMap map[string]interface{} `json:"serviceToJwtsMap,omitempty"`
}

// Validate validates this lol loadouts create loadout request d t o
func (m *LolLoadoutsCreateLoadoutRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLoadoutsCreateLoadoutRequestDTO) validateLoadout(formats strfmt.Registry) error {

	if swag.IsZero(m.Loadout) { // not required
		return nil
	}

	if m.Loadout != nil {
		if err := m.Loadout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadout")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLoadoutsCreateLoadoutRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoadoutsCreateLoadoutRequestDTO) UnmarshalBinary(b []byte) error {
	var res LolLoadoutsCreateLoadoutRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}