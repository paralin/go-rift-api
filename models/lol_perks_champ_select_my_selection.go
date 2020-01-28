// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPerksChampSelectMySelection lol perks champ select my selection
// swagger:model LolPerksChampSelectMySelection
type LolPerksChampSelectMySelection struct {

	// selected skin Id
	SelectedSkinID int32 `json:"selectedSkinId,omitempty"`

	// spell1 Id
	Spell1ID int64 `json:"spell1Id,omitempty"`

	// spell2 Id
	Spell2ID int64 `json:"spell2Id,omitempty"`

	// ward skin Id
	WardSkinID int64 `json:"wardSkinId,omitempty"`
}

// Validate validates this lol perks champ select my selection
func (m *LolPerksChampSelectMySelection) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksChampSelectMySelection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksChampSelectMySelection) UnmarshalBinary(b []byte) error {
	var res LolPerksChampSelectMySelection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}