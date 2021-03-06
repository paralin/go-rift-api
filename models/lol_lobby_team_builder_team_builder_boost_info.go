// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderTeamBuilderBoostInfo lol lobby team builder team builder boost info
// swagger:model LolLobbyTeamBuilderTeamBuilderBoostInfo
type LolLobbyTeamBuilderTeamBuilderBoostInfo struct {

	// activator cell Id
	ActivatorCellID int64 `json:"activatorCellId,omitempty"`

	// allow battle boost
	AllowBattleBoost bool `json:"allowBattleBoost,omitempty"`

	// battle boost activated
	BattleBoostActivated bool `json:"battleBoostActivated,omitempty"`

	// boostable skin count
	BoostableSkinCount int32 `json:"boostableSkinCount,omitempty"`

	// cost
	Cost int64 `json:"cost,omitempty"`

	// unlocked skin ids
	UnlockedSkinIds []int64 `json:"unlockedSkinIds"`
}

// Validate validates this lol lobby team builder team builder boost info
func (m *LolLobbyTeamBuilderTeamBuilderBoostInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTeamBuilderBoostInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTeamBuilderBoostInfo) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderTeamBuilderBoostInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
