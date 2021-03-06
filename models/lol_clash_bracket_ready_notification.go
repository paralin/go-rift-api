// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClashBracketReadyNotification lol clash bracket ready notification
// swagger:model LolClashBracketReadyNotification
type LolClashBracketReadyNotification struct {

	// bracket Id
	BracketID int64 `json:"bracketId,omitempty"`

	// tournament Id
	TournamentID int64 `json:"tournamentId,omitempty"`
}

// Validate validates this lol clash bracket ready notification
func (m *LolClashBracketReadyNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClashBracketReadyNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashBracketReadyNotification) UnmarshalBinary(b []byte) error {
	var res LolClashBracketReadyNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
