// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyGameModeDto lol lobby game mode dto
// swagger:model LolLobbyGameModeDto
type LolLobbyGameModeDto struct {

	// bot difficulty
	BotDifficulty string `json:"botDifficulty,omitempty"`

	// game customization
	GameCustomization map[string]string `json:"gameCustomization,omitempty"`

	// game type
	GameType string `json:"gameType,omitempty"`

	// max party size
	MaxPartySize int32 `json:"maxPartySize,omitempty"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`
}

// Validate validates this lol lobby game mode dto
func (m *LolLobbyGameModeDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyGameModeDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyGameModeDto) UnmarshalBinary(b []byte) error {
	var res LolLobbyGameModeDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
