// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolInventoryCurrencyDTO lol inventory currency d t o
// swagger:model LolInventoryCurrencyDTO
type LolInventoryCurrencyDTO struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// sub currencies
	SubCurrencies map[string]int32 `json:"subCurrencies,omitempty"`
}

// Validate validates this lol inventory currency d t o
func (m *LolInventoryCurrencyDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolInventoryCurrencyDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolInventoryCurrencyDTO) UnmarshalBinary(b []byte) error {
	var res LolInventoryCurrencyDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}