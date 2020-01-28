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

// ClientConfigConfigReadinessEnum client config config readiness enum
// swagger:model ClientConfigConfigReadinessEnum
type ClientConfigConfigReadinessEnum string

const (

	// ClientConfigConfigReadinessEnumNotReady captures enum value "NotReady"
	ClientConfigConfigReadinessEnumNotReady ClientConfigConfigReadinessEnum = "NotReady"

	// ClientConfigConfigReadinessEnumReady captures enum value "Ready"
	ClientConfigConfigReadinessEnumReady ClientConfigConfigReadinessEnum = "Ready"

	// ClientConfigConfigReadinessEnumDisabled captures enum value "Disabled"
	ClientConfigConfigReadinessEnumDisabled ClientConfigConfigReadinessEnum = "Disabled"
)

// for schema
var clientConfigConfigReadinessEnumEnum []interface{}

func init() {
	var res []ClientConfigConfigReadinessEnum
	if err := json.Unmarshal([]byte(`["NotReady","Ready","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientConfigConfigReadinessEnumEnum = append(clientConfigConfigReadinessEnumEnum, v)
	}
}

func (m ClientConfigConfigReadinessEnum) validateClientConfigConfigReadinessEnumEnum(path, location string, value ClientConfigConfigReadinessEnum) error {
	if err := validate.Enum(path, location, value, clientConfigConfigReadinessEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this client config config readiness enum
func (m ClientConfigConfigReadinessEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClientConfigConfigReadinessEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}