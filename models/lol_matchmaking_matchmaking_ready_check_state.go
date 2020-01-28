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

// LolMatchmakingMatchmakingReadyCheckState lol matchmaking matchmaking ready check state
// swagger:model LolMatchmakingMatchmakingReadyCheckState
type LolMatchmakingMatchmakingReadyCheckState string

const (

	// LolMatchmakingMatchmakingReadyCheckStateInvalid captures enum value "Invalid"
	LolMatchmakingMatchmakingReadyCheckStateInvalid LolMatchmakingMatchmakingReadyCheckState = "Invalid"

	// LolMatchmakingMatchmakingReadyCheckStateInProgress captures enum value "InProgress"
	LolMatchmakingMatchmakingReadyCheckStateInProgress LolMatchmakingMatchmakingReadyCheckState = "InProgress"

	// LolMatchmakingMatchmakingReadyCheckStateEveryoneReady captures enum value "EveryoneReady"
	LolMatchmakingMatchmakingReadyCheckStateEveryoneReady LolMatchmakingMatchmakingReadyCheckState = "EveryoneReady"

	// LolMatchmakingMatchmakingReadyCheckStateStrangerNotReady captures enum value "StrangerNotReady"
	LolMatchmakingMatchmakingReadyCheckStateStrangerNotReady LolMatchmakingMatchmakingReadyCheckState = "StrangerNotReady"

	// LolMatchmakingMatchmakingReadyCheckStatePartyNotReady captures enum value "PartyNotReady"
	LolMatchmakingMatchmakingReadyCheckStatePartyNotReady LolMatchmakingMatchmakingReadyCheckState = "PartyNotReady"

	// LolMatchmakingMatchmakingReadyCheckStateError captures enum value "Error"
	LolMatchmakingMatchmakingReadyCheckStateError LolMatchmakingMatchmakingReadyCheckState = "Error"
)

// for schema
var lolMatchmakingMatchmakingReadyCheckStateEnum []interface{}

func init() {
	var res []LolMatchmakingMatchmakingReadyCheckState
	if err := json.Unmarshal([]byte(`["Invalid","InProgress","EveryoneReady","StrangerNotReady","PartyNotReady","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolMatchmakingMatchmakingReadyCheckStateEnum = append(lolMatchmakingMatchmakingReadyCheckStateEnum, v)
	}
}

func (m LolMatchmakingMatchmakingReadyCheckState) validateLolMatchmakingMatchmakingReadyCheckStateEnum(path, location string, value LolMatchmakingMatchmakingReadyCheckState) error {
	if err := validate.Enum(path, location, value, lolMatchmakingMatchmakingReadyCheckStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol matchmaking matchmaking ready check state
func (m LolMatchmakingMatchmakingReadyCheckState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolMatchmakingMatchmakingReadyCheckStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
