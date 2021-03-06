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

// LolCareerStatsPositionStatistics lol career stats position statistics
// swagger:model LolCareerStatsPositionStatistics
type LolCareerStatsPositionStatistics struct {

	// experts
	Experts []*LolCareerStatsExpertPlayer `json:"experts"`

	// position
	Position LolCareerStatsSummonersRiftPosition `json:"position,omitempty"`

	// queue stats
	QueueStats []*LolCareerStatsStatisticsByQueue `json:"queueStats"`
}

// Validate validates this lol career stats position statistics
func (m *LolCareerStatsPositionStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExperts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCareerStatsPositionStatistics) validateExperts(formats strfmt.Registry) error {

	if swag.IsZero(m.Experts) { // not required
		return nil
	}

	for i := 0; i < len(m.Experts); i++ {
		if swag.IsZero(m.Experts[i]) { // not required
			continue
		}

		if m.Experts[i] != nil {
			if err := m.Experts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("experts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolCareerStatsPositionStatistics) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := m.Position.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("position")
		}
		return err
	}

	return nil
}

func (m *LolCareerStatsPositionStatistics) validateQueueStats(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueStats) { // not required
		return nil
	}

	for i := 0; i < len(m.QueueStats); i++ {
		if swag.IsZero(m.QueueStats[i]) { // not required
			continue
		}

		if m.QueueStats[i] != nil {
			if err := m.QueueStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queueStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCareerStatsPositionStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCareerStatsPositionStatistics) UnmarshalBinary(b []byte) error {
	var res LolCareerStatsPositionStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
