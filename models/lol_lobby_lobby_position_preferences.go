// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyLobbyPositionPreferences lol lobby lobby position preferences
// swagger:model LolLobbyLobbyPositionPreferences
type LolLobbyLobbyPositionPreferences struct {

	// first preference
	FirstPreference string `json:"firstPreference,omitempty"`

	// second preference
	SecondPreference string `json:"secondPreference,omitempty"`
}

// Validate validates this lol lobby lobby position preferences
func (m *LolLobbyLobbyPositionPreferences) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyLobbyPositionPreferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyLobbyPositionPreferences) UnmarshalBinary(b []byte) error {
	var res LolLobbyLobbyPositionPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
