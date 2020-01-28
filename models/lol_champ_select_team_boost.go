// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectTeamBoost lol champ select team boost
// swagger:model LolChampSelectTeamBoost
type LolChampSelectTeamBoost struct {

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

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// unlocked
	Unlocked bool `json:"unlocked,omitempty"`
}

// Validate validates this lol champ select team boost
func (m *LolChampSelectTeamBoost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectTeamBoost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectTeamBoost) UnmarshalBinary(b []byte) error {
	var res LolChampSelectTeamBoost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
