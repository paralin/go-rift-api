// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderTbdInventory lol lobby team builder tbd inventory
// swagger:model LolLobbyTeamBuilderTbdInventory
type LolLobbyTeamBuilderTbdInventory struct {

	// all champion ids
	AllChampionIds []int32 `json:"allChampionIds"`

	// disabled champion ids
	DisabledChampionIds []int32 `json:"disabledChampionIds"`

	// initial spell ids
	InitialSpellIds []int32 `json:"initialSpellIds"`

	// last selected skin Id by champion Id
	LastSelectedSkinIDByChampionID map[string]int32 `json:"lastSelectedSkinIdByChampionId,omitempty"`

	// skin ids
	SkinIds []int32 `json:"skinIds"`

	// spell ids
	SpellIds []int32 `json:"spellIds"`
}

// Validate validates this lol lobby team builder tbd inventory
func (m *LolLobbyTeamBuilderTbdInventory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTbdInventory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTbdInventory) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderTbdInventory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}