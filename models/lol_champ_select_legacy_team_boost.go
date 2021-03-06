// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectLegacyTeamBoost lol champ select legacy team boost
// swagger:model LolChampSelectLegacyTeamBoost
type LolChampSelectLegacyTeamBoost struct {

	// available skins
	AvailableSkins []int64 `json:"availableSkins"`

	// ip reward
	IPReward int64 `json:"ipReward,omitempty"`

	// ip reward for purchaser
	IPRewardForPurchaser int64 `json:"ipRewardForPurchaser,omitempty"`

	// price
	Price int64 `json:"price,omitempty"`

	// skin unlock mode
	SkinUnlockMode string `json:"skinUnlockMode,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`

	// unlocked
	Unlocked bool `json:"unlocked,omitempty"`
}

// Validate validates this lol champ select legacy team boost
func (m *LolChampSelectLegacyTeamBoost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectLegacyTeamBoost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectLegacyTeamBoost) UnmarshalBinary(b []byte) error {
	var res LolChampSelectLegacyTeamBoost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
