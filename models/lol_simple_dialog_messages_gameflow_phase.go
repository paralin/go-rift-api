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

// LolSimpleDialogMessagesGameflowPhase lol simple dialog messages gameflow phase
// swagger:model LolSimpleDialogMessagesGameflowPhase
type LolSimpleDialogMessagesGameflowPhase string

const (

	// LolSimpleDialogMessagesGameflowPhaseNone captures enum value "None"
	LolSimpleDialogMessagesGameflowPhaseNone LolSimpleDialogMessagesGameflowPhase = "None"

	// LolSimpleDialogMessagesGameflowPhaseLobby captures enum value "Lobby"
	LolSimpleDialogMessagesGameflowPhaseLobby LolSimpleDialogMessagesGameflowPhase = "Lobby"

	// LolSimpleDialogMessagesGameflowPhaseMatchmaking captures enum value "Matchmaking"
	LolSimpleDialogMessagesGameflowPhaseMatchmaking LolSimpleDialogMessagesGameflowPhase = "Matchmaking"

	// LolSimpleDialogMessagesGameflowPhaseCheckedIntoTournament captures enum value "CheckedIntoTournament"
	LolSimpleDialogMessagesGameflowPhaseCheckedIntoTournament LolSimpleDialogMessagesGameflowPhase = "CheckedIntoTournament"

	// LolSimpleDialogMessagesGameflowPhaseReadyCheck captures enum value "ReadyCheck"
	LolSimpleDialogMessagesGameflowPhaseReadyCheck LolSimpleDialogMessagesGameflowPhase = "ReadyCheck"

	// LolSimpleDialogMessagesGameflowPhaseChampSelect captures enum value "ChampSelect"
	LolSimpleDialogMessagesGameflowPhaseChampSelect LolSimpleDialogMessagesGameflowPhase = "ChampSelect"

	// LolSimpleDialogMessagesGameflowPhaseGameStart captures enum value "GameStart"
	LolSimpleDialogMessagesGameflowPhaseGameStart LolSimpleDialogMessagesGameflowPhase = "GameStart"

	// LolSimpleDialogMessagesGameflowPhaseFailedToLaunch captures enum value "FailedToLaunch"
	LolSimpleDialogMessagesGameflowPhaseFailedToLaunch LolSimpleDialogMessagesGameflowPhase = "FailedToLaunch"

	// LolSimpleDialogMessagesGameflowPhaseInProgress captures enum value "InProgress"
	LolSimpleDialogMessagesGameflowPhaseInProgress LolSimpleDialogMessagesGameflowPhase = "InProgress"

	// LolSimpleDialogMessagesGameflowPhaseReconnect captures enum value "Reconnect"
	LolSimpleDialogMessagesGameflowPhaseReconnect LolSimpleDialogMessagesGameflowPhase = "Reconnect"

	// LolSimpleDialogMessagesGameflowPhaseWaitingForStats captures enum value "WaitingForStats"
	LolSimpleDialogMessagesGameflowPhaseWaitingForStats LolSimpleDialogMessagesGameflowPhase = "WaitingForStats"

	// LolSimpleDialogMessagesGameflowPhasePreEndOfGame captures enum value "PreEndOfGame"
	LolSimpleDialogMessagesGameflowPhasePreEndOfGame LolSimpleDialogMessagesGameflowPhase = "PreEndOfGame"

	// LolSimpleDialogMessagesGameflowPhaseEndOfGame captures enum value "EndOfGame"
	LolSimpleDialogMessagesGameflowPhaseEndOfGame LolSimpleDialogMessagesGameflowPhase = "EndOfGame"

	// LolSimpleDialogMessagesGameflowPhaseTerminatedInError captures enum value "TerminatedInError"
	LolSimpleDialogMessagesGameflowPhaseTerminatedInError LolSimpleDialogMessagesGameflowPhase = "TerminatedInError"
)

// for schema
var lolSimpleDialogMessagesGameflowPhaseEnum []interface{}

func init() {
	var res []LolSimpleDialogMessagesGameflowPhase
	if err := json.Unmarshal([]byte(`["None","Lobby","Matchmaking","CheckedIntoTournament","ReadyCheck","ChampSelect","GameStart","FailedToLaunch","InProgress","Reconnect","WaitingForStats","PreEndOfGame","EndOfGame","TerminatedInError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolSimpleDialogMessagesGameflowPhaseEnum = append(lolSimpleDialogMessagesGameflowPhaseEnum, v)
	}
}

func (m LolSimpleDialogMessagesGameflowPhase) validateLolSimpleDialogMessagesGameflowPhaseEnum(path, location string, value LolSimpleDialogMessagesGameflowPhase) error {
	if err := validate.Enum(path, location, value, lolSimpleDialogMessagesGameflowPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol simple dialog messages gameflow phase
func (m LolSimpleDialogMessagesGameflowPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolSimpleDialogMessagesGameflowPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
