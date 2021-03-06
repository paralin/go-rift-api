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

// LolLootGameflowPhase lol loot gameflow phase
// swagger:model LolLootGameflowPhase
type LolLootGameflowPhase string

const (

	// LolLootGameflowPhaseNone captures enum value "None"
	LolLootGameflowPhaseNone LolLootGameflowPhase = "None"

	// LolLootGameflowPhaseLobby captures enum value "Lobby"
	LolLootGameflowPhaseLobby LolLootGameflowPhase = "Lobby"

	// LolLootGameflowPhaseMatchmaking captures enum value "Matchmaking"
	LolLootGameflowPhaseMatchmaking LolLootGameflowPhase = "Matchmaking"

	// LolLootGameflowPhaseCheckedIntoTournament captures enum value "CheckedIntoTournament"
	LolLootGameflowPhaseCheckedIntoTournament LolLootGameflowPhase = "CheckedIntoTournament"

	// LolLootGameflowPhaseReadyCheck captures enum value "ReadyCheck"
	LolLootGameflowPhaseReadyCheck LolLootGameflowPhase = "ReadyCheck"

	// LolLootGameflowPhaseChampSelect captures enum value "ChampSelect"
	LolLootGameflowPhaseChampSelect LolLootGameflowPhase = "ChampSelect"

	// LolLootGameflowPhaseGameStart captures enum value "GameStart"
	LolLootGameflowPhaseGameStart LolLootGameflowPhase = "GameStart"

	// LolLootGameflowPhaseFailedToLaunch captures enum value "FailedToLaunch"
	LolLootGameflowPhaseFailedToLaunch LolLootGameflowPhase = "FailedToLaunch"

	// LolLootGameflowPhaseInProgress captures enum value "InProgress"
	LolLootGameflowPhaseInProgress LolLootGameflowPhase = "InProgress"

	// LolLootGameflowPhaseReconnect captures enum value "Reconnect"
	LolLootGameflowPhaseReconnect LolLootGameflowPhase = "Reconnect"

	// LolLootGameflowPhaseWaitingForStats captures enum value "WaitingForStats"
	LolLootGameflowPhaseWaitingForStats LolLootGameflowPhase = "WaitingForStats"

	// LolLootGameflowPhasePreEndOfGame captures enum value "PreEndOfGame"
	LolLootGameflowPhasePreEndOfGame LolLootGameflowPhase = "PreEndOfGame"

	// LolLootGameflowPhaseEndOfGame captures enum value "EndOfGame"
	LolLootGameflowPhaseEndOfGame LolLootGameflowPhase = "EndOfGame"

	// LolLootGameflowPhaseTerminatedInError captures enum value "TerminatedInError"
	LolLootGameflowPhaseTerminatedInError LolLootGameflowPhase = "TerminatedInError"
)

// for schema
var lolLootGameflowPhaseEnum []interface{}

func init() {
	var res []LolLootGameflowPhase
	if err := json.Unmarshal([]byte(`["None","Lobby","Matchmaking","CheckedIntoTournament","ReadyCheck","ChampSelect","GameStart","FailedToLaunch","InProgress","Reconnect","WaitingForStats","PreEndOfGame","EndOfGame","TerminatedInError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLootGameflowPhaseEnum = append(lolLootGameflowPhaseEnum, v)
	}
}

func (m LolLootGameflowPhase) validateLolLootGameflowPhaseEnum(path, location string, value LolLootGameflowPhase) error {
	if err := validate.Enum(path, location, value, lolLootGameflowPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol loot gameflow phase
func (m LolLootGameflowPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLootGameflowPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
