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

// LolLobbyTeamBuilderMatchmakingReadyCheckState lol lobby team builder matchmaking ready check state
// swagger:model LolLobbyTeamBuilderMatchmakingReadyCheckState
type LolLobbyTeamBuilderMatchmakingReadyCheckState string

const (

	// LolLobbyTeamBuilderMatchmakingReadyCheckStateInvalid captures enum value "Invalid"
	LolLobbyTeamBuilderMatchmakingReadyCheckStateInvalid LolLobbyTeamBuilderMatchmakingReadyCheckState = "Invalid"

	// LolLobbyTeamBuilderMatchmakingReadyCheckStateInProgress captures enum value "InProgress"
	LolLobbyTeamBuilderMatchmakingReadyCheckStateInProgress LolLobbyTeamBuilderMatchmakingReadyCheckState = "InProgress"

	// LolLobbyTeamBuilderMatchmakingReadyCheckStateEveryoneReady captures enum value "EveryoneReady"
	LolLobbyTeamBuilderMatchmakingReadyCheckStateEveryoneReady LolLobbyTeamBuilderMatchmakingReadyCheckState = "EveryoneReady"

	// LolLobbyTeamBuilderMatchmakingReadyCheckStateStrangerNotReady captures enum value "StrangerNotReady"
	LolLobbyTeamBuilderMatchmakingReadyCheckStateStrangerNotReady LolLobbyTeamBuilderMatchmakingReadyCheckState = "StrangerNotReady"

	// LolLobbyTeamBuilderMatchmakingReadyCheckStatePartyNotReady captures enum value "PartyNotReady"
	LolLobbyTeamBuilderMatchmakingReadyCheckStatePartyNotReady LolLobbyTeamBuilderMatchmakingReadyCheckState = "PartyNotReady"

	// LolLobbyTeamBuilderMatchmakingReadyCheckStateError captures enum value "Error"
	LolLobbyTeamBuilderMatchmakingReadyCheckStateError LolLobbyTeamBuilderMatchmakingReadyCheckState = "Error"
)

// for schema
var lolLobbyTeamBuilderMatchmakingReadyCheckStateEnum []interface{}

func init() {
	var res []LolLobbyTeamBuilderMatchmakingReadyCheckState
	if err := json.Unmarshal([]byte(`["Invalid","InProgress","EveryoneReady","StrangerNotReady","PartyNotReady","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLobbyTeamBuilderMatchmakingReadyCheckStateEnum = append(lolLobbyTeamBuilderMatchmakingReadyCheckStateEnum, v)
	}
}

func (m LolLobbyTeamBuilderMatchmakingReadyCheckState) validateLolLobbyTeamBuilderMatchmakingReadyCheckStateEnum(path, location string, value LolLobbyTeamBuilderMatchmakingReadyCheckState) error {
	if err := validate.Enum(path, location, value, lolLobbyTeamBuilderMatchmakingReadyCheckStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol lobby team builder matchmaking ready check state
func (m LolLobbyTeamBuilderMatchmakingReadyCheckState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLobbyTeamBuilderMatchmakingReadyCheckStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
