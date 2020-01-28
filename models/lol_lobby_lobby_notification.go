// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyLobbyNotification lol lobby lobby notification
// swagger:model LolLobbyLobbyNotification
type LolLobbyLobbyNotification struct {

	// notification Id
	NotificationID string `json:"notificationId,omitempty"`

	// notification reason
	NotificationReason string `json:"notificationReason,omitempty"`

	// summoner ids
	SummonerIds []int64 `json:"summonerIds"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this lol lobby lobby notification
func (m *LolLobbyLobbyNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyLobbyNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyLobbyNotification) UnmarshalBinary(b []byte) error {
	var res LolLobbyLobbyNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
