// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderQueueReward lol lobby team builder queue reward
// swagger:model LolLobbyTeamBuilderQueueReward
type LolLobbyTeamBuilderQueueReward struct {

	// is champion points enabled
	IsChampionPointsEnabled bool `json:"isChampionPointsEnabled,omitempty"`

	// is Ip enabled
	IsIPEnabled bool `json:"isIpEnabled,omitempty"`

	// is xp enabled
	IsXpEnabled bool `json:"isXpEnabled,omitempty"`

	// party size Ip rewards
	PartySizeIPRewards []int32 `json:"partySizeIpRewards"`
}

// Validate validates this lol lobby team builder queue reward
func (m *LolLobbyTeamBuilderQueueReward) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderQueueReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderQueueReward) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderQueueReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
