// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyLcdsPartyRewardsConfig lol lobby lcds party rewards config
// swagger:model LolLobbyLcdsPartyRewardsConfig
type LolLobbyLcdsPartyRewardsConfig struct {

	// enabled
	Enabled bool `json:"Enabled,omitempty"`
}

// Validate validates this lol lobby lcds party rewards config
func (m *LolLobbyLcdsPartyRewardsConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyLcdsPartyRewardsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyLcdsPartyRewardsConfig) UnmarshalBinary(b []byte) error {
	var res LolLobbyLcdsPartyRewardsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}