// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLootLootItem lol loot loot item
// swagger:model LolLootLootItem
type LolLootLootItem struct {

	// asset
	Asset string `json:"asset,omitempty"`

	// display categories
	DisplayCategories string `json:"displayCategories,omitempty"`

	// expiry time
	ExpiryTime int64 `json:"expiryTime,omitempty"`

	// loot name
	LootName string `json:"lootName,omitempty"`

	// rarity
	Rarity string `json:"rarity,omitempty"`

	// rental games
	RentalGames int32 `json:"rentalGames,omitempty"`

	// rental seconds
	RentalSeconds int64 `json:"rentalSeconds,omitempty"`

	// store item Id
	StoreItemID int32 `json:"storeItemId,omitempty"`

	// tags
	Tags string `json:"tags,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// upgrade loot name
	UpgradeLootName string `json:"upgradeLootName,omitempty"`

	// value
	Value int32 `json:"value,omitempty"`
}

// Validate validates this lol loot loot item
func (m *LolLootLootItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLootLootItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLootLootItem) UnmarshalBinary(b []byte) error {
	var res LolLootLootItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}