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

// LolQueueEligibilityLoginSessionStates lol queue eligibility login session states
// swagger:model LolQueueEligibilityLoginSessionStates
type LolQueueEligibilityLoginSessionStates string

const (

	// LolQueueEligibilityLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolQueueEligibilityLoginSessionStatesINPROGRESS LolQueueEligibilityLoginSessionStates = "IN_PROGRESS"

	// LolQueueEligibilityLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolQueueEligibilityLoginSessionStatesSUCCEEDED LolQueueEligibilityLoginSessionStates = "SUCCEEDED"

	// LolQueueEligibilityLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolQueueEligibilityLoginSessionStatesLOGGINGOUT LolQueueEligibilityLoginSessionStates = "LOGGING_OUT"

	// LolQueueEligibilityLoginSessionStatesERROR captures enum value "ERROR"
	LolQueueEligibilityLoginSessionStatesERROR LolQueueEligibilityLoginSessionStates = "ERROR"
)

// for schema
var lolQueueEligibilityLoginSessionStatesEnum []interface{}

func init() {
	var res []LolQueueEligibilityLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolQueueEligibilityLoginSessionStatesEnum = append(lolQueueEligibilityLoginSessionStatesEnum, v)
	}
}

func (m LolQueueEligibilityLoginSessionStates) validateLolQueueEligibilityLoginSessionStatesEnum(path, location string, value LolQueueEligibilityLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolQueueEligibilityLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol queue eligibility login session states
func (m LolQueueEligibilityLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolQueueEligibilityLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}