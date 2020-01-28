// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PlayerLootDTO player loot d t o
// swagger:model PlayerLootDTO
type PlayerLootDTO struct {

	// count
	Count int32 `json:"count,omitempty"`

	// expiry time
	ExpiryTime int64 `json:"expiryTime,omitempty"`

	// loot name
	LootName string `json:"lootName,omitempty"`

	// ref Id
	RefID string `json:"refId,omitempty"`
}

// Validate validates this player loot d t o
func (m *PlayerLootDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlayerLootDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerLootDTO) UnmarshalBinary(b []byte) error {
	var res PlayerLootDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
