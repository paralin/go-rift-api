// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashTournamentGameEnd lol clash tournament game end
// swagger:model LolClashTournamentGameEnd
type LolClashTournamentGameEnd struct {

	// bracket Id
	BracketID int64 `json:"bracketId,omitempty"`

	// old bracket
	OldBracket *LolClashBracket `json:"oldBracket,omitempty"`

	// tournament Id
	TournamentID int64 `json:"tournamentId,omitempty"`

	// tournament name loc key
	TournamentNameLocKey string `json:"tournamentNameLocKey,omitempty"`

	// tournament name loc key secondary
	TournamentNameLocKeySecondary string `json:"tournamentNameLocKeySecondary,omitempty"`
}

// Validate validates this lol clash tournament game end
func (m *LolClashTournamentGameEnd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOldBracket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashTournamentGameEnd) validateOldBracket(formats strfmt.Registry) error {

	if swag.IsZero(m.OldBracket) { // not required
		return nil
	}

	if m.OldBracket != nil {
		if err := m.OldBracket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oldBracket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashTournamentGameEnd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashTournamentGameEnd) UnmarshalBinary(b []byte) error {
	var res LolClashTournamentGameEnd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
