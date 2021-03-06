// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolServiceStatusLoginDataPacket lol service status login data packet
// swagger:model LolServiceStatusLoginDataPacket
type LolServiceStatusLoginDataPacket struct {

	// broadcast notification
	BroadcastNotification *LolServiceStatusBroadcastNotification `json:"broadcastNotification,omitempty"`
}

// Validate validates this lol service status login data packet
func (m *LolServiceStatusLoginDataPacket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBroadcastNotification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolServiceStatusLoginDataPacket) validateBroadcastNotification(formats strfmt.Registry) error {

	if swag.IsZero(m.BroadcastNotification) { // not required
		return nil
	}

	if m.BroadcastNotification != nil {
		if err := m.BroadcastNotification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("broadcastNotification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolServiceStatusLoginDataPacket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolServiceStatusLoginDataPacket) UnmarshalBinary(b []byte) error {
	var res LolServiceStatusLoginDataPacket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
