// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyRegistrationCredentials lol lobby registration credentials
// swagger:model LolLobbyRegistrationCredentials
type LolLobbyRegistrationCredentials struct {

	// game client version
	GameClientVersion string `json:"gameClientVersion,omitempty"`

	// inventory token
	InventoryToken string `json:"inventoryToken,omitempty"`

	// inventory tokens
	InventoryTokens []string `json:"inventoryTokens"`

	// ranked overview token
	RankedOverviewToken string `json:"rankedOverviewToken,omitempty"`

	// simple inventory token
	SimpleInventoryToken string `json:"simpleInventoryToken,omitempty"`

	// summoner token
	SummonerToken string `json:"summonerToken,omitempty"`

	// user info token
	UserInfoToken string `json:"userInfoToken,omitempty"`
}

// Validate validates this lol lobby registration credentials
func (m *LolLobbyRegistrationCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyRegistrationCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyRegistrationCredentials) UnmarshalBinary(b []byte) error {
	var res LolLobbyRegistrationCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}