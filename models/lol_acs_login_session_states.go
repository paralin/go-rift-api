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

// LolAcsLoginSessionStates lol acs login session states
// swagger:model LolAcsLoginSessionStates
type LolAcsLoginSessionStates string

const (

	// LolAcsLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolAcsLoginSessionStatesINPROGRESS LolAcsLoginSessionStates = "IN_PROGRESS"

	// LolAcsLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolAcsLoginSessionStatesSUCCEEDED LolAcsLoginSessionStates = "SUCCEEDED"

	// LolAcsLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolAcsLoginSessionStatesLOGGINGOUT LolAcsLoginSessionStates = "LOGGING_OUT"

	// LolAcsLoginSessionStatesERROR captures enum value "ERROR"
	LolAcsLoginSessionStatesERROR LolAcsLoginSessionStates = "ERROR"
)

// for schema
var lolAcsLoginSessionStatesEnum []interface{}

func init() {
	var res []LolAcsLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolAcsLoginSessionStatesEnum = append(lolAcsLoginSessionStatesEnum, v)
	}
}

func (m LolAcsLoginSessionStates) validateLolAcsLoginSessionStatesEnum(path, location string, value LolAcsLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolAcsLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol acs login session states
func (m LolAcsLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolAcsLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
