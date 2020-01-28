// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ActiveBoostsLcdsSummonerActiveBoostsDTO active boosts lcds summoner active boosts d t o
// swagger:model ActiveBoostsLcdsSummonerActiveBoostsDTO
type ActiveBoostsLcdsSummonerActiveBoostsDTO struct {

	// ip boost end date
	IPBoostEndDate int64 `json:"ipBoostEndDate,omitempty"`

	// ip boost per win count
	IPBoostPerWinCount int32 `json:"ipBoostPerWinCount,omitempty"`

	// ip loyalty boost
	IPLoyaltyBoost int32 `json:"ipLoyaltyBoost,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// xp boost end date
	XpBoostEndDate int64 `json:"xpBoostEndDate,omitempty"`

	// xp boost per win count
	XpBoostPerWinCount int32 `json:"xpBoostPerWinCount,omitempty"`

	// xp loyalty boost
	XpLoyaltyBoost int32 `json:"xpLoyaltyBoost,omitempty"`
}

// Validate validates this active boosts lcds summoner active boosts d t o
func (m *ActiveBoostsLcdsSummonerActiveBoostsDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveBoostsLcdsSummonerActiveBoostsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveBoostsLcdsSummonerActiveBoostsDTO) UnmarshalBinary(b []byte) error {
	var res ActiveBoostsLcdsSummonerActiveBoostsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
