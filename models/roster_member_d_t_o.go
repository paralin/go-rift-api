// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RosterMemberDTO roster member d t o
// swagger:model RosterMemberDTO
type RosterMemberDTO struct {

	// bid type
	BidType TicketType `json:"bidType,omitempty"`

	// current bid
	CurrentBid int32 `json:"currentBid,omitempty"`

	// join time
	JoinTime int64 `json:"joinTime,omitempty"`

	// pending premium spend
	PendingPremiumSpend int32 `json:"pendingPremiumSpend,omitempty"`

	// pending spend
	PendingSpend int32 `json:"pendingSpend,omitempty"`

	// player Id
	PlayerID int64 `json:"playerId,omitempty"`

	// position
	Position Position `json:"position,omitempty"`

	// roster Id
	RosterID int64 `json:"rosterId,omitempty"`

	// tier
	Tier int32 `json:"tier,omitempty"`

	// tournament Id
	TournamentID int64 `json:"tournamentId,omitempty"`
}

// Validate validates this roster member d t o
func (m *RosterMemberDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBidType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RosterMemberDTO) validateBidType(formats strfmt.Registry) error {

	if swag.IsZero(m.BidType) { // not required
		return nil
	}

	if err := m.BidType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bidType")
		}
		return err
	}

	return nil
}

func (m *RosterMemberDTO) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := m.Position.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("position")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RosterMemberDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RosterMemberDTO) UnmarshalBinary(b []byte) error {
	var res RosterMemberDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
