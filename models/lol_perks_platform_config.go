// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPerksPlatformConfig lol perks platform config
// swagger:model LolPerksPlatformConfig
type LolPerksPlatformConfig struct {

	// auto repair pages enabled
	AutoRepairPagesEnabled bool `json:"AutoRepairPagesEnabled,omitempty"`

	// perks enabled
	PerksEnabled bool `json:"PerksEnabled,omitempty"`
}

// Validate validates this lol perks platform config
func (m *LolPerksPlatformConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksPlatformConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksPlatformConfig) UnmarshalBinary(b []byte) error {
	var res LolPerksPlatformConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
