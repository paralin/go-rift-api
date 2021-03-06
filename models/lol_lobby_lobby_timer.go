// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyLobbyTimer lol lobby lobby timer
// swagger:model LolLobbyLobbyTimer
type LolLobbyLobbyTimer struct {

	// countdown
	Countdown int64 `json:"countdown,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this lol lobby lobby timer
func (m *LolLobbyLobbyTimer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyLobbyTimer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyLobbyTimer) UnmarshalBinary(b []byte) error {
	var res LolLobbyLobbyTimer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
