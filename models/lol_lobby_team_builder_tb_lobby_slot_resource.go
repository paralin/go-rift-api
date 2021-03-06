// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderTbLobbySlotResource lol lobby team builder tb lobby slot resource
// swagger:model LolLobbyTeamBuilderTbLobbySlotResource
type LolLobbyTeamBuilderTbLobbySlotResource struct {

	// auto fill eligible
	AutoFillEligible bool `json:"autoFillEligible,omitempty"`

	// auto fill protected for promos
	AutoFillProtectedForPromos bool `json:"autoFillProtectedForPromos,omitempty"`

	// auto fill protected for soloing
	AutoFillProtectedForSoloing bool `json:"autoFillProtectedForSoloing,omitempty"`

	// auto fill protected for streaking
	AutoFillProtectedForStreaking bool `json:"autoFillProtectedForStreaking,omitempty"`

	// draft position preferences
	DraftPositionPreferences []string `json:"draftPositionPreferences"`

	// excluded position preference
	ExcludedPositionPreference string `json:"excludedPositionPreference,omitempty"`

	// show position excluder
	ShowPositionExcluder bool `json:"showPositionExcluder,omitempty"`

	// slot Id
	SlotID int32 `json:"slotId,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol lobby team builder tb lobby slot resource
func (m *LolLobbyTeamBuilderTbLobbySlotResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTbLobbySlotResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTbLobbySlotResource) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderTbLobbySlotResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
