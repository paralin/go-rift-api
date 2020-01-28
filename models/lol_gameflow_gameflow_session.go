// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolGameflowGameflowSession lol gameflow gameflow session
// swagger:model LolGameflowGameflowSession
type LolGameflowGameflowSession struct {

	// game client
	GameClient *LolGameflowGameflowGameClient `json:"gameClient,omitempty"`

	// game data
	GameData *LolGameflowGameflowGameData `json:"gameData,omitempty"`

	// game dodge
	GameDodge *LolGameflowGameflowGameDodge `json:"gameDodge,omitempty"`

	// map
	Map *LolGameflowGameflowGameMap `json:"map,omitempty"`

	// phase
	Phase LolGameflowGameflowPhase `json:"phase,omitempty"`
}

// Validate validates this lol gameflow gameflow session
func (m *LolGameflowGameflowSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameDodge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMap(formats); err != nil {
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

func (m *LolGameflowGameflowSession) validateGameClient(formats strfmt.Registry) error {

	if swag.IsZero(m.GameClient) { // not required
		return nil
	}

	if m.GameClient != nil {
		if err := m.GameClient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameClient")
			}
			return err
		}
	}

	return nil
}

func (m *LolGameflowGameflowSession) validateGameData(formats strfmt.Registry) error {

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

func (m *LolGameflowGameflowSession) validateGameDodge(formats strfmt.Registry) error {

	if swag.IsZero(m.GameDodge) { // not required
		return nil
	}

	if m.GameDodge != nil {
		if err := m.GameDodge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameDodge")
			}
			return err
		}
	}

	return nil
}

func (m *LolGameflowGameflowSession) validateMap(formats strfmt.Registry) error {

	if swag.IsZero(m.Map) { // not required
		return nil
	}

	if m.Map != nil {
		if err := m.Map.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

func (m *LolGameflowGameflowSession) validatePhase(formats strfmt.Registry) error {

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
func (m *LolGameflowGameflowSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGameflowGameflowSession) UnmarshalBinary(b []byte) error {
	var res LolGameflowGameflowSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
