// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashSetTicketRequest lol clash set ticket request
// swagger:model LolClashSetTicketRequest
type LolClashSetTicketRequest struct {

	// ticket amount
	TicketAmount int32 `json:"ticketAmount,omitempty"`

	// ticket type
	TicketType TicketType `json:"ticketType,omitempty"`
}

// Validate validates this lol clash set ticket request
func (m *LolClashSetTicketRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTicketType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashSetTicketRequest) validateTicketType(formats strfmt.Registry) error {

	if swag.IsZero(m.TicketType) { // not required
		return nil
	}

	if err := m.TicketType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ticketType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashSetTicketRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashSetTicketRequest) UnmarshalBinary(b []byte) error {
	var res LolClashSetTicketRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
