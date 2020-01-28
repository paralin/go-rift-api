// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LootLcdsLootDescriptionDTO loot lcds loot description d t o
// swagger:model LootLcdsLootDescriptionDTO
type LootLcdsLootDescriptionDTO struct {

	// child loot table names
	ChildLootTableNames []string `json:"childLootTableNames"`

	// localization long description map
	LocalizationLongDescriptionMap map[string]string `json:"localizationLongDescriptionMap,omitempty"`

	// localization map
	LocalizationMap map[string]string `json:"localizationMap,omitempty"`

	// loot name
	LootName string `json:"lootName,omitempty"`
}

// Validate validates this loot lcds loot description d t o
func (m *LootLcdsLootDescriptionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LootLcdsLootDescriptionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LootLcdsLootDescriptionDTO) UnmarshalBinary(b []byte) error {
	var res LootLcdsLootDescriptionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}