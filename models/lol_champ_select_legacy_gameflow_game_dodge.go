// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolChampSelectLegacyGameflowGameDodge lol champ select legacy gameflow game dodge
// swagger:model LolChampSelectLegacyGameflowGameDodge
type LolChampSelectLegacyGameflowGameDodge struct {

	// dodge ids
	DodgeIds []int64 `json:"dodgeIds"`

	// state
	State LolChampSelectLegacyGameflowGameDodgeState `json:"state,omitempty"`
}

// Validate validates this lol champ select legacy gameflow game dodge
func (m *LolChampSelectLegacyGameflowGameDodge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChampSelectLegacyGameflowGameDodge) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectLegacyGameflowGameDodge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectLegacyGameflowGameDodge) UnmarshalBinary(b []byte) error {
	var res LolChampSelectLegacyGameflowGameDodge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}