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

// LolMatchmakingMatchmakingDodgeState lol matchmaking matchmaking dodge state
// swagger:model LolMatchmakingMatchmakingDodgeState
type LolMatchmakingMatchmakingDodgeState string

const (

	// LolMatchmakingMatchmakingDodgeStateInvalid captures enum value "Invalid"
	LolMatchmakingMatchmakingDodgeStateInvalid LolMatchmakingMatchmakingDodgeState = "Invalid"

	// LolMatchmakingMatchmakingDodgeStatePartyDodged captures enum value "PartyDodged"
	LolMatchmakingMatchmakingDodgeStatePartyDodged LolMatchmakingMatchmakingDodgeState = "PartyDodged"

	// LolMatchmakingMatchmakingDodgeStateStrangerDodged captures enum value "StrangerDodged"
	LolMatchmakingMatchmakingDodgeStateStrangerDodged LolMatchmakingMatchmakingDodgeState = "StrangerDodged"

	// LolMatchmakingMatchmakingDodgeStateTournamentDodged captures enum value "TournamentDodged"
	LolMatchmakingMatchmakingDodgeStateTournamentDodged LolMatchmakingMatchmakingDodgeState = "TournamentDodged"
)

// for schema
var lolMatchmakingMatchmakingDodgeStateEnum []interface{}

func init() {
	var res []LolMatchmakingMatchmakingDodgeState
	if err := json.Unmarshal([]byte(`["Invalid","PartyDodged","StrangerDodged","TournamentDodged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolMatchmakingMatchmakingDodgeStateEnum = append(lolMatchmakingMatchmakingDodgeStateEnum, v)
	}
}

func (m LolMatchmakingMatchmakingDodgeState) validateLolMatchmakingMatchmakingDodgeStateEnum(path, location string, value LolMatchmakingMatchmakingDodgeState) error {
	if err := validate.Enum(path, location, value, lolMatchmakingMatchmakingDodgeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol matchmaking matchmaking dodge state
func (m LolMatchmakingMatchmakingDodgeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolMatchmakingMatchmakingDodgeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}