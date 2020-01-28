// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStoreBundledItemCost lol store bundled item cost
// swagger:model LolStoreBundledItemCost
type LolStoreBundledItemCost struct {

	// cost
	Cost int64 `json:"cost,omitempty"`

	// cost type
	CostType string `json:"costType,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// discount
	Discount float32 `json:"discount,omitempty"`
}

// Validate validates this lol store bundled item cost
func (m *LolStoreBundledItemCost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreBundledItemCost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreBundledItemCost) UnmarshalBinary(b []byte) error {
	var res LolStoreBundledItemCost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
