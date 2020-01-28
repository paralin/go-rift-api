// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPerksPerkSettingResource lol perks perk setting resource
// swagger:model LolPerksPerkSettingResource
type LolPerksPerkSettingResource struct {

	// perk ids
	PerkIds []int32 `json:"perkIds"`

	// perk style
	PerkStyle int32 `json:"perkStyle,omitempty"`

	// perk sub style
	PerkSubStyle int32 `json:"perkSubStyle,omitempty"`
}

// Validate validates this lol perks perk setting resource
func (m *LolPerksPerkSettingResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksPerkSettingResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksPerkSettingResource) UnmarshalBinary(b []byte) error {
	var res LolPerksPerkSettingResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
