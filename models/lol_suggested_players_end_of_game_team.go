// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolSuggestedPlayersEndOfGameTeam lol suggested players end of game team
// swagger:model LolSuggestedPlayersEndOfGameTeam
type LolSuggestedPlayersEndOfGameTeam struct {

	// is winning team
	IsWinningTeam bool `json:"isWinningTeam,omitempty"`

	// players
	Players []*LolSuggestedPlayersEndOfGamePlayer `json:"players"`
}

// Validate validates this lol suggested players end of game team
func (m *LolSuggestedPlayersEndOfGameTeam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolSuggestedPlayersEndOfGameTeam) validatePlayers(formats strfmt.Registry) error {

	if swag.IsZero(m.Players) { // not required
		return nil
	}

	for i := 0; i < len(m.Players); i++ {
		if swag.IsZero(m.Players[i]) { // not required
			continue
		}

		if m.Players[i] != nil {
			if err := m.Players[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("players" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolSuggestedPlayersEndOfGameTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSuggestedPlayersEndOfGameTeam) UnmarshalBinary(b []byte) error {
	var res LolSuggestedPlayersEndOfGameTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
