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

// LolGameflowGameflowPhase lol gameflow gameflow phase
// swagger:model LolGameflowGameflowPhase
type LolGameflowGameflowPhase string

const (

	// LolGameflowGameflowPhaseNone captures enum value "None"
	LolGameflowGameflowPhaseNone LolGameflowGameflowPhase = "None"

	// LolGameflowGameflowPhaseLobby captures enum value "Lobby"
	LolGameflowGameflowPhaseLobby LolGameflowGameflowPhase = "Lobby"

	// LolGameflowGameflowPhaseMatchmaking captures enum value "Matchmaking"
	LolGameflowGameflowPhaseMatchmaking LolGameflowGameflowPhase = "Matchmaking"

	// LolGameflowGameflowPhaseCheckedIntoTournament captures enum value "CheckedIntoTournament"
	LolGameflowGameflowPhaseCheckedIntoTournament LolGameflowGameflowPhase = "CheckedIntoTournament"

	// LolGameflowGameflowPhaseReadyCheck captures enum value "ReadyCheck"
	LolGameflowGameflowPhaseReadyCheck LolGameflowGameflowPhase = "ReadyCheck"

	// LolGameflowGameflowPhaseChampSelect captures enum value "ChampSelect"
	LolGameflowGameflowPhaseChampSelect LolGameflowGameflowPhase = "ChampSelect"

	// LolGameflowGameflowPhaseGameStart captures enum value "GameStart"
	LolGameflowGameflowPhaseGameStart LolGameflowGameflowPhase = "GameStart"

	// LolGameflowGameflowPhaseFailedToLaunch captures enum value "FailedToLaunch"
	LolGameflowGameflowPhaseFailedToLaunch LolGameflowGameflowPhase = "FailedToLaunch"

	// LolGameflowGameflowPhaseInProgress captures enum value "InProgress"
	LolGameflowGameflowPhaseInProgress LolGameflowGameflowPhase = "InProgress"

	// LolGameflowGameflowPhaseReconnect captures enum value "Reconnect"
	LolGameflowGameflowPhaseReconnect LolGameflowGameflowPhase = "Reconnect"

	// LolGameflowGameflowPhaseWaitingForStats captures enum value "WaitingForStats"
	LolGameflowGameflowPhaseWaitingForStats LolGameflowGameflowPhase = "WaitingForStats"

	// LolGameflowGameflowPhasePreEndOfGame captures enum value "PreEndOfGame"
	LolGameflowGameflowPhasePreEndOfGame LolGameflowGameflowPhase = "PreEndOfGame"

	// LolGameflowGameflowPhaseEndOfGame captures enum value "EndOfGame"
	LolGameflowGameflowPhaseEndOfGame LolGameflowGameflowPhase = "EndOfGame"

	// LolGameflowGameflowPhaseTerminatedInError captures enum value "TerminatedInError"
	LolGameflowGameflowPhaseTerminatedInError LolGameflowGameflowPhase = "TerminatedInError"
)

// for schema
var lolGameflowGameflowPhaseEnum []interface{}

func init() {
	var res []LolGameflowGameflowPhase
	if err := json.Unmarshal([]byte(`["None","Lobby","Matchmaking","CheckedIntoTournament","ReadyCheck","ChampSelect","GameStart","FailedToLaunch","InProgress","Reconnect","WaitingForStats","PreEndOfGame","EndOfGame","TerminatedInError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolGameflowGameflowPhaseEnum = append(lolGameflowGameflowPhaseEnum, v)
	}
}

func (m LolGameflowGameflowPhase) validateLolGameflowGameflowPhaseEnum(path, location string, value LolGameflowGameflowPhase) error {
	if err := validate.Enum(path, location, value, lolGameflowGameflowPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol gameflow gameflow phase
func (m LolGameflowGameflowPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolGameflowGameflowPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}