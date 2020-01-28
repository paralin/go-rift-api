// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolRegaliaRankedStats lol regalia ranked stats
// swagger:model LolRegaliaRankedStats
type LolRegaliaRankedStats struct {

	// highest previous season achieved tier
	HighestPreviousSeasonAchievedTier LolRegaliaLeagueTier `json:"highestPreviousSeasonAchievedTier,omitempty"`

	// highest ranked entry
	HighestRankedEntry *LolRegaliaRankedQueueStats `json:"highestRankedEntry,omitempty"`

	// queues
	Queues []*LolRegaliaRankedQueueStats `json:"queues"`

	// ranked regalia level
	RankedRegaliaLevel int32 `json:"rankedRegaliaLevel,omitempty"`
}

// Validate validates this lol regalia ranked stats
func (m *LolRegaliaRankedStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHighestPreviousSeasonAchievedTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHighestRankedEntry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolRegaliaRankedStats) validateHighestPreviousSeasonAchievedTier(formats strfmt.Registry) error {

	if swag.IsZero(m.HighestPreviousSeasonAchievedTier) { // not required
		return nil
	}

	if err := m.HighestPreviousSeasonAchievedTier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("highestPreviousSeasonAchievedTier")
		}
		return err
	}

	return nil
}

func (m *LolRegaliaRankedStats) validateHighestRankedEntry(formats strfmt.Registry) error {

	if swag.IsZero(m.HighestRankedEntry) { // not required
		return nil
	}

	if m.HighestRankedEntry != nil {
		if err := m.HighestRankedEntry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("highestRankedEntry")
			}
			return err
		}
	}

	return nil
}

func (m *LolRegaliaRankedStats) validateQueues(formats strfmt.Registry) error {

	if swag.IsZero(m.Queues) { // not required
		return nil
	}

	for i := 0; i < len(m.Queues); i++ {
		if swag.IsZero(m.Queues[i]) { // not required
			continue
		}

		if m.Queues[i] != nil {
			if err := m.Queues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolRegaliaRankedStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRegaliaRankedStats) UnmarshalBinary(b []byte) error {
	var res LolRegaliaRankedStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}