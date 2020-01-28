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

// CapacityEnum capacity enum
// swagger:model CapacityEnum
type CapacityEnum string

const (

	// CapacityEnumLOW captures enum value "LOW"
	CapacityEnumLOW CapacityEnum = "LOW"

	// CapacityEnumMEDIUM captures enum value "MEDIUM"
	CapacityEnumMEDIUM CapacityEnum = "MEDIUM"

	// CapacityEnumHIGH captures enum value "HIGH"
	CapacityEnumHIGH CapacityEnum = "HIGH"

	// CapacityEnumFULL captures enum value "FULL"
	CapacityEnumFULL CapacityEnum = "FULL"
)

// for schema
var capacityEnumEnum []interface{}

func init() {
	var res []CapacityEnum
	if err := json.Unmarshal([]byte(`["LOW","MEDIUM","HIGH","FULL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		capacityEnumEnum = append(capacityEnumEnum, v)
	}
}

func (m CapacityEnum) validateCapacityEnumEnum(path, location string, value CapacityEnum) error {
	if err := validate.Enum(path, location, value, capacityEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this capacity enum
func (m CapacityEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCapacityEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
