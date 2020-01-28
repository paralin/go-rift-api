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

// LolEndOfGameGameflowPhase lol end of game gameflow phase
// swagger:model LolEndOfGameGameflowPhase
type LolEndOfGameGameflowPhase string

const (

	// LolEndOfGameGameflowPhaseNone captures enum value "None"
	LolEndOfGameGameflowPhaseNone LolEndOfGameGameflowPhase = "None"

	// LolEndOfGameGameflowPhaseLobby captures enum value "Lobby"
	LolEndOfGameGameflowPhaseLobby LolEndOfGameGameflowPhase = "Lobby"

	// LolEndOfGameGameflowPhaseMatchmaking captures enum value "Matchmaking"
	LolEndOfGameGameflowPhaseMatchmaking LolEndOfGameGameflowPhase = "Matchmaking"

	// LolEndOfGameGameflowPhaseCheckedIntoTournament captures enum value "CheckedIntoTournament"
	LolEndOfGameGameflowPhaseCheckedIntoTournament LolEndOfGameGameflowPhase = "CheckedIntoTournament"

	// LolEndOfGameGameflowPhaseReadyCheck captures enum value "ReadyCheck"
	LolEndOfGameGameflowPhaseReadyCheck LolEndOfGameGameflowPhase = "ReadyCheck"

	// LolEndOfGameGameflowPhaseChampSelect captures enum value "ChampSelect"
	LolEndOfGameGameflowPhaseChampSelect LolEndOfGameGameflowPhase = "ChampSelect"

	// LolEndOfGameGameflowPhaseGameStart captures enum value "GameStart"
	LolEndOfGameGameflowPhaseGameStart LolEndOfGameGameflowPhase = "GameStart"

	// LolEndOfGameGameflowPhaseFailedToLaunch captures enum value "FailedToLaunch"
	LolEndOfGameGameflowPhaseFailedToLaunch LolEndOfGameGameflowPhase = "FailedToLaunch"

	// LolEndOfGameGameflowPhaseInProgress captures enum value "InProgress"
	LolEndOfGameGameflowPhaseInProgress LolEndOfGameGameflowPhase = "InProgress"

	// LolEndOfGameGameflowPhaseReconnect captures enum value "Reconnect"
	LolEndOfGameGameflowPhaseReconnect LolEndOfGameGameflowPhase = "Reconnect"

	// LolEndOfGameGameflowPhaseWaitingForStats captures enum value "WaitingForStats"
	LolEndOfGameGameflowPhaseWaitingForStats LolEndOfGameGameflowPhase = "WaitingForStats"

	// LolEndOfGameGameflowPhasePreEndOfGame captures enum value "PreEndOfGame"
	LolEndOfGameGameflowPhasePreEndOfGame LolEndOfGameGameflowPhase = "PreEndOfGame"

	// LolEndOfGameGameflowPhaseEndOfGame captures enum value "EndOfGame"
	LolEndOfGameGameflowPhaseEndOfGame LolEndOfGameGameflowPhase = "EndOfGame"

	// LolEndOfGameGameflowPhaseTerminatedInError captures enum value "TerminatedInError"
	LolEndOfGameGameflowPhaseTerminatedInError LolEndOfGameGameflowPhase = "TerminatedInError"
)

// for schema
var lolEndOfGameGameflowPhaseEnum []interface{}

func init() {
	var res []LolEndOfGameGameflowPhase
	if err := json.Unmarshal([]byte(`["None","Lobby","Matchmaking","CheckedIntoTournament","ReadyCheck","ChampSelect","GameStart","FailedToLaunch","InProgress","Reconnect","WaitingForStats","PreEndOfGame","EndOfGame","TerminatedInError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolEndOfGameGameflowPhaseEnum = append(lolEndOfGameGameflowPhaseEnum, v)
	}
}

func (m LolEndOfGameGameflowPhase) validateLolEndOfGameGameflowPhaseEnum(path, location string, value LolEndOfGameGameflowPhase) error {
	if err := validate.Enum(path, location, value, lolEndOfGameGameflowPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol end of game gameflow phase
func (m LolEndOfGameGameflowPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolEndOfGameGameflowPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}