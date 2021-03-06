// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolCareerStatsPositionStatsQueryRequest lol career stats position stats query request
// swagger:model LolCareerStatsPositionStatsQueryRequest
type LolCareerStatsPositionStatsQueryRequest struct {

	// position
	Position LolCareerStatsSummonersRiftPosition `json:"position,omitempty"`

	// queue type
	QueueType LolCareerStatsCareerStatsQueueType `json:"queueType,omitempty"`

	// rank tier
	RankTier LolCareerStatsRankedTier `json:"rankTier,omitempty"`

	// season
	Season int32 `json:"season,omitempty"`
}

// Validate validates this lol career stats position stats query request
func (m *LolCareerStatsPositionStatsQueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRankTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCareerStatsPositionStatsQueryRequest) validatePosition(formats strfmt.Registry) error {

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

func (m *LolCareerStatsPositionStatsQueryRequest) validateQueueType(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueType) { // not required
		return nil
	}

	if err := m.QueueType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("queueType")
		}
		return err
	}

	return nil
}

func (m *LolCareerStatsPositionStatsQueryRequest) validateRankTier(formats strfmt.Registry) error {

	if swag.IsZero(m.RankTier) { // not required
		return nil
	}

	if err := m.RankTier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rankTier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCareerStatsPositionStatsQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCareerStatsPositionStatsQueryRequest) UnmarshalBinary(b []byte) error {
	var res LolCareerStatsPositionStatsQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
