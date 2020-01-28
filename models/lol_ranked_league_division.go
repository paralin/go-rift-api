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

// LolRankedLeagueDivision lol ranked league division
// swagger:model LolRankedLeagueDivision
type LolRankedLeagueDivision string

const (

	// LolRankedLeagueDivisionI captures enum value "I"
	LolRankedLeagueDivisionI LolRankedLeagueDivision = "I"

	// LolRankedLeagueDivisionII captures enum value "II"
	LolRankedLeagueDivisionII LolRankedLeagueDivision = "II"

	// LolRankedLeagueDivisionIII captures enum value "III"
	LolRankedLeagueDivisionIII LolRankedLeagueDivision = "III"

	// LolRankedLeagueDivisionIV captures enum value "IV"
	LolRankedLeagueDivisionIV LolRankedLeagueDivision = "IV"

	// LolRankedLeagueDivisionV captures enum value "V"
	LolRankedLeagueDivisionV LolRankedLeagueDivision = "V"

	// LolRankedLeagueDivisionNA captures enum value "NA"
	LolRankedLeagueDivisionNA LolRankedLeagueDivision = "NA"
)

// for schema
var lolRankedLeagueDivisionEnum []interface{}

func init() {
	var res []LolRankedLeagueDivision
	if err := json.Unmarshal([]byte(`["I","II","III","IV","V","NA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolRankedLeagueDivisionEnum = append(lolRankedLeagueDivisionEnum, v)
	}
}

func (m LolRankedLeagueDivision) validateLolRankedLeagueDivisionEnum(path, location string, value LolRankedLeagueDivision) error {
	if err := validate.Enum(path, location, value, lolRankedLeagueDivisionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol ranked league division
func (m LolRankedLeagueDivision) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolRankedLeagueDivisionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
