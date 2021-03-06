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

// LolLobbyLobbyDto lol lobby lobby dto
// swagger:model LolLobbyLobbyDto
type LolLobbyLobbyDto struct {

	// can start activity
	CanStartActivity bool `json:"canStartActivity,omitempty"`

	// chat room Id
	ChatRoomID string `json:"chatRoomId,omitempty"`

	// chat room key
	ChatRoomKey string `json:"chatRoomKey,omitempty"`

	// game config
	GameConfig *LolLobbyLobbyGameConfigDto `json:"gameConfig,omitempty"`

	// invitations
	Invitations []*LolLobbyLobbyInvitationDto `json:"invitations"`

	// local member
	LocalMember *LolLobbyLobbyParticipantDto `json:"localMember,omitempty"`

	// members
	Members []*LolLobbyLobbyParticipantDto `json:"members"`

	// party Id
	PartyID string `json:"partyId,omitempty"`

	// party type
	PartyType string `json:"partyType,omitempty"`

	// restrictions
	Restrictions []*LolLobbyEligibilityRestriction `json:"restrictions"`
}

// Validate validates this lol lobby lobby dto
func (m *LolLobbyLobbyDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvitations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyLobbyDto) validateGameConfig(formats strfmt.Registry) error {

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

func (m *LolLobbyLobbyDto) validateInvitations(formats strfmt.Registry) error {

	if swag.IsZero(m.Invitations) { // not required
		return nil
	}

	for i := 0; i < len(m.Invitations); i++ {
		if swag.IsZero(m.Invitations[i]) { // not required
			continue
		}

		if m.Invitations[i] != nil {
			if err := m.Invitations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invitations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLobbyLobbyDto) validateLocalMember(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalMember) { // not required
		return nil
	}

	if m.LocalMember != nil {
		if err := m.LocalMember.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localMember")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyLobbyDto) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLobbyLobbyDto) validateRestrictions(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolLobbyLobbyDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyLobbyDto) UnmarshalBinary(b []byte) error {
	var res LolLobbyLobbyDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
