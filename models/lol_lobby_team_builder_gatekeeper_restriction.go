// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderGatekeeperRestriction lol lobby team builder gatekeeper restriction
// swagger:model LolLobbyTeamBuilderGatekeeperRestriction
type LolLobbyTeamBuilderGatekeeperRestriction struct {

	// payload
	Payload string `json:"payload,omitempty"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// remaining millis
	RemainingMillis int32 `json:"remainingMillis,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol lobby team builder gatekeeper restriction
func (m *LolLobbyTeamBuilderGatekeeperRestriction) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderGatekeeperRestriction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderGatekeeperRestriction) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderGatekeeperRestriction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
