// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLobbyGameflowSession lol lobby gameflow session
// swagger:model LolLobbyGameflowSession
type LolLobbyGameflowSession struct {

	// game client
	GameClient *LolLobbyGameflowGameClient `json:"gameClient,omitempty"`

	// game data
	GameData *LolLobbyGameflowGameData `json:"gameData,omitempty"`

	// game dodge
	GameDodge *LolLobbyGameflowGameDodge `json:"gameDodge,omitempty"`

	// phase
	Phase LolLobbyGameflowPhase `json:"phase,omitempty"`
}

// Validate validates this lol lobby gameflow session
func (m *LolLobbyGameflowSession) Validate(formats strfmt.Registry) error {
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

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyGameflowSession) validateGameClient(formats strfmt.Registry) error {

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

func (m *LolLobbyGameflowSession) validateGameData(formats strfmt.Registry) error {

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

func (m *LolLobbyGameflowSession) validateGameDodge(formats strfmt.Registry) error {

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

func (m *LolLobbyGameflowSession) validatePhase(formats strfmt.Registry) error {

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
func (m *LolLobbyGameflowSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyGameflowSession) UnmarshalBinary(b []byte) error {
	var res LolLobbyGameflowSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
