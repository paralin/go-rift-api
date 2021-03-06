// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderTradeV1 lol lobby team builder trade v1
// swagger:model LolLobbyTeamBuilderTradeV1
type LolLobbyTeamBuilderTradeV1 struct {

	// cell Id
	CellID int32 `json:"cellId,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this lol lobby team builder trade v1
func (m *LolLobbyTeamBuilderTradeV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTradeV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderTradeV1) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderTradeV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
