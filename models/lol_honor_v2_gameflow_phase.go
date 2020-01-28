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

// LolHonorV2GameflowPhase lol honor v2 gameflow phase
// swagger:model LolHonorV2GameflowPhase
type LolHonorV2GameflowPhase string

const (

	// LolHonorV2GameflowPhaseNone captures enum value "None"
	LolHonorV2GameflowPhaseNone LolHonorV2GameflowPhase = "None"

	// LolHonorV2GameflowPhaseLobby captures enum value "Lobby"
	LolHonorV2GameflowPhaseLobby LolHonorV2GameflowPhase = "Lobby"

	// LolHonorV2GameflowPhaseMatchmaking captures enum value "Matchmaking"
	LolHonorV2GameflowPhaseMatchmaking LolHonorV2GameflowPhase = "Matchmaking"

	// LolHonorV2GameflowPhaseCheckedIntoTournament captures enum value "CheckedIntoTournament"
	LolHonorV2GameflowPhaseCheckedIntoTournament LolHonorV2GameflowPhase = "CheckedIntoTournament"

	// LolHonorV2GameflowPhaseReadyCheck captures enum value "ReadyCheck"
	LolHonorV2GameflowPhaseReadyCheck LolHonorV2GameflowPhase = "ReadyCheck"

	// LolHonorV2GameflowPhaseChampSelect captures enum value "ChampSelect"
	LolHonorV2GameflowPhaseChampSelect LolHonorV2GameflowPhase = "ChampSelect"

	// LolHonorV2GameflowPhaseGameStart captures enum value "GameStart"
	LolHonorV2GameflowPhaseGameStart LolHonorV2GameflowPhase = "GameStart"

	// LolHonorV2GameflowPhaseFailedToLaunch captures enum value "FailedToLaunch"
	LolHonorV2GameflowPhaseFailedToLaunch LolHonorV2GameflowPhase = "FailedToLaunch"

	// LolHonorV2GameflowPhaseInProgress captures enum value "InProgress"
	LolHonorV2GameflowPhaseInProgress LolHonorV2GameflowPhase = "InProgress"

	// LolHonorV2GameflowPhaseReconnect captures enum value "Reconnect"
	LolHonorV2GameflowPhaseReconnect LolHonorV2GameflowPhase = "Reconnect"

	// LolHonorV2GameflowPhaseWaitingForStats captures enum value "WaitingForStats"
	LolHonorV2GameflowPhaseWaitingForStats LolHonorV2GameflowPhase = "WaitingForStats"

	// LolHonorV2GameflowPhasePreEndOfGame captures enum value "PreEndOfGame"
	LolHonorV2GameflowPhasePreEndOfGame LolHonorV2GameflowPhase = "PreEndOfGame"

	// LolHonorV2GameflowPhaseEndOfGame captures enum value "EndOfGame"
	LolHonorV2GameflowPhaseEndOfGame LolHonorV2GameflowPhase = "EndOfGame"

	// LolHonorV2GameflowPhaseTerminatedInError captures enum value "TerminatedInError"
	LolHonorV2GameflowPhaseTerminatedInError LolHonorV2GameflowPhase = "TerminatedInError"
)

// for schema
var lolHonorV2GameflowPhaseEnum []interface{}

func init() {
	var res []LolHonorV2GameflowPhase
	if err := json.Unmarshal([]byte(`["None","Lobby","Matchmaking","CheckedIntoTournament","ReadyCheck","ChampSelect","GameStart","FailedToLaunch","InProgress","Reconnect","WaitingForStats","PreEndOfGame","EndOfGame","TerminatedInError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolHonorV2GameflowPhaseEnum = append(lolHonorV2GameflowPhaseEnum, v)
	}
}

func (m LolHonorV2GameflowPhase) validateLolHonorV2GameflowPhaseEnum(path, location string, value LolHonorV2GameflowPhase) error {
	if err := validate.Enum(path, location, value, lolHonorV2GameflowPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol honor v2 gameflow phase
func (m LolHonorV2GameflowPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolHonorV2GameflowPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
