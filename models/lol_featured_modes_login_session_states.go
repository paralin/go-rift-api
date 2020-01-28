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

// LolFeaturedModesLoginSessionStates lol featured modes login session states
// swagger:model LolFeaturedModesLoginSessionStates
type LolFeaturedModesLoginSessionStates string

const (

	// LolFeaturedModesLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolFeaturedModesLoginSessionStatesINPROGRESS LolFeaturedModesLoginSessionStates = "IN_PROGRESS"

	// LolFeaturedModesLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolFeaturedModesLoginSessionStatesSUCCEEDED LolFeaturedModesLoginSessionStates = "SUCCEEDED"

	// LolFeaturedModesLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolFeaturedModesLoginSessionStatesLOGGINGOUT LolFeaturedModesLoginSessionStates = "LOGGING_OUT"

	// LolFeaturedModesLoginSessionStatesERROR captures enum value "ERROR"
	LolFeaturedModesLoginSessionStatesERROR LolFeaturedModesLoginSessionStates = "ERROR"
)

// for schema
var lolFeaturedModesLoginSessionStatesEnum []interface{}

func init() {
	var res []LolFeaturedModesLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolFeaturedModesLoginSessionStatesEnum = append(lolFeaturedModesLoginSessionStatesEnum, v)
	}
}

func (m LolFeaturedModesLoginSessionStates) validateLolFeaturedModesLoginSessionStatesEnum(path, location string, value LolFeaturedModesLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolFeaturedModesLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol featured modes login session states
func (m LolFeaturedModesLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolFeaturedModesLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
