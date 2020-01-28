// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPerksPlayerInventory lol perks player inventory
// swagger:model LolPerksPlayerInventory
type LolPerksPlayerInventory struct {

	// owned page count
	OwnedPageCount int32 `json:"ownedPageCount,omitempty"`
}

// Validate validates this lol perks player inventory
func (m *LolPerksPlayerInventory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksPlayerInventory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksPlayerInventory) UnmarshalBinary(b []byte) error {
	var res LolPerksPlayerInventory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
