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

// TournamentStatusEnum tournament status enum
// swagger:model TournamentStatusEnum
type TournamentStatusEnum string

const (

	// TournamentStatusEnumDEFAULT captures enum value "DEFAULT"
	TournamentStatusEnumDEFAULT TournamentStatusEnum = "DEFAULT"

	// TournamentStatusEnumCANCELLED captures enum value "CANCELLED"
	TournamentStatusEnumCANCELLED TournamentStatusEnum = "CANCELLED"

	// TournamentStatusEnumPAUSED captures enum value "PAUSED"
	TournamentStatusEnumPAUSED TournamentStatusEnum = "PAUSED"

	// TournamentStatusEnumPRERESUME captures enum value "PRERESUME"
	TournamentStatusEnumPRERESUME TournamentStatusEnum = "PRERESUME"
)

// for schema
var tournamentStatusEnumEnum []interface{}

func init() {
	var res []TournamentStatusEnum
	if err := json.Unmarshal([]byte(`["DEFAULT","CANCELLED","PAUSED","PRERESUME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tournamentStatusEnumEnum = append(tournamentStatusEnumEnum, v)
	}
}

func (m TournamentStatusEnum) validateTournamentStatusEnumEnum(path, location string, value TournamentStatusEnum) error {
	if err := validate.Enum(path, location, value, tournamentStatusEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tournament status enum
func (m TournamentStatusEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTournamentStatusEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
