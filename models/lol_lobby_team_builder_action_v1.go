// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderActionV1 lol lobby team builder action v1
// swagger:model LolLobbyTeamBuilderActionV1
type LolLobbyTeamBuilderActionV1 struct {

	// action Id
	ActionID int32 `json:"actionId,omitempty"`

	// actor cell Id
	ActorCellID int32 `json:"actorCellId,omitempty"`

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// completed
	Completed bool `json:"completed,omitempty"`

	// duration
	Duration int64 `json:"duration,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol lobby team builder action v1
func (m *LolLobbyTeamBuilderActionV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderActionV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderActionV1) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderActionV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
