// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyLcdsGameMetaData lol lobby lcds game meta data
// swagger:model LolLobbyLcdsGameMetaData
type LolLobbyLcdsGameMetaData struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// map Id
	MapID int32 `json:"mapId,omitempty"`
}

// Validate validates this lol lobby lcds game meta data
func (m *LolLobbyLcdsGameMetaData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyLcdsGameMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyLcdsGameMetaData) UnmarshalBinary(b []byte) error {
	var res LolLobbyLcdsGameMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
