// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TimeSeriesEventMarkerV1 time series event marker v1
// swagger:model TimeSeriesEventMarkerV1
type TimeSeriesEventMarkerV1 struct {

	// event name
	EventName string `json:"eventName,omitempty"`

	// event marker name
	MarkerName string `json:"markerName,omitempty"`

	// timestamp in microseconds of when the event occurred
	When int64 `json:"when,omitempty"`
}

// Validate validates this time series event marker v1
func (m *TimeSeriesEventMarkerV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeSeriesEventMarkerV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeSeriesEventMarkerV1) UnmarshalBinary(b []byte) error {
	var res TimeSeriesEventMarkerV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
