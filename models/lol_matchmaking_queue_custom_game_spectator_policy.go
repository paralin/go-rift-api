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

// LolMatchmakingQueueCustomGameSpectatorPolicy lol matchmaking queue custom game spectator policy
// swagger:model LolMatchmakingQueueCustomGameSpectatorPolicy
type LolMatchmakingQueueCustomGameSpectatorPolicy string

const (

	// LolMatchmakingQueueCustomGameSpectatorPolicyNotAllowed captures enum value "NotAllowed"
	LolMatchmakingQueueCustomGameSpectatorPolicyNotAllowed LolMatchmakingQueueCustomGameSpectatorPolicy = "NotAllowed"

	// LolMatchmakingQueueCustomGameSpectatorPolicyLobbyAllowed captures enum value "LobbyAllowed"
	LolMatchmakingQueueCustomGameSpectatorPolicyLobbyAllowed LolMatchmakingQueueCustomGameSpectatorPolicy = "LobbyAllowed"

	// LolMatchmakingQueueCustomGameSpectatorPolicyFriendsAllowed captures enum value "FriendsAllowed"
	LolMatchmakingQueueCustomGameSpectatorPolicyFriendsAllowed LolMatchmakingQueueCustomGameSpectatorPolicy = "FriendsAllowed"

	// LolMatchmakingQueueCustomGameSpectatorPolicyAllAllowed captures enum value "AllAllowed"
	LolMatchmakingQueueCustomGameSpectatorPolicyAllAllowed LolMatchmakingQueueCustomGameSpectatorPolicy = "AllAllowed"
)

// for schema
var lolMatchmakingQueueCustomGameSpectatorPolicyEnum []interface{}

func init() {
	var res []LolMatchmakingQueueCustomGameSpectatorPolicy
	if err := json.Unmarshal([]byte(`["NotAllowed","LobbyAllowed","FriendsAllowed","AllAllowed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolMatchmakingQueueCustomGameSpectatorPolicyEnum = append(lolMatchmakingQueueCustomGameSpectatorPolicyEnum, v)
	}
}

func (m LolMatchmakingQueueCustomGameSpectatorPolicy) validateLolMatchmakingQueueCustomGameSpectatorPolicyEnum(path, location string, value LolMatchmakingQueueCustomGameSpectatorPolicy) error {
	if err := validate.Enum(path, location, value, lolMatchmakingQueueCustomGameSpectatorPolicyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol matchmaking queue custom game spectator policy
func (m LolMatchmakingQueueCustomGameSpectatorPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolMatchmakingQueueCustomGameSpectatorPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}