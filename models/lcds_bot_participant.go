// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LcdsBotParticipant lcds bot participant
// swagger:model LcdsBotParticipant
type LcdsBotParticipant struct {

	// bot skill level
	BotSkillLevel int32 `json:"botSkillLevel,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner internal name
	SummonerInternalName string `json:"summonerInternalName,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`

	// team Id
	TeamID string `json:"teamId,omitempty"`
}

// Validate validates this lcds bot participant
func (m *LcdsBotParticipant) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LcdsBotParticipant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LcdsBotParticipant) UnmarshalBinary(b []byte) error {
	var res LcdsBotParticipant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
