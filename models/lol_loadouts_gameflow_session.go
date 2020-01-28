// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLoadoutsGameflowSession lol loadouts gameflow session
// swagger:model LolLoadoutsGameflowSession
type LolLoadoutsGameflowSession struct {

	// game data
	GameData *LolLoadoutsGameflowGameData `json:"gameData,omitempty"`

	// phase
	Phase LolLoadoutsGameflowPhase `json:"phase,omitempty"`
}

// Validate validates this lol loadouts gameflow session
func (m *LolLoadoutsGameflowSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLoadoutsGameflowSession) validateGameData(formats strfmt.Registry) error {

	if swag.IsZero(m.GameData) { // not required
		return nil
	}

	if m.GameData != nil {
		if err := m.GameData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameData")
			}
			return err
		}
	}

	return nil
}

func (m *LolLoadoutsGameflowSession) validatePhase(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolLoadoutsGameflowSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoadoutsGameflowSession) UnmarshalBinary(b []byte) error {
	var res LolLoadoutsGameflowSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
