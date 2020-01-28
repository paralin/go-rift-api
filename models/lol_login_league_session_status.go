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

// LolLoginLeagueSessionStatus lol login league session status
// swagger:model LolLoginLeagueSessionStatus
type LolLoginLeagueSessionStatus string

const (

	// LolLoginLeagueSessionStatusUNINITIALIZED captures enum value "UNINITIALIZED"
	LolLoginLeagueSessionStatusUNINITIALIZED LolLoginLeagueSessionStatus = "UNINITIALIZED"

	// LolLoginLeagueSessionStatusINITIALIZED captures enum value "INITIALIZED"
	LolLoginLeagueSessionStatusINITIALIZED LolLoginLeagueSessionStatus = "INITIALIZED"

	// LolLoginLeagueSessionStatusEXPIRED captures enum value "EXPIRED"
	LolLoginLeagueSessionStatusEXPIRED LolLoginLeagueSessionStatus = "EXPIRED"

	// LolLoginLeagueSessionStatusDUPLICATED captures enum value "DUPLICATED"
	LolLoginLeagueSessionStatusDUPLICATED LolLoginLeagueSessionStatus = "DUPLICATED"
)

// for schema
var lolLoginLeagueSessionStatusEnum []interface{}

func init() {
	var res []LolLoginLeagueSessionStatus
	if err := json.Unmarshal([]byte(`["UNINITIALIZED","INITIALIZED","EXPIRED","DUPLICATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLoginLeagueSessionStatusEnum = append(lolLoginLeagueSessionStatusEnum, v)
	}
}

func (m LolLoginLeagueSessionStatus) validateLolLoginLeagueSessionStatusEnum(path, location string, value LolLoginLeagueSessionStatus) error {
	if err := validate.Enum(path, location, value, lolLoginLeagueSessionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol login league session status
func (m LolLoginLeagueSessionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLoginLeagueSessionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}