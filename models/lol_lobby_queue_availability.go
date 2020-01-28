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

// LolLobbyQueueAvailability lol lobby queue availability
// swagger:model LolLobbyQueueAvailability
type LolLobbyQueueAvailability string

const (

	// LolLobbyQueueAvailabilityAvailable captures enum value "Available"
	LolLobbyQueueAvailabilityAvailable LolLobbyQueueAvailability = "Available"

	// LolLobbyQueueAvailabilityPlatformDisabled captures enum value "PlatformDisabled"
	LolLobbyQueueAvailabilityPlatformDisabled LolLobbyQueueAvailability = "PlatformDisabled"

	// LolLobbyQueueAvailabilityDoesntMeetRequirements captures enum value "DoesntMeetRequirements"
	LolLobbyQueueAvailabilityDoesntMeetRequirements LolLobbyQueueAvailability = "DoesntMeetRequirements"
)

// for schema
var lolLobbyQueueAvailabilityEnum []interface{}

func init() {
	var res []LolLobbyQueueAvailability
	if err := json.Unmarshal([]byte(`["Available","PlatformDisabled","DoesntMeetRequirements"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLobbyQueueAvailabilityEnum = append(lolLobbyQueueAvailabilityEnum, v)
	}
}

func (m LolLobbyQueueAvailability) validateLolLobbyQueueAvailabilityEnum(path, location string, value LolLobbyQueueAvailability) error {
	if err := validate.Enum(path, location, value, lolLobbyQueueAvailabilityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol lobby queue availability
func (m LolLobbyQueueAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLobbyQueueAvailabilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
