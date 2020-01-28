// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolMatchHistoryMatchHistoryParticipant lol match history match history participant
// swagger:model LolMatchHistoryMatchHistoryParticipant
type LolMatchHistoryMatchHistoryParticipant struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// highest achieved season tier
	HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier,omitempty"`

	// participant Id
	ParticipantID int64 `json:"participantId,omitempty"`

	// spell1 Id
	Spell1ID int64 `json:"spell1Id,omitempty"`

	// spell2 Id
	Spell2ID int64 `json:"spell2Id,omitempty"`

	// stats
	Stats *LolMatchHistoryMatchHistoryParticipantStatistics `json:"stats,omitempty"`

	// team Id
	TeamID int64 `json:"teamId,omitempty"`

	// timeline
	Timeline *LolMatchHistoryMatchHistoryTimeline `json:"timeline,omitempty"`
}

// Validate validates this lol match history match history participant
func (m *LolMatchHistoryMatchHistoryParticipant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolMatchHistoryMatchHistoryParticipant) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *LolMatchHistoryMatchHistoryParticipant) validateTimeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeline) { // not required
		return nil
	}

	if m.Timeline != nil {
		if err := m.Timeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeline")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolMatchHistoryMatchHistoryParticipant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMatchHistoryMatchHistoryParticipant) UnmarshalBinary(b []byte) error {
	var res LolMatchHistoryMatchHistoryParticipant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}