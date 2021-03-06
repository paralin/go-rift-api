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

// LolReplaysGameflowPhase lol replays gameflow phase
// swagger:model LolReplaysGameflowPhase
type LolReplaysGameflowPhase string

const (

	// LolReplaysGameflowPhaseNone captures enum value "None"
	LolReplaysGameflowPhaseNone LolReplaysGameflowPhase = "None"

	// LolReplaysGameflowPhaseLobby captures enum value "Lobby"
	LolReplaysGameflowPhaseLobby LolReplaysGameflowPhase = "Lobby"

	// LolReplaysGameflowPhaseMatchmaking captures enum value "Matchmaking"
	LolReplaysGameflowPhaseMatchmaking LolReplaysGameflowPhase = "Matchmaking"

	// LolReplaysGameflowPhaseCheckedIntoTournament captures enum value "CheckedIntoTournament"
	LolReplaysGameflowPhaseCheckedIntoTournament LolReplaysGameflowPhase = "CheckedIntoTournament"

	// LolReplaysGameflowPhaseReadyCheck captures enum value "ReadyCheck"
	LolReplaysGameflowPhaseReadyCheck LolReplaysGameflowPhase = "ReadyCheck"

	// LolReplaysGameflowPhaseChampSelect captures enum value "ChampSelect"
	LolReplaysGameflowPhaseChampSelect LolReplaysGameflowPhase = "ChampSelect"

	// LolReplaysGameflowPhaseGameStart captures enum value "GameStart"
	LolReplaysGameflowPhaseGameStart LolReplaysGameflowPhase = "GameStart"

	// LolReplaysGameflowPhaseFailedToLaunch captures enum value "FailedToLaunch"
	LolReplaysGameflowPhaseFailedToLaunch LolReplaysGameflowPhase = "FailedToLaunch"

	// LolReplaysGameflowPhaseInProgress captures enum value "InProgress"
	LolReplaysGameflowPhaseInProgress LolReplaysGameflowPhase = "InProgress"

	// LolReplaysGameflowPhaseReconnect captures enum value "Reconnect"
	LolReplaysGameflowPhaseReconnect LolReplaysGameflowPhase = "Reconnect"

	// LolReplaysGameflowPhaseWaitingForStats captures enum value "WaitingForStats"
	LolReplaysGameflowPhaseWaitingForStats LolReplaysGameflowPhase = "WaitingForStats"

	// LolReplaysGameflowPhasePreEndOfGame captures enum value "PreEndOfGame"
	LolReplaysGameflowPhasePreEndOfGame LolReplaysGameflowPhase = "PreEndOfGame"

	// LolReplaysGameflowPhaseEndOfGame captures enum value "EndOfGame"
	LolReplaysGameflowPhaseEndOfGame LolReplaysGameflowPhase = "EndOfGame"

	// LolReplaysGameflowPhaseTerminatedInError captures enum value "TerminatedInError"
	LolReplaysGameflowPhaseTerminatedInError LolReplaysGameflowPhase = "TerminatedInError"
)

// for schema
var lolReplaysGameflowPhaseEnum []interface{}

func init() {
	var res []LolReplaysGameflowPhase
	if err := json.Unmarshal([]byte(`["None","Lobby","Matchmaking","CheckedIntoTournament","ReadyCheck","ChampSelect","GameStart","FailedToLaunch","InProgress","Reconnect","WaitingForStats","PreEndOfGame","EndOfGame","TerminatedInError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolReplaysGameflowPhaseEnum = append(lolReplaysGameflowPhaseEnum, v)
	}
}

func (m LolReplaysGameflowPhase) validateLolReplaysGameflowPhaseEnum(path, location string, value LolReplaysGameflowPhase) error {
	if err := validate.Enum(path, location, value, lolReplaysGameflowPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol replays gameflow phase
func (m LolReplaysGameflowPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolReplaysGameflowPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
