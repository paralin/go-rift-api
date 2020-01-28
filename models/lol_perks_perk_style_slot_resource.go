// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPerksPerkStyleSlotResource lol perks perk style slot resource
// swagger:model LolPerksPerkStyleSlotResource
type LolPerksPerkStyleSlotResource struct {

	// perks
	Perks []int32 `json:"perks"`

	// slot label
	SlotLabel string `json:"slotLabel,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol perks perk style slot resource
func (m *LolPerksPerkStyleSlotResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksPerkStyleSlotResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksPerkStyleSlotResource) UnmarshalBinary(b []byte) error {
	var res LolPerksPerkStyleSlotResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
