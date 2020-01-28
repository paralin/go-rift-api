// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyPartyInviteAnalytics lol lobby party invite analytics
// swagger:model LolLobbyPartyInviteAnalytics
type LolLobbyPartyInviteAnalytics struct {

	// event data
	EventData interface{} `json:"eventData,omitempty"`

	// event timestamp
	EventTimestamp int64 `json:"eventTimestamp,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// from summoner Id
	FromSummonerID int64 `json:"fromSummonerId,omitempty"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// party Id
	PartyID string `json:"partyId,omitempty"`

	// party type
	PartyType string `json:"partyType,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// to summoner Id
	ToSummonerID int64 `json:"toSummonerId,omitempty"`
}

// Validate validates this lol lobby party invite analytics
func (m *LolLobbyPartyInviteAnalytics) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyPartyInviteAnalytics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyPartyInviteAnalytics) UnmarshalBinary(b []byte) error {
	var res LolLobbyPartyInviteAnalytics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
