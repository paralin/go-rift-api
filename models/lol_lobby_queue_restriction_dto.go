// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLobbyQueueRestrictionDto lol lobby queue restriction dto
// swagger:model LolLobbyQueueRestrictionDto
type LolLobbyQueueRestrictionDto struct {

	// gatekeeper restrictions
	GatekeeperRestrictions []*LolLobbyGatekeeperRestrictionDto `json:"gatekeeperRestrictions"`
}

// Validate validates this lol lobby queue restriction dto
func (m *LolLobbyQueueRestrictionDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGatekeeperRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyQueueRestrictionDto) validateGatekeeperRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.GatekeeperRestrictions) { // not required
		return nil
	}

	for i := 0; i < len(m.GatekeeperRestrictions); i++ {
		if swag.IsZero(m.GatekeeperRestrictions[i]) { // not required
			continue
		}

		if m.GatekeeperRestrictions[i] != nil {
			if err := m.GatekeeperRestrictions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gatekeeperRestrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyQueueRestrictionDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyQueueRestrictionDto) UnmarshalBinary(b []byte) error {
	var res LolLobbyQueueRestrictionDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
