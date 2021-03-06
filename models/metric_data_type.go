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

// MetricDataType metric data type
// swagger:model MetricDataType
type MetricDataType string

const (

	// MetricDataTypeUnknown captures enum value "unknown"
	MetricDataTypeUnknown MetricDataType = "unknown"

	// MetricDataTypeInt captures enum value "int"
	MetricDataTypeInt MetricDataType = "int"

	// MetricDataTypeUint captures enum value "uint"
	MetricDataTypeUint MetricDataType = "uint"

	// MetricDataTypeFloat captures enum value "float"
	MetricDataTypeFloat MetricDataType = "float"

	// MetricDataTypeBool captures enum value "bool"
	MetricDataTypeBool MetricDataType = "bool"

	// MetricDataTypeString captures enum value "string"
	MetricDataTypeString MetricDataType = "string"
)

// for schema
var metricDataTypeEnum []interface{}

func init() {
	var res []MetricDataType
	if err := json.Unmarshal([]byte(`["unknown","int","uint","float","bool","string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricDataTypeEnum = append(metricDataTypeEnum, v)
	}
}

func (m MetricDataType) validateMetricDataTypeEnum(path, location string, value MetricDataType) error {
	if err := validate.Enum(path, location, value, metricDataTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this metric data type
func (m MetricDataType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMetricDataTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
