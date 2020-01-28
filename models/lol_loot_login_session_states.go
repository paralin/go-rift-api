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

// LolLootLoginSessionStates lol loot login session states
// swagger:model LolLootLoginSessionStates
type LolLootLoginSessionStates string

const (

	// LolLootLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolLootLoginSessionStatesINPROGRESS LolLootLoginSessionStates = "IN_PROGRESS"

	// LolLootLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolLootLoginSessionStatesSUCCEEDED LolLootLoginSessionStates = "SUCCEEDED"

	// LolLootLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolLootLoginSessionStatesLOGGINGOUT LolLootLoginSessionStates = "LOGGING_OUT"

	// LolLootLoginSessionStatesERROR captures enum value "ERROR"
	LolLootLoginSessionStatesERROR LolLootLoginSessionStates = "ERROR"
)

// for schema
var lolLootLoginSessionStatesEnum []interface{}

func init() {
	var res []LolLootLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLootLoginSessionStatesEnum = append(lolLootLoginSessionStatesEnum, v)
	}
}

func (m LolLootLoginSessionStates) validateLolLootLoginSessionStatesEnum(path, location string, value LolLootLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolLootLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol loot login session states
func (m LolLootLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLootLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}