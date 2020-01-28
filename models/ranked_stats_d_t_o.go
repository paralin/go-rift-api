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

// RankedStatsDTO ranked stats d t o
// swagger:model RankedStatsDTO
type RankedStatsDTO struct {

	// earned regalia reward ids
	EarnedRegaliaRewardIds []string `json:"earnedRegaliaRewardIds"`

	// highest previous season achieved rank
	HighestPreviousSeasonAchievedRank string `json:"highestPreviousSeasonAchievedRank,omitempty"`

	// highest previous season achieved tier
	HighestPreviousSeasonAchievedTier string `json:"highestPreviousSeasonAchievedTier,omitempty"`

	// highest previous season end rank
	HighestPreviousSeasonEndRank string `json:"highestPreviousSeasonEndRank,omitempty"`

	// highest previous season end tier
	HighestPreviousSeasonEndTier string `json:"highestPreviousSeasonEndTier,omitempty"`

	// queues
	Queues []*RankedQueueStatsDTO `json:"queues"`

	// splits progress
	SplitsProgress map[string]int32 `json:"splitsProgress,omitempty"`
}

// Validate validates this ranked stats d t o
func (m *RankedStatsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RankedStatsDTO) validateQueues(formats strfmt.Registry) error {

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
func (m *RankedStatsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RankedStatsDTO) UnmarshalBinary(b []byte) error {
	var res RankedStatsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}