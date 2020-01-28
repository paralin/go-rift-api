// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LolClashRosterStats lol clash roster stats
// swagger:model LolClashRosterStats
type LolClashRosterStats struct {

	// end time ms
	EndTimeMs int64 `json:"endTimeMs,omitempty"`

	// period stats
	PeriodStats []*LolClashRosterPeriodAggregatedStats `json:"periodStats"`

	// player stats
	PlayerStats map[string]LolClashRosterPlayerAggregatedStats `json:"playerStats,omitempty"`

	// roster icon color Id
	RosterIconColorID int32 `json:"rosterIconColorId,omitempty"`

	// roster icon Id
	RosterIconID int32 `json:"rosterIconId,omitempty"`

	// roster Id
	RosterID int64 `json:"rosterId,omitempty"`

	// roster name
	RosterName string `json:"rosterName,omitempty"`

	// roster short name
	RosterShortName string `json:"rosterShortName,omitempty"`

	// start time ms
	StartTimeMs int64 `json:"startTimeMs,omitempty"`

	// tier
	Tier int32 `json:"tier,omitempty"`

	// tournament name loc key
	TournamentNameLocKey string `json:"tournamentNameLocKey,omitempty"`

	// tournament name loc key secondary
	TournamentNameLocKeySecondary string `json:"tournamentNameLocKeySecondary,omitempty"`

	// tournament periods
	TournamentPeriods int32 `json:"tournamentPeriods,omitempty"`

	// tournament theme Id
	TournamentThemeID int32 `json:"tournamentThemeId,omitempty"`
}

// Validate validates this lol clash roster stats
func (m *LolClashRosterStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeriodStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayerStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashRosterStats) validatePeriodStats(formats strfmt.Registry) error {

	if swag.IsZero(m.PeriodStats) { // not required
		return nil
	}

	for i := 0; i < len(m.PeriodStats); i++ {
		if swag.IsZero(m.PeriodStats[i]) { // not required
			continue
		}

		if m.PeriodStats[i] != nil {
			if err := m.PeriodStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periodStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClashRosterStats) validatePlayerStats(formats strfmt.Registry) error {

	if swag.IsZero(m.PlayerStats) { // not required
		return nil
	}

	for k := range m.PlayerStats {

		if err := validate.Required("playerStats"+"."+k, "body", m.PlayerStats[k]); err != nil {
			return err
		}
		if val, ok := m.PlayerStats[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashRosterStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashRosterStats) UnmarshalBinary(b []byte) error {
	var res LolClashRosterStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}