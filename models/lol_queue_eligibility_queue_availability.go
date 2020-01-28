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

// LolQueueEligibilityQueueAvailability lol queue eligibility queue availability
// swagger:model LolQueueEligibilityQueueAvailability
type LolQueueEligibilityQueueAvailability string

const (

	// LolQueueEligibilityQueueAvailabilityAvailable captures enum value "Available"
	LolQueueEligibilityQueueAvailabilityAvailable LolQueueEligibilityQueueAvailability = "Available"

	// LolQueueEligibilityQueueAvailabilityPlatformDisabled captures enum value "PlatformDisabled"
	LolQueueEligibilityQueueAvailabilityPlatformDisabled LolQueueEligibilityQueueAvailability = "PlatformDisabled"

	// LolQueueEligibilityQueueAvailabilityDoesntMeetRequirements captures enum value "DoesntMeetRequirements"
	LolQueueEligibilityQueueAvailabilityDoesntMeetRequirements LolQueueEligibilityQueueAvailability = "DoesntMeetRequirements"
)

// for schema
var lolQueueEligibilityQueueAvailabilityEnum []interface{}

func init() {
	var res []LolQueueEligibilityQueueAvailability
	if err := json.Unmarshal([]byte(`["Available","PlatformDisabled","DoesntMeetRequirements"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolQueueEligibilityQueueAvailabilityEnum = append(lolQueueEligibilityQueueAvailabilityEnum, v)
	}
}

func (m LolQueueEligibilityQueueAvailability) validateLolQueueEligibilityQueueAvailabilityEnum(path, location string, value LolQueueEligibilityQueueAvailability) error {
	if err := validate.Enum(path, location, value, lolQueueEligibilityQueueAvailabilityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol queue eligibility queue availability
func (m LolQueueEligibilityQueueAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolQueueEligibilityQueueAvailabilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}