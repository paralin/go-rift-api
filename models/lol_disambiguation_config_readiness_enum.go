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

// LolDisambiguationConfigReadinessEnum lol disambiguation config readiness enum
// swagger:model LolDisambiguationConfigReadinessEnum
type LolDisambiguationConfigReadinessEnum string

const (

	// LolDisambiguationConfigReadinessEnumNotReady captures enum value "NotReady"
	LolDisambiguationConfigReadinessEnumNotReady LolDisambiguationConfigReadinessEnum = "NotReady"

	// LolDisambiguationConfigReadinessEnumReady captures enum value "Ready"
	LolDisambiguationConfigReadinessEnumReady LolDisambiguationConfigReadinessEnum = "Ready"
)

// for schema
var lolDisambiguationConfigReadinessEnumEnum []interface{}

func init() {
	var res []LolDisambiguationConfigReadinessEnum
	if err := json.Unmarshal([]byte(`["NotReady","Ready"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolDisambiguationConfigReadinessEnumEnum = append(lolDisambiguationConfigReadinessEnumEnum, v)
	}
}

func (m LolDisambiguationConfigReadinessEnum) validateLolDisambiguationConfigReadinessEnumEnum(path, location string, value LolDisambiguationConfigReadinessEnum) error {
	if err := validate.Enum(path, location, value, lolDisambiguationConfigReadinessEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol disambiguation config readiness enum
func (m LolDisambiguationConfigReadinessEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolDisambiguationConfigReadinessEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
