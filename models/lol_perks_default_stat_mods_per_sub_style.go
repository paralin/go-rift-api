// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPerksDefaultStatModsPerSubStyle lol perks default stat mods per sub style
// swagger:model LolPerksDefaultStatModsPerSubStyle
type LolPerksDefaultStatModsPerSubStyle struct {

	// id
	ID int32 `json:"id,omitempty"`

	// perks
	Perks []int32 `json:"perks"`
}

// Validate validates this lol perks default stat mods per sub style
func (m *LolPerksDefaultStatModsPerSubStyle) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksDefaultStatModsPerSubStyle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksDefaultStatModsPerSubStyle) UnmarshalBinary(b []byte) error {
	var res LolPerksDefaultStatModsPerSubStyle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
