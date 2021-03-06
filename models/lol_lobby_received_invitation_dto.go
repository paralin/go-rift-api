// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLobbyReceivedInvitationDto lol lobby received invitation dto
// swagger:model LolLobbyReceivedInvitationDto
type LolLobbyReceivedInvitationDto struct {

	// can accept invitation
	CanAcceptInvitation bool `json:"canAcceptInvitation,omitempty"`

	// from summoner Id
	FromSummonerID int64 `json:"fromSummonerId,omitempty"`

	// from summoner name
	FromSummonerName string `json:"fromSummonerName,omitempty"`

	// game config
	GameConfig *LolLobbyReceivedInvitationGameConfigDto `json:"gameConfig,omitempty"`

	// invitation Id
	InvitationID string `json:"invitationId,omitempty"`

	// restrictions
	Restrictions []*LolLobbyEligibilityRestriction `json:"restrictions"`

	// state
	State LolLobbyLobbyInvitationState `json:"state,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this lol lobby received invitation dto
func (m *LolLobbyReceivedInvitationDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyReceivedInvitationDto) validateGameConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.GameConfig) { // not required
		return nil
	}

	if m.GameConfig != nil {
		if err := m.GameConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameConfig")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyReceivedInvitationDto) validateRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.Restrictions) { // not required
		return nil
	}

	for i := 0; i < len(m.Restrictions); i++ {
		if swag.IsZero(m.Restrictions[i]) { // not required
			continue
		}

		if m.Restrictions[i] != nil {
			if err := m.Restrictions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLobbyReceivedInvitationDto) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyReceivedInvitationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyReceivedInvitationDto) UnmarshalBinary(b []byte) error {
	var res LolLobbyReceivedInvitationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
