// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLootLootItemGdsResource lol loot loot item gds resource
// swagger:model LolLootLootItemGdsResource
type LolLootLootItemGdsResource struct {

	// auto redeem
	AutoRedeem bool `json:"autoRedeem,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end date
	EndDate string `json:"endDate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// lifetime max
	LifetimeMax int32 `json:"lifetimeMax,omitempty"`

	// mapped store Id
	MappedStoreID int32 `json:"mappedStoreId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rarity
	Rarity LolLootLootRarity `json:"rarity,omitempty"`

	// recipe menu active
	RecipeMenuActive bool `json:"recipeMenuActive,omitempty"`

	// recipe menu subtitle
	RecipeMenuSubtitle string `json:"recipeMenuSubtitle,omitempty"`

	// recipe menu title
	RecipeMenuTitle string `json:"recipeMenuTitle,omitempty"`

	// start date
	StartDate string `json:"startDate,omitempty"`

	// type
	Type LolLootLootType `json:"type,omitempty"`
}

// Validate validates this lol loot loot item gds resource
func (m *LolLootLootItemGdsResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRarity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLootLootItemGdsResource) validateRarity(formats strfmt.Registry) error {

	if swag.IsZero(m.Rarity) { // not required
		return nil
	}

	if err := m.Rarity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rarity")
		}
		return err
	}

	return nil
}

func (m *LolLootLootItemGdsResource) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLootLootItemGdsResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLootLootItemGdsResource) UnmarshalBinary(b []byte) error {
	var res LolLootLootItemGdsResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
