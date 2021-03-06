// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MetricMetadataAlert metric metadata alert
// swagger:model MetricMetadataAlert
type MetricMetadataAlert struct {

	// description
	Description string `json:"description,omitempty"`

	// info
	Info string `json:"info,omitempty"`

	// level
	Level string `json:"level,omitempty"`

	// max
	Max float64 `json:"max,omitempty"`

	// min
	Min float64 `json:"min,omitempty"`

	// notify
	Notify *MetricMetadataNotify `json:"notify,omitempty"`

	// pretty name
	PrettyName string `json:"pretty_name,omitempty"`
}

// Validate validates this metric metadata alert
func (m *MetricMetadataAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotify(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricMetadataAlert) validateNotify(formats strfmt.Registry) error {

	if swag.IsZero(m.Notify) { // not required
		return nil
	}

	if m.Notify != nil {
		if err := m.Notify.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notify")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricMetadataAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricMetadataAlert) UnmarshalBinary(b []byte) error {
	var res MetricMetadataAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
