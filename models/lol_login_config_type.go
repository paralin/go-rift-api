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

// LolLoginConfigType lol login config type
// swagger:model LolLoginConfigType
type LolLoginConfigType string

const (

	// LolLoginConfigTypePublic captures enum value "public"
	LolLoginConfigTypePublic LolLoginConfigType = "public"

	// LolLoginConfigTypePlayer captures enum value "player"
	LolLoginConfigTypePlayer LolLoginConfigType = "player"
)

// for schema
var lolLoginConfigTypeEnum []interface{}

func init() {
	var res []LolLoginConfigType
	if err := json.Unmarshal([]byte(`["public","player"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLoginConfigTypeEnum = append(lolLoginConfigTypeEnum, v)
	}
}

func (m LolLoginConfigType) validateLolLoginConfigTypeEnum(path, location string, value LolLoginConfigType) error {
	if err := validate.Enum(path, location, value, lolLoginConfigTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol login config type
func (m LolLoginConfigType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLoginConfigTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
