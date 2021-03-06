// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LootLcdsRecipeOutputDTO loot lcds recipe output d t o
// swagger:model LootLcdsRecipeOutputDTO
type LootLcdsRecipeOutputDTO struct {

	// allow duplicates
	AllowDuplicates bool `json:"allowDuplicates,omitempty"`

	// loot name
	LootName string `json:"lootName,omitempty"`

	// probability
	Probability float64 `json:"probability,omitempty"`

	// quantity expression
	QuantityExpression string `json:"quantityExpression,omitempty"`
}

// Validate validates this loot lcds recipe output d t o
func (m *LootLcdsRecipeOutputDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LootLcdsRecipeOutputDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LootLcdsRecipeOutputDTO) UnmarshalBinary(b []byte) error {
	var res LootLcdsRecipeOutputDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
