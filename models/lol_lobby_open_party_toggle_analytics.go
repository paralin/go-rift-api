// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyOpenPartyToggleAnalytics lol lobby open party toggle analytics
// swagger:model LolLobbyOpenPartyToggleAnalytics
type LolLobbyOpenPartyToggleAnalytics struct {

	// elapsed time
	ElapsedTime int64 `json:"elapsedTime,omitempty"`

	// end timestamp
	EndTimestamp int64 `json:"endTimestamp,omitempty"`

	// end transition
	EndTransition string `json:"endTransition,omitempty"`

	// event data
	EventData interface{} `json:"eventData,omitempty"`

	// event timestamp
	EventTimestamp int64 `json:"eventTimestamp,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// final state
	FinalState string `json:"finalState,omitempty"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// initial state
	InitialState string `json:"initialState,omitempty"`

	// num of toggles
	NumOfToggles int32 `json:"numOfToggles,omitempty"`

	// party Id
	PartyID string `json:"partyId,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// player Id
	PlayerID string `json:"playerId,omitempty"`

	// start timestamp
	StartTimestamp int64 `json:"startTimestamp,omitempty"`

	// start transition
	StartTransition string `json:"startTransition,omitempty"`
}

// Validate validates this lol lobby open party toggle analytics
func (m *LolLobbyOpenPartyToggleAnalytics) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyOpenPartyToggleAnalytics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyOpenPartyToggleAnalytics) UnmarshalBinary(b []byte) error {
	var res LolLobbyOpenPartyToggleAnalytics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
