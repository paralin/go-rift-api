// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolPerksSummoner lol perks summoner
// swagger:model LolPerksSummoner
type LolPerksSummoner struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// internal name
	InternalName string `json:"internalName,omitempty"`

	// percent complete for next level
	PercentCompleteForNextLevel int32 `json:"percentCompleteForNextLevel,omitempty"`

	// profile icon Id
	ProfileIconID int32 `json:"profileIconId,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`

	// reroll points
	RerollPoints *LolPerksSummonerRerollPoints `json:"rerollPoints,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner level
	SummonerLevel int32 `json:"summonerLevel,omitempty"`

	// xp since last level
	XpSinceLastLevel int64 `json:"xpSinceLastLevel,omitempty"`

	// xp until next level
	XpUntilNextLevel int64 `json:"xpUntilNextLevel,omitempty"`
}

// Validate validates this lol perks summoner
func (m *LolPerksSummoner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRerollPoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPerksSummoner) validateRerollPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.RerollPoints) { // not required
		return nil
	}

	if m.RerollPoints != nil {
		if err := m.RerollPoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rerollPoints")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksSummoner) UnmarshalBinary(b []byte) error {
	var res LolPerksSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}