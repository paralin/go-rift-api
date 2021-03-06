// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLootPlayerLootNotification lol loot player loot notification
// swagger:model LolLootPlayerLootNotification
type LolLootPlayerLootNotification struct {

	// acknowledged
	Acknowledged bool `json:"acknowledged,omitempty"`

	// count
	Count int32 `json:"count,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this lol loot player loot notification
func (m *LolLootPlayerLootNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLootPlayerLootNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLootPlayerLootNotification) UnmarshalBinary(b []byte) error {
	var res LolLootPlayerLootNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
