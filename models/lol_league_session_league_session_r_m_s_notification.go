// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLeagueSessionLeagueSessionRMSNotification lol league session league session r m s notification
// swagger:model LolLeagueSessionLeagueSessionRMSNotification
type LolLeagueSessionLeagueSessionRMSNotification struct {

	// puuid
	Puuid string `json:"puuid,omitempty"`

	// session Id
	SessionID string `json:"sessionId,omitempty"`

	// session initiated at
	SessionInitiatedAt int64 `json:"sessionInitiatedAt,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this lol league session league session r m s notification
func (m *LolLeagueSessionLeagueSessionRMSNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLeagueSessionLeagueSessionRMSNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLeagueSessionLeagueSessionRMSNotification) UnmarshalBinary(b []byte) error {
	var res LolLeagueSessionLeagueSessionRMSNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}