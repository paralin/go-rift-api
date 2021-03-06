// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEndOfGameLobbyInvitation lol end of game lobby invitation
// swagger:model LolEndOfGameLobbyInvitation
type LolEndOfGameLobbyInvitation struct {

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this lol end of game lobby invitation
func (m *LolEndOfGameLobbyInvitation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameLobbyInvitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameLobbyInvitation) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameLobbyInvitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
