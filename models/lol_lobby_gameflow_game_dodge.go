// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLobbyGameflowGameDodge lol lobby gameflow game dodge
// swagger:model LolLobbyGameflowGameDodge
type LolLobbyGameflowGameDodge struct {

	// dodge ids
	DodgeIds []int64 `json:"dodgeIds"`

	// phase
	Phase LolLobbyGameflowPhase `json:"phase,omitempty"`

	// state
	State LolLobbyMatchmakingDodgeState `json:"state,omitempty"`
}

// Validate validates this lol lobby gameflow game dodge
func (m *LolLobbyGameflowGameDodge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyGameflowGameDodge) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if err := m.Phase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("phase")
		}
		return err
	}

	return nil
}

func (m *LolLobbyGameflowGameDodge) validateState(formats strfmt.Registry) error {

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
func (m *LolLobbyGameflowGameDodge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyGameflowGameDodge) UnmarshalBinary(b []byte) error {
	var res LolLobbyGameflowGameDodge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
