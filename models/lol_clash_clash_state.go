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

// LolClashClashState lol clash clash state
// swagger:model LolClashClashState
type LolClashClashState string

const (

	// LolClashClashStateDarkDisabled captures enum value "DarkDisabled"
	LolClashClashStateDarkDisabled LolClashClashState = "DarkDisabled"

	// LolClashClashStateDarkEnabled captures enum value "DarkEnabled"
	LolClashClashStateDarkEnabled LolClashClashState = "DarkEnabled"

	// LolClashClashStateDisabled captures enum value "Disabled"
	LolClashClashStateDisabled LolClashClashState = "Disabled"

	// LolClashClashStateEnabled captures enum value "Enabled"
	LolClashClashStateEnabled LolClashClashState = "Enabled"
)

// for schema
var lolClashClashStateEnum []interface{}

func init() {
	var res []LolClashClashState
	if err := json.Unmarshal([]byte(`["DarkDisabled","DarkEnabled","Disabled","Enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolClashClashStateEnum = append(lolClashClashStateEnum, v)
	}
}

func (m LolClashClashState) validateLolClashClashStateEnum(path, location string, value LolClashClashState) error {
	if err := validate.Enum(path, location, value, lolClashClashStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol clash clash state
func (m LolClashClashState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolClashClashStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
