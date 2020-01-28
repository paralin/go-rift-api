// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PhaseInMember phase in member
// swagger:model PhaseInMember
type PhaseInMember struct {

	// bet
	Bet int32 `json:"bet,omitempty"`

	// pending pay
	PendingPay int32 `json:"pendingPay,omitempty"`

	// player Id
	PlayerID int64 `json:"playerId,omitempty"`

	// self bet
	SelfBet int32 `json:"selfBet,omitempty"`
}

// Validate validates this phase in member
func (m *PhaseInMember) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhaseInMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhaseInMember) UnmarshalBinary(b []byte) error {
	var res PhaseInMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}