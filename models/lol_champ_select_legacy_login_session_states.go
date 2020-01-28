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

// LolChampSelectLegacyLoginSessionStates lol champ select legacy login session states
// swagger:model LolChampSelectLegacyLoginSessionStates
type LolChampSelectLegacyLoginSessionStates string

const (

	// LolChampSelectLegacyLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolChampSelectLegacyLoginSessionStatesINPROGRESS LolChampSelectLegacyLoginSessionStates = "IN_PROGRESS"

	// LolChampSelectLegacyLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolChampSelectLegacyLoginSessionStatesSUCCEEDED LolChampSelectLegacyLoginSessionStates = "SUCCEEDED"

	// LolChampSelectLegacyLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolChampSelectLegacyLoginSessionStatesLOGGINGOUT LolChampSelectLegacyLoginSessionStates = "LOGGING_OUT"

	// LolChampSelectLegacyLoginSessionStatesERROR captures enum value "ERROR"
	LolChampSelectLegacyLoginSessionStatesERROR LolChampSelectLegacyLoginSessionStates = "ERROR"
)

// for schema
var lolChampSelectLegacyLoginSessionStatesEnum []interface{}

func init() {
	var res []LolChampSelectLegacyLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolChampSelectLegacyLoginSessionStatesEnum = append(lolChampSelectLegacyLoginSessionStatesEnum, v)
	}
}

func (m LolChampSelectLegacyLoginSessionStates) validateLolChampSelectLegacyLoginSessionStatesEnum(path, location string, value LolChampSelectLegacyLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolChampSelectLegacyLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol champ select legacy login session states
func (m LolChampSelectLegacyLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolChampSelectLegacyLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
