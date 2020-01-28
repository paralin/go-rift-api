// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PendingRosterMemberDTO pending roster member d t o
// swagger:model PendingRosterMemberDTO
type PendingRosterMemberDTO struct {

	// bet
	Bet int32 `json:"bet,omitempty"`

	// bet type
	BetType TicketType `json:"betType,omitempty"`

	// join time
	JoinTime int64 `json:"joinTime,omitempty"`

	// member state
	MemberState PendingRosterMemberState `json:"memberState,omitempty"`

	// pending pay
	PendingPay int32 `json:"pendingPay,omitempty"`

	// pending premium pay
	PendingPremiumPay int32 `json:"pendingPremiumPay,omitempty"`

	// player Id
	PlayerID int64 `json:"playerId,omitempty"`

	// position
	Position Position `json:"position,omitempty"`

	// self bet
	SelfBet int32 `json:"selfBet,omitempty"`

	// tier
	Tier int32 `json:"tier,omitempty"`
}

// Validate validates this pending roster member d t o
func (m *PendingRosterMemberDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberState(formats); err != nil {
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

func (m *PendingRosterMemberDTO) validateBetType(formats strfmt.Registry) error {

	if swag.IsZero(m.BetType) { // not required
		return nil
	}

	if err := m.BetType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("betType")
		}
		return err
	}

	return nil
}

func (m *PendingRosterMemberDTO) validateMemberState(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberState) { // not required
		return nil
	}

	if err := m.MemberState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("memberState")
		}
		return err
	}

	return nil
}

func (m *PendingRosterMemberDTO) validatePosition(formats strfmt.Registry) error {

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
func (m *PendingRosterMemberDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PendingRosterMemberDTO) UnmarshalBinary(b []byte) error {
	var res PendingRosterMemberDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}