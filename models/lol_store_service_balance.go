// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStoreServiceBalance lol store service balance
// swagger:model LolStoreServiceBalance
type LolStoreServiceBalance struct {

	// amount
	Amount int64 `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

// Validate validates this lol store service balance
func (m *LolStoreServiceBalance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreServiceBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreServiceBalance) UnmarshalBinary(b []byte) error {
	var res LolStoreServiceBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
