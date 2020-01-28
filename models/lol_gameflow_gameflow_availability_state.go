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

// LolGameflowGameflowAvailabilityState lol gameflow gameflow availability state
// swagger:model LolGameflowGameflowAvailabilityState
type LolGameflowGameflowAvailabilityState string

const (

	// LolGameflowGameflowAvailabilityStateAvailable captures enum value "Available"
	LolGameflowGameflowAvailabilityStateAvailable LolGameflowGameflowAvailabilityState = "Available"

	// LolGameflowGameflowAvailabilityStateInitializing captures enum value "Initializing"
	LolGameflowGameflowAvailabilityStateInitializing LolGameflowGameflowAvailabilityState = "Initializing"

	// LolGameflowGameflowAvailabilityStatePatching captures enum value "Patching"
	LolGameflowGameflowAvailabilityStatePatching LolGameflowGameflowAvailabilityState = "Patching"

	// LolGameflowGameflowAvailabilityStatePlayerBanned captures enum value "PlayerBanned"
	LolGameflowGameflowAvailabilityStatePlayerBanned LolGameflowGameflowAvailabilityState = "PlayerBanned"

	// LolGameflowGameflowAvailabilityStateInGameFlow captures enum value "InGameFlow"
	LolGameflowGameflowAvailabilityStateInGameFlow LolGameflowGameflowAvailabilityState = "InGameFlow"

	// LolGameflowGameflowAvailabilityStateConfiguration captures enum value "Configuration"
	LolGameflowGameflowAvailabilityStateConfiguration LolGameflowGameflowAvailabilityState = "Configuration"

	// LolGameflowGameflowAvailabilityStateEligibilityInfoMissing captures enum value "EligibilityInfoMissing"
	LolGameflowGameflowAvailabilityStateEligibilityInfoMissing LolGameflowGameflowAvailabilityState = "EligibilityInfoMissing"
)

// for schema
var lolGameflowGameflowAvailabilityStateEnum []interface{}

func init() {
	var res []LolGameflowGameflowAvailabilityState
	if err := json.Unmarshal([]byte(`["Available","Initializing","Patching","PlayerBanned","InGameFlow","Configuration","EligibilityInfoMissing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolGameflowGameflowAvailabilityStateEnum = append(lolGameflowGameflowAvailabilityStateEnum, v)
	}
}

func (m LolGameflowGameflowAvailabilityState) validateLolGameflowGameflowAvailabilityStateEnum(path, location string, value LolGameflowGameflowAvailabilityState) error {
	if err := validate.Enum(path, location, value, lolGameflowGameflowAvailabilityStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol gameflow gameflow availability state
func (m LolGameflowGameflowAvailabilityState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolGameflowGameflowAvailabilityStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}