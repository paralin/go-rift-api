// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLoyaltyInventoryDTO lol loyalty inventory d t o
// swagger:model LolLoyaltyInventoryDTO
type LolLoyaltyInventoryDTO struct {

	// items
	Items map[string]interface{} `json:"items,omitempty"`
}

// Validate validates this lol loyalty inventory d t o
func (m *LolLoyaltyInventoryDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLoyaltyInventoryDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoyaltyInventoryDTO) UnmarshalBinary(b []byte) error {
	var res LolLoyaltyInventoryDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
