// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MetricType metric type
// swagger:model MetricType
type MetricType string

const (

	// MetricTypeUnknown captures enum value "unknown"
	MetricTypeUnknown MetricType = "unknown"

	// MetricTypePercentage captures enum value "percentage"
	MetricTypePercentage MetricType = "percentage"

	// MetricTypeCount captures enum value "count"
	MetricTypeCount MetricType = "count"

	// MetricTypeDuration captures enum value "duration"
	MetricTypeDuration MetricType = "duration"

	// MetricTypeRate captures enum value "rate"
	MetricTypeRate MetricType = "rate"
)

// for schema
var metricTypeEnum []interface{}

func init() {
	var res []MetricType
	if err := json.Unmarshal([]byte(`["unknown","percentage","count","duration","rate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricTypeEnum = append(metricTypeEnum, v)
	}
}

func (m MetricType) validateMetricTypeEnum(path, location string, value MetricType) error {
	if err := validate.Enum(path, location, value, metricTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this metric type
func (m MetricType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMetricTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
