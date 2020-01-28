// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LolLobbyPremadePartyDto lol lobby premade party dto
// swagger:model LolLobbyPremadePartyDto
type LolLobbyPremadePartyDto struct {

	// party Id
	PartyID string `json:"partyId,omitempty"`

	// players
	Players map[string]LolLobbyPremadeMemberDto `json:"players,omitempty"`
}

// Validate validates this lol lobby premade party dto
func (m *LolLobbyPremadePartyDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyPremadePartyDto) validatePlayers(formats strfmt.Registry) error {

	if swag.IsZero(m.Players) { // not required
		return nil
	}

	for k := range m.Players {

		if err := validate.Required("players"+"."+k, "body", m.Players[k]); err != nil {
			return err
		}
		if val, ok := m.Players[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyPremadePartyDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyPremadePartyDto) UnmarshalBinary(b []byte) error {
	var res LolLobbyPremadePartyDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
