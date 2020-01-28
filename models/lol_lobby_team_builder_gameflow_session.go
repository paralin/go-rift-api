// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderGameflowSession lol lobby team builder gameflow session
// swagger:model LolLobbyTeamBuilderGameflowSession
type LolLobbyTeamBuilderGameflowSession struct {

	// game client
	GameClient *LolLobbyTeamBuilderGameflowGameClient `json:"gameClient,omitempty"`
}

// Validate validates this lol lobby team builder gameflow session
func (m *LolLobbyTeamBuilderGameflowSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameClient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyTeamBuilderGameflowSession) validateGameClient(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderGameflowSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderGameflowSession) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderGameflowSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
