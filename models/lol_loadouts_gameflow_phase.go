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

// LolLoadoutsGameflowPhase lol loadouts gameflow phase
// swagger:model LolLoadoutsGameflowPhase
type LolLoadoutsGameflowPhase string

const (

	// LolLoadoutsGameflowPhaseNone captures enum value "None"
	LolLoadoutsGameflowPhaseNone LolLoadoutsGameflowPhase = "None"

	// LolLoadoutsGameflowPhaseLobby captures enum value "Lobby"
	LolLoadoutsGameflowPhaseLobby LolLoadoutsGameflowPhase = "Lobby"

	// LolLoadoutsGameflowPhaseMatchmaking captures enum value "Matchmaking"
	LolLoadoutsGameflowPhaseMatchmaking LolLoadoutsGameflowPhase = "Matchmaking"

	// LolLoadoutsGameflowPhaseCheckedIntoTournament captures enum value "CheckedIntoTournament"
	LolLoadoutsGameflowPhaseCheckedIntoTournament LolLoadoutsGameflowPhase = "CheckedIntoTournament"

	// LolLoadoutsGameflowPhaseReadyCheck captures enum value "ReadyCheck"
	LolLoadoutsGameflowPhaseReadyCheck LolLoadoutsGameflowPhase = "ReadyCheck"

	// LolLoadoutsGameflowPhaseChampSelect captures enum value "ChampSelect"
	LolLoadoutsGameflowPhaseChampSelect LolLoadoutsGameflowPhase = "ChampSelect"

	// LolLoadoutsGameflowPhaseGameStart captures enum value "GameStart"
	LolLoadoutsGameflowPhaseGameStart LolLoadoutsGameflowPhase = "GameStart"

	// LolLoadoutsGameflowPhaseFailedToLaunch captures enum value "FailedToLaunch"
	LolLoadoutsGameflowPhaseFailedToLaunch LolLoadoutsGameflowPhase = "FailedToLaunch"

	// LolLoadoutsGameflowPhaseInProgress captures enum value "InProgress"
	LolLoadoutsGameflowPhaseInProgress LolLoadoutsGameflowPhase = "InProgress"

	// LolLoadoutsGameflowPhaseReconnect captures enum value "Reconnect"
	LolLoadoutsGameflowPhaseReconnect LolLoadoutsGameflowPhase = "Reconnect"

	// LolLoadoutsGameflowPhaseWaitingForStats captures enum value "WaitingForStats"
	LolLoadoutsGameflowPhaseWaitingForStats LolLoadoutsGameflowPhase = "WaitingForStats"

	// LolLoadoutsGameflowPhasePreEndOfGame captures enum value "PreEndOfGame"
	LolLoadoutsGameflowPhasePreEndOfGame LolLoadoutsGameflowPhase = "PreEndOfGame"

	// LolLoadoutsGameflowPhaseEndOfGame captures enum value "EndOfGame"
	LolLoadoutsGameflowPhaseEndOfGame LolLoadoutsGameflowPhase = "EndOfGame"

	// LolLoadoutsGameflowPhaseTerminatedInError captures enum value "TerminatedInError"
	LolLoadoutsGameflowPhaseTerminatedInError LolLoadoutsGameflowPhase = "TerminatedInError"
)

// for schema
var lolLoadoutsGameflowPhaseEnum []interface{}

func init() {
	var res []LolLoadoutsGameflowPhase
	if err := json.Unmarshal([]byte(`["None","Lobby","Matchmaking","CheckedIntoTournament","ReadyCheck","ChampSelect","GameStart","FailedToLaunch","InProgress","Reconnect","WaitingForStats","PreEndOfGame","EndOfGame","TerminatedInError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLoadoutsGameflowPhaseEnum = append(lolLoadoutsGameflowPhaseEnum, v)
	}
}

func (m LolLoadoutsGameflowPhase) validateLolLoadoutsGameflowPhaseEnum(path, location string, value LolLoadoutsGameflowPhase) error {
	if err := validate.Enum(path, location, value, lolLoadoutsGameflowPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol loadouts gameflow phase
func (m LolLoadoutsGameflowPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLoadoutsGameflowPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
