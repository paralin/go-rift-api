// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolMatchmakingMatchmakingDodgeData lol matchmaking matchmaking dodge data
// swagger:model LolMatchmakingMatchmakingDodgeData
type LolMatchmakingMatchmakingDodgeData struct {

	// dodger Id
	DodgerID int64 `json:"dodgerId,omitempty"`

	// state
	State LolMatchmakingMatchmakingDodgeState `json:"state,omitempty"`
}

// Validate validates this lol matchmaking matchmaking dodge data
func (m *LolMatchmakingMatchmakingDodgeData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolMatchmakingMatchmakingDodgeData) validateState(formats strfmt.Registry) error {

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
func (m *LolMatchmakingMatchmakingDodgeData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMatchmakingMatchmakingDodgeData) UnmarshalBinary(b []byte) error {
	var res LolMatchmakingMatchmakingDodgeData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
