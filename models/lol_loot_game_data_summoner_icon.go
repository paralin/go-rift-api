// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLootGameDataSummonerIcon lol loot game data summoner icon
// swagger:model LolLootGameDataSummonerIcon
type LolLootGameDataSummonerIcon struct {

	// icon path
	IconPath string `json:"iconPath,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this lol loot game data summoner icon
func (m *LolLootGameDataSummonerIcon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLootGameDataSummonerIcon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLootGameDataSummonerIcon) UnmarshalBinary(b []byte) error {
	var res LolLootGameDataSummonerIcon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
