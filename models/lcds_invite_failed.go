// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LcdsInviteFailed lcds invite failed
// swagger:model LcdsInviteFailed
type LcdsInviteFailed struct {

	// exception
	Exception *LcdsGameInviteBaseRuntimeException `json:"exception,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`
}

// Validate validates this lcds invite failed
func (m *LcdsInviteFailed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateException(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LcdsInviteFailed) validateException(formats strfmt.Registry) error {

	if swag.IsZero(m.Exception) { // not required
		return nil
	}

	if m.Exception != nil {
		if err := m.Exception.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exception")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LcdsInviteFailed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LcdsInviteFailed) UnmarshalBinary(b []byte) error {
	var res LcdsInviteFailed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
