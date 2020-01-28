// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatChampSelection lol chat champ selection
// swagger:model LolChatChampSelection
type LolChatChampSelection struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// selected skin index
	SelectedSkinIndex int32 `json:"selectedSkinIndex,omitempty"`

	// summoner internal name
	SummonerInternalName string `json:"summonerInternalName,omitempty"`
}

// Validate validates this lol chat champ selection
func (m *LolChatChampSelection) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatChampSelection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatChampSelection) UnmarshalBinary(b []byte) error {
	var res LolChatChampSelection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
