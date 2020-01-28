// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BracketMatch bracket match
// swagger:model BracketMatch
type BracketMatch struct {

	// fail roster status
	FailRosterStatus int32 `json:"failRosterStatus,omitempty"`

	// forfeit roster Id
	ForfeitRosterID int64 `json:"forfeitRosterId,omitempty"`

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// game start time
	GameStartTime int64 `json:"gameStartTime,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// loser bracket
	LoserBracket bool `json:"loserBracket,omitempty"`

	// order
	Order int32 `json:"order,omitempty"`

	// roster id1
	RosterId1 int64 `json:"rosterId1,omitempty"`

	// roster id2
	RosterId2 int64 `json:"rosterId2,omitempty"`

	// round
	Round int32 `json:"round,omitempty"`

	// round start time
	RoundStartTime int64 `json:"roundStartTime,omitempty"`

	// status
	Status ClientBracketMatchStatus `json:"status,omitempty"`

	// winner Id
	WinnerID int64 `json:"winnerId,omitempty"`
}

// Validate validates this bracket match
func (m *BracketMatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BracketMatch) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BracketMatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BracketMatch) UnmarshalBinary(b []byte) error {
	var res BracketMatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}