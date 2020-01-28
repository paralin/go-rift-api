// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyGameflowGameClient lol lobby gameflow game client
// swagger:model LolLobbyGameflowGameClient
type LolLobbyGameflowGameClient struct {

	// running
	Running bool `json:"running,omitempty"`
}

// Validate validates this lol lobby gameflow game client
func (m *LolLobbyGameflowGameClient) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyGameflowGameClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyGameflowGameClient) UnmarshalBinary(b []byte) error {
	var res LolLobbyGameflowGameClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
