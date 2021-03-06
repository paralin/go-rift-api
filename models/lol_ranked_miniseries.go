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

// LolRankedMiniseries lol ranked miniseries
// swagger:model LolRankedMiniseries
type LolRankedMiniseries string

const (

	// LolRankedMiniseriesW captures enum value "W"
	LolRankedMiniseriesW LolRankedMiniseries = "W"

	// LolRankedMiniseriesL captures enum value "L"
	LolRankedMiniseriesL LolRankedMiniseries = "L"

	// LolRankedMiniseriesN captures enum value "N"
	LolRankedMiniseriesN LolRankedMiniseries = "N"
)

// for schema
var lolRankedMiniseriesEnum []interface{}

func init() {
	var res []LolRankedMiniseries
	if err := json.Unmarshal([]byte(`["W","L","N"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolRankedMiniseriesEnum = append(lolRankedMiniseriesEnum, v)
	}
}

func (m LolRankedMiniseries) validateLolRankedMiniseriesEnum(path, location string, value LolRankedMiniseries) error {
	if err := validate.Enum(path, location, value, lolRankedMiniseriesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol ranked miniseries
func (m LolRankedMiniseries) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolRankedMiniseriesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
