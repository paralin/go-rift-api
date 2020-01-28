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

// LolCareerStatsSummonersRiftPosition lol career stats summoners rift position
// swagger:model LolCareerStatsSummonersRiftPosition
type LolCareerStatsSummonersRiftPosition string

const (

	// LolCareerStatsSummonersRiftPositionALL captures enum value "ALL"
	LolCareerStatsSummonersRiftPositionALL LolCareerStatsSummonersRiftPosition = "ALL"

	// LolCareerStatsSummonersRiftPositionUNKNOWN captures enum value "UNKNOWN"
	LolCareerStatsSummonersRiftPositionUNKNOWN LolCareerStatsSummonersRiftPosition = "UNKNOWN"

	// LolCareerStatsSummonersRiftPositionTOP captures enum value "TOP"
	LolCareerStatsSummonersRiftPositionTOP LolCareerStatsSummonersRiftPosition = "TOP"

	// LolCareerStatsSummonersRiftPositionJUNGLE captures enum value "JUNGLE"
	LolCareerStatsSummonersRiftPositionJUNGLE LolCareerStatsSummonersRiftPosition = "JUNGLE"

	// LolCareerStatsSummonersRiftPositionMID captures enum value "MID"
	LolCareerStatsSummonersRiftPositionMID LolCareerStatsSummonersRiftPosition = "MID"

	// LolCareerStatsSummonersRiftPositionBOTTOM captures enum value "BOTTOM"
	LolCareerStatsSummonersRiftPositionBOTTOM LolCareerStatsSummonersRiftPosition = "BOTTOM"

	// LolCareerStatsSummonersRiftPositionSUPPORT captures enum value "SUPPORT"
	LolCareerStatsSummonersRiftPositionSUPPORT LolCareerStatsSummonersRiftPosition = "SUPPORT"
)

// for schema
var lolCareerStatsSummonersRiftPositionEnum []interface{}

func init() {
	var res []LolCareerStatsSummonersRiftPosition
	if err := json.Unmarshal([]byte(`["ALL","UNKNOWN","TOP","JUNGLE","MID","BOTTOM","SUPPORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolCareerStatsSummonersRiftPositionEnum = append(lolCareerStatsSummonersRiftPositionEnum, v)
	}
}

func (m LolCareerStatsSummonersRiftPosition) validateLolCareerStatsSummonersRiftPositionEnum(path, location string, value LolCareerStatsSummonersRiftPosition) error {
	if err := validate.Enum(path, location, value, lolCareerStatsSummonersRiftPositionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol career stats summoners rift position
func (m LolCareerStatsSummonersRiftPosition) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolCareerStatsSummonersRiftPositionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}