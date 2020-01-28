// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEndOfGameGameDataTftChampion lol end of game game data tft champion
// swagger:model LolEndOfGameGameDataTftChampion
type LolEndOfGameGameDataTftChampion struct {

	// character id
	CharacterID string `json:"character_id,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// rarity
	Rarity int32 `json:"rarity,omitempty"`

	// square icon path
	SquareIconPath string `json:"squareIconPath,omitempty"`
}

// Validate validates this lol end of game game data tft champion
func (m *LolEndOfGameGameDataTftChampion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameGameDataTftChampion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameGameDataTftChampion) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameGameDataTftChampion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
