// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyCollectionsChampion lol lobby collections champion
// swagger:model LolLobbyCollectionsChampion
type LolLobbyCollectionsChampion struct {

	// active
	Active bool `json:"active,omitempty"`

	// bot enabled
	BotEnabled bool `json:"botEnabled,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this lol lobby collections champion
func (m *LolLobbyCollectionsChampion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyCollectionsChampion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyCollectionsChampion) UnmarshalBinary(b []byte) error {
	var res LolLobbyCollectionsChampion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
