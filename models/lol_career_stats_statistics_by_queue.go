// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolCareerStatsStatisticsByQueue lol career stats statistics by queue
// swagger:model LolCareerStatsStatisticsByQueue
type LolCareerStatsStatisticsByQueue struct {

	// queue type
	QueueType LolCareerStatsCareerStatsQueueType `json:"queueType,omitempty"`

	// stats
	Stats interface{} `json:"stats,omitempty"`
}

// Validate validates this lol career stats statistics by queue
func (m *LolCareerStatsStatisticsByQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCareerStatsStatisticsByQueue) validateQueueType(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolCareerStatsStatisticsByQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCareerStatsStatisticsByQueue) UnmarshalBinary(b []byte) error {
	var res LolCareerStatsStatisticsByQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
