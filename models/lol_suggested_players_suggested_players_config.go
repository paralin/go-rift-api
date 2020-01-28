// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolSuggestedPlayersSuggestedPlayersConfig lol suggested players suggested players config
// swagger:model LolSuggestedPlayersSuggestedPlayersConfig
type LolSuggestedPlayersSuggestedPlayersConfig struct {

	// enabled
	Enabled bool `json:"Enabled,omitempty"`

	// friends of friends enabled
	FriendsOfFriendsEnabled bool `json:"FriendsOfFriendsEnabled,omitempty"`

	// friends of friends limit
	FriendsOfFriendsLimit int32 `json:"FriendsOfFriendsLimit,omitempty"`

	// max num replacements
	MaxNumReplacements int32 `json:"MaxNumReplacements,omitempty"`

	// max num suggested players
	MaxNumSuggestedPlayers int32 `json:"MaxNumSuggestedPlayers,omitempty"`

	// online friends limit
	OnlineFriendsLimit int32 `json:"OnlineFriendsLimit,omitempty"`

	// previous premades limit
	PreviousPremadesLimit int32 `json:"PreviousPremadesLimit,omitempty"`

	// vicorious comrades limit
	VicoriousComradesLimit int32 `json:"VicoriousComradesLimit,omitempty"`
}

// Validate validates this lol suggested players suggested players config
func (m *LolSuggestedPlayersSuggestedPlayersConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersConfig) UnmarshalBinary(b []byte) error {
	var res LolSuggestedPlayersSuggestedPlayersConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}