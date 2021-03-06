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

// LolStoreLoginSessionStates lol store login session states
// swagger:model LolStoreLoginSessionStates
type LolStoreLoginSessionStates string

const (

	// LolStoreLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolStoreLoginSessionStatesINPROGRESS LolStoreLoginSessionStates = "IN_PROGRESS"

	// LolStoreLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolStoreLoginSessionStatesSUCCEEDED LolStoreLoginSessionStates = "SUCCEEDED"

	// LolStoreLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolStoreLoginSessionStatesLOGGINGOUT LolStoreLoginSessionStates = "LOGGING_OUT"

	// LolStoreLoginSessionStatesERROR captures enum value "ERROR"
	LolStoreLoginSessionStatesERROR LolStoreLoginSessionStates = "ERROR"
)

// for schema
var lolStoreLoginSessionStatesEnum []interface{}

func init() {
	var res []LolStoreLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolStoreLoginSessionStatesEnum = append(lolStoreLoginSessionStatesEnum, v)
	}
}

func (m LolStoreLoginSessionStates) validateLolStoreLoginSessionStatesEnum(path, location string, value LolStoreLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolStoreLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol store login session states
func (m LolStoreLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolStoreLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
