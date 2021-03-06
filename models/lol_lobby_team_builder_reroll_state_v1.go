// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderRerollStateV1 lol lobby team builder reroll state v1
// swagger:model LolLobbyTeamBuilderRerollStateV1
type LolLobbyTeamBuilderRerollStateV1 struct {

	// allow rerolling
	AllowRerolling bool `json:"allowRerolling,omitempty"`

	// rerolls remaining
	RerollsRemaining int32 `json:"rerollsRemaining,omitempty"`
}

// Validate validates this lol lobby team builder reroll state v1
func (m *LolLobbyTeamBuilderRerollStateV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderRerollStateV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderRerollStateV1) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderRerollStateV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
