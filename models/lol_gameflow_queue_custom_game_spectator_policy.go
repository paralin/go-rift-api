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

// LolGameflowQueueCustomGameSpectatorPolicy lol gameflow queue custom game spectator policy
// swagger:model LolGameflowQueueCustomGameSpectatorPolicy
type LolGameflowQueueCustomGameSpectatorPolicy string

const (

	// LolGameflowQueueCustomGameSpectatorPolicyNotAllowed captures enum value "NotAllowed"
	LolGameflowQueueCustomGameSpectatorPolicyNotAllowed LolGameflowQueueCustomGameSpectatorPolicy = "NotAllowed"

	// LolGameflowQueueCustomGameSpectatorPolicyLobbyAllowed captures enum value "LobbyAllowed"
	LolGameflowQueueCustomGameSpectatorPolicyLobbyAllowed LolGameflowQueueCustomGameSpectatorPolicy = "LobbyAllowed"

	// LolGameflowQueueCustomGameSpectatorPolicyFriendsAllowed captures enum value "FriendsAllowed"
	LolGameflowQueueCustomGameSpectatorPolicyFriendsAllowed LolGameflowQueueCustomGameSpectatorPolicy = "FriendsAllowed"

	// LolGameflowQueueCustomGameSpectatorPolicyAllAllowed captures enum value "AllAllowed"
	LolGameflowQueueCustomGameSpectatorPolicyAllAllowed LolGameflowQueueCustomGameSpectatorPolicy = "AllAllowed"
)

// for schema
var lolGameflowQueueCustomGameSpectatorPolicyEnum []interface{}

func init() {
	var res []LolGameflowQueueCustomGameSpectatorPolicy
	if err := json.Unmarshal([]byte(`["NotAllowed","LobbyAllowed","FriendsAllowed","AllAllowed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolGameflowQueueCustomGameSpectatorPolicyEnum = append(lolGameflowQueueCustomGameSpectatorPolicyEnum, v)
	}
}

func (m LolGameflowQueueCustomGameSpectatorPolicy) validateLolGameflowQueueCustomGameSpectatorPolicyEnum(path, location string, value LolGameflowQueueCustomGameSpectatorPolicy) error {
	if err := validate.Enum(path, location, value, lolGameflowQueueCustomGameSpectatorPolicyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol gameflow queue custom game spectator policy
func (m LolGameflowQueueCustomGameSpectatorPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolGameflowQueueCustomGameSpectatorPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
