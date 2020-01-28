// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashRosterWithdrawNotification lol clash roster withdraw notification
// swagger:model LolClashRosterWithdrawNotification
type LolClashRosterWithdrawNotification struct {

	// notify reason
	NotifyReason LolClashRosterNotifyReason `json:"notifyReason,omitempty"`

	// roster Id
	RosterID int64 `json:"rosterId,omitempty"`

	// source player Id
	SourcePlayerID int64 `json:"sourcePlayerId,omitempty"`

	// tournament Id
	TournamentID int64 `json:"tournamentId,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`

	// withdraw
	Withdraw *RosterWithdraw `json:"withdraw,omitempty"`
}

// Validate validates this lol clash roster withdraw notification
func (m *LolClashRosterWithdrawNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifyReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithdraw(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashRosterWithdrawNotification) validateNotifyReason(formats strfmt.Registry) error {

	if swag.IsZero(m.NotifyReason) { // not required
		return nil
	}

	if err := m.NotifyReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notifyReason")
		}
		return err
	}

	return nil
}

func (m *LolClashRosterWithdrawNotification) validateWithdraw(formats strfmt.Registry) error {

	if swag.IsZero(m.Withdraw) { // not required
		return nil
	}

	if m.Withdraw != nil {
		if err := m.Withdraw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("withdraw")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashRosterWithdrawNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashRosterWithdrawNotification) UnmarshalBinary(b []byte) error {
	var res LolClashRosterWithdrawNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}