// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// LolLootLootRarity lol loot loot rarity
// swagger:model LolLootLootRarity
type LolLootLootRarity string

const (

	// LolLootLootRarityDefault captures enum value "Default"
	LolLootLootRarityDefault LolLootLootRarity = "Default"

	// LolLootLootRarityEpic captures enum value "Epic"
	LolLootLootRarityEpic LolLootLootRarity = "Epic"

	// LolLootLootRarityLegendary captures enum value "Legendary"
	LolLootLootRarityLegendary LolLootLootRarity = "Legendary"

	// LolLootLootRarityMythic captures enum value "Mythic"
	LolLootLootRarityMythic LolLootLootRarity = "Mythic"

	// LolLootLootRarityUltimate captures enum value "Ultimate"
	LolLootLootRarityUltimate LolLootLootRarity = "Ultimate"
)

// for schema
var lolLootLootRarityEnum []interface{}

func init() {
	var res []LolLootLootRarity
	if err := json.Unmarshal([]byte(`["Default","Epic","Legendary","Mythic","Ultimate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLootLootRarityEnum = append(lolLootLootRarityEnum, v)
	}
}

func (m LolLootLootRarity) validateLolLootLootRarityEnum(path, location string, value LolLootLootRarity) error {
	if err := validate.Enum(path, location, value, lolLootLootRarityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol loot loot rarity
func (m LolLootLootRarity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLootLootRarityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
