// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ChampSelectLcdsPlayerChampionSelectionDTO champ select lcds player champion selection d t o
// swagger:model ChampSelectLcdsPlayerChampionSelectionDTO
type ChampSelectLcdsPlayerChampionSelectionDTO struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// selected skin index
	SelectedSkinIndex int32 `json:"selectedSkinIndex,omitempty"`

	// spell1 Id
	Spell1ID int32 `json:"spell1Id,omitempty"`

	// spell2 Id
	Spell2ID int32 `json:"spell2Id,omitempty"`

	// summoner internal name
	SummonerInternalName string `json:"summonerInternalName,omitempty"`
}

// Validate validates this champ select lcds player champion selection d t o
func (m *ChampSelectLcdsPlayerChampionSelectionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChampSelectLcdsPlayerChampionSelectionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChampSelectLcdsPlayerChampionSelectionDTO) UnmarshalBinary(b []byte) error {
	var res ChampSelectLcdsPlayerChampionSelectionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
