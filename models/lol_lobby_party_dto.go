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

// LolLobbyPartyDto lol lobby party dto
// swagger:model LolLobbyPartyDto
type LolLobbyPartyDto struct {

	// active restrictions
	ActiveRestrictions *LolLobbyQueueRestrictionDto `json:"activeRestrictions,omitempty"`

	// activity locked
	ActivityLocked bool `json:"activityLocked,omitempty"`

	// activity resume utc millis
	ActivityResumeUtcMillis int64 `json:"activityResumeUtcMillis,omitempty"`

	// activity started utc millis
	ActivityStartedUtcMillis int64 `json:"activityStartedUtcMillis,omitempty"`

	// chat
	Chat *LolLobbyPartyChatDto `json:"chat,omitempty"`

	// eligibility hash
	EligibilityHash int64 `json:"eligibilityHash,omitempty"`

	// eligibility restrictions
	EligibilityRestrictions []*LolLobbyGatekeeperRestrictionDto `json:"eligibilityRestrictions"`

	// game mode
	GameMode *LolLobbyGameModeDto `json:"gameMode,omitempty"`

	// max party size
	MaxPartySize int32 `json:"maxPartySize,omitempty"`

	// party Id
	PartyID string `json:"partyId,omitempty"`

	// party type
	PartyType string `json:"partyType,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// players
	Players []*LolLobbyPartyMemberDto `json:"players"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this lol lobby party dto
func (m *LolLobbyPartyDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEligibilityRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyPartyDto) validateActiveRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveRestrictions) { // not required
		return nil
	}

	if m.ActiveRestrictions != nil {
		if err := m.ActiveRestrictions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeRestrictions")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyPartyDto) validateChat(formats strfmt.Registry) error {

	if swag.IsZero(m.Chat) { // not required
		return nil
	}

	if m.Chat != nil {
		if err := m.Chat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chat")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyPartyDto) validateEligibilityRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.EligibilityRestrictions) { // not required
		return nil
	}

	for i := 0; i < len(m.EligibilityRestrictions); i++ {
		if swag.IsZero(m.EligibilityRestrictions[i]) { // not required
			continue
		}

		if m.EligibilityRestrictions[i] != nil {
			if err := m.EligibilityRestrictions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eligibilityRestrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLobbyPartyDto) validateGameMode(formats strfmt.Registry) error {

	if swag.IsZero(m.GameMode) { // not required
		return nil
	}

	if m.GameMode != nil {
		if err := m.GameMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameMode")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyPartyDto) validatePlayers(formats strfmt.Registry) error {

	if swag.IsZero(m.Players) { // not required
		return nil
	}

	for i := 0; i < len(m.Players); i++ {
		if swag.IsZero(m.Players[i]) { // not required
			continue
		}

		if m.Players[i] != nil {
			if err := m.Players[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("players" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyPartyDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyPartyDto) UnmarshalBinary(b []byte) error {
	var res LolLobbyPartyDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}