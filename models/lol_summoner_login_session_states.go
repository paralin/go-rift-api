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

// LolSummonerLoginSessionStates lol summoner login session states
// swagger:model LolSummonerLoginSessionStates
type LolSummonerLoginSessionStates string

const (

	// LolSummonerLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolSummonerLoginSessionStatesINPROGRESS LolSummonerLoginSessionStates = "IN_PROGRESS"

	// LolSummonerLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolSummonerLoginSessionStatesSUCCEEDED LolSummonerLoginSessionStates = "SUCCEEDED"

	// LolSummonerLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolSummonerLoginSessionStatesLOGGINGOUT LolSummonerLoginSessionStates = "LOGGING_OUT"

	// LolSummonerLoginSessionStatesERROR captures enum value "ERROR"
	LolSummonerLoginSessionStatesERROR LolSummonerLoginSessionStates = "ERROR"
)

// for schema
var lolSummonerLoginSessionStatesEnum []interface{}

func init() {
	var res []LolSummonerLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolSummonerLoginSessionStatesEnum = append(lolSummonerLoginSessionStatesEnum, v)
	}
}

func (m LolSummonerLoginSessionStates) validateLolSummonerLoginSessionStatesEnum(path, location string, value LolSummonerLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolSummonerLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol summoner login session states
func (m LolSummonerLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolSummonerLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
