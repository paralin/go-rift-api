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

// LolLoginLoginSessionStates lol login login session states
// swagger:model LolLoginLoginSessionStates
type LolLoginLoginSessionStates string

const (

	// LolLoginLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolLoginLoginSessionStatesINPROGRESS LolLoginLoginSessionStates = "IN_PROGRESS"

	// LolLoginLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolLoginLoginSessionStatesSUCCEEDED LolLoginLoginSessionStates = "SUCCEEDED"

	// LolLoginLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolLoginLoginSessionStatesLOGGINGOUT LolLoginLoginSessionStates = "LOGGING_OUT"

	// LolLoginLoginSessionStatesERROR captures enum value "ERROR"
	LolLoginLoginSessionStatesERROR LolLoginLoginSessionStates = "ERROR"
)

// for schema
var lolLoginLoginSessionStatesEnum []interface{}

func init() {
	var res []LolLoginLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLoginLoginSessionStatesEnum = append(lolLoginLoginSessionStatesEnum, v)
	}
}

func (m LolLoginLoginSessionStates) validateLolLoginLoginSessionStatesEnum(path, location string, value LolLoginLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolLoginLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol login login session states
func (m LolLoginLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLoginLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
