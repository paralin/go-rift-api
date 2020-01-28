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

// LolChatQueueCustomGameSpectatorPolicy lol chat queue custom game spectator policy
// swagger:model LolChatQueueCustomGameSpectatorPolicy
type LolChatQueueCustomGameSpectatorPolicy string

const (

	// LolChatQueueCustomGameSpectatorPolicyNotAllowed captures enum value "NotAllowed"
	LolChatQueueCustomGameSpectatorPolicyNotAllowed LolChatQueueCustomGameSpectatorPolicy = "NotAllowed"

	// LolChatQueueCustomGameSpectatorPolicyLobbyAllowed captures enum value "LobbyAllowed"
	LolChatQueueCustomGameSpectatorPolicyLobbyAllowed LolChatQueueCustomGameSpectatorPolicy = "LobbyAllowed"

	// LolChatQueueCustomGameSpectatorPolicyFriendsAllowed captures enum value "FriendsAllowed"
	LolChatQueueCustomGameSpectatorPolicyFriendsAllowed LolChatQueueCustomGameSpectatorPolicy = "FriendsAllowed"

	// LolChatQueueCustomGameSpectatorPolicyAllAllowed captures enum value "AllAllowed"
	LolChatQueueCustomGameSpectatorPolicyAllAllowed LolChatQueueCustomGameSpectatorPolicy = "AllAllowed"
)

// for schema
var lolChatQueueCustomGameSpectatorPolicyEnum []interface{}

func init() {
	var res []LolChatQueueCustomGameSpectatorPolicy
	if err := json.Unmarshal([]byte(`["NotAllowed","LobbyAllowed","FriendsAllowed","AllAllowed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolChatQueueCustomGameSpectatorPolicyEnum = append(lolChatQueueCustomGameSpectatorPolicyEnum, v)
	}
}

func (m LolChatQueueCustomGameSpectatorPolicy) validateLolChatQueueCustomGameSpectatorPolicyEnum(path, location string, value LolChatQueueCustomGameSpectatorPolicy) error {
	if err := validate.Enum(path, location, value, lolChatQueueCustomGameSpectatorPolicyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol chat queue custom game spectator policy
func (m LolChatQueueCustomGameSpectatorPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolChatQueueCustomGameSpectatorPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}