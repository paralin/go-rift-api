// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLobbyLobbyStatus lol lobby lobby status
// swagger:model LolLobbyLobbyStatus
type LolLobbyLobbyStatus struct {

	// allowed play again
	AllowedPlayAgain bool `json:"allowedPlayAgain,omitempty"`

	// custom spectator policy
	CustomSpectatorPolicy LolLobbyQueueCustomGameSpectatorPolicy `json:"customSpectatorPolicy,omitempty"`

	// invited summoner ids
	InvitedSummonerIds []int64 `json:"invitedSummonerIds"`

	// is custom
	IsCustom bool `json:"isCustom,omitempty"`

	// is leader
	IsLeader bool `json:"isLeader,omitempty"`

	// is practice tool
	IsPracticeTool bool `json:"isPracticeTool,omitempty"`

	// is spectator
	IsSpectator bool `json:"isSpectator,omitempty"`

	// lobby Id
	LobbyID string `json:"lobbyId,omitempty"`

	// member summoner ids
	MemberSummonerIds []int64 `json:"memberSummonerIds"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`
}

// Validate validates this lol lobby lobby status
func (m *LolLobbyLobbyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomSpectatorPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyLobbyStatus) validateCustomSpectatorPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomSpectatorPolicy) { // not required
		return nil
	}

	if err := m.CustomSpectatorPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("customSpectatorPolicy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyLobbyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyLobbyStatus) UnmarshalBinary(b []byte) error {
	var res LolLobbyLobbyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
