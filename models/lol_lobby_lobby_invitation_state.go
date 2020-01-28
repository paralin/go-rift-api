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

// LolLobbyLobbyInvitationState lol lobby lobby invitation state
// swagger:model LolLobbyLobbyInvitationState
type LolLobbyLobbyInvitationState string

const (

	// LolLobbyLobbyInvitationStateRequested captures enum value "Requested"
	LolLobbyLobbyInvitationStateRequested LolLobbyLobbyInvitationState = "Requested"

	// LolLobbyLobbyInvitationStatePending captures enum value "Pending"
	LolLobbyLobbyInvitationStatePending LolLobbyLobbyInvitationState = "Pending"

	// LolLobbyLobbyInvitationStateAccepted captures enum value "Accepted"
	LolLobbyLobbyInvitationStateAccepted LolLobbyLobbyInvitationState = "Accepted"

	// LolLobbyLobbyInvitationStateJoined captures enum value "Joined"
	LolLobbyLobbyInvitationStateJoined LolLobbyLobbyInvitationState = "Joined"

	// LolLobbyLobbyInvitationStateDeclined captures enum value "Declined"
	LolLobbyLobbyInvitationStateDeclined LolLobbyLobbyInvitationState = "Declined"

	// LolLobbyLobbyInvitationStateKicked captures enum value "Kicked"
	LolLobbyLobbyInvitationStateKicked LolLobbyLobbyInvitationState = "Kicked"

	// LolLobbyLobbyInvitationStateOnHold captures enum value "OnHold"
	LolLobbyLobbyInvitationStateOnHold LolLobbyLobbyInvitationState = "OnHold"

	// LolLobbyLobbyInvitationStateError captures enum value "Error"
	LolLobbyLobbyInvitationStateError LolLobbyLobbyInvitationState = "Error"
)

// for schema
var lolLobbyLobbyInvitationStateEnum []interface{}

func init() {
	var res []LolLobbyLobbyInvitationState
	if err := json.Unmarshal([]byte(`["Requested","Pending","Accepted","Joined","Declined","Kicked","OnHold","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLobbyLobbyInvitationStateEnum = append(lolLobbyLobbyInvitationStateEnum, v)
	}
}

func (m LolLobbyLobbyInvitationState) validateLolLobbyLobbyInvitationStateEnum(path, location string, value LolLobbyLobbyInvitationState) error {
	if err := validate.Enum(path, location, value, lolLobbyLobbyInvitationStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol lobby lobby invitation state
func (m LolLobbyLobbyInvitationState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLobbyLobbyInvitationStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}