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

// LogSeverityLevels Allowable severity levels for log events.
// swagger:model LogSeverityLevels
type LogSeverityLevels string

const (

	// LogSeverityLevelsOkay captures enum value "Okay"
	LogSeverityLevelsOkay LogSeverityLevels = "Okay"

	// LogSeverityLevelsWarning captures enum value "Warning"
	LogSeverityLevelsWarning LogSeverityLevels = "Warning"

	// LogSeverityLevelsError captures enum value "Error"
	LogSeverityLevelsError LogSeverityLevels = "Error"

	// LogSeverityLevelsAlways captures enum value "Always"
	LogSeverityLevelsAlways LogSeverityLevels = "Always"
)

// for schema
var logSeverityLevelsEnum []interface{}

func init() {
	var res []LogSeverityLevels
	if err := json.Unmarshal([]byte(`["Okay","Warning","Error","Always"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logSeverityLevelsEnum = append(logSeverityLevelsEnum, v)
	}
}

func (m LogSeverityLevels) validateLogSeverityLevelsEnum(path, location string, value LogSeverityLevels) error {
	if err := validate.Enum(path, location, value, logSeverityLevelsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this log severity levels
func (m LogSeverityLevels) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLogSeverityLevelsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
