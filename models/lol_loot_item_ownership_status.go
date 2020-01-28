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

// LolLootItemOwnershipStatus lol loot item ownership status
// swagger:model LolLootItemOwnershipStatus
type LolLootItemOwnershipStatus string

const (

	// LolLootItemOwnershipStatusNONE captures enum value "NONE"
	LolLootItemOwnershipStatusNONE LolLootItemOwnershipStatus = "NONE"

	// LolLootItemOwnershipStatusFREE captures enum value "FREE"
	LolLootItemOwnershipStatusFREE LolLootItemOwnershipStatus = "FREE"

	// LolLootItemOwnershipStatusRENTAL captures enum value "RENTAL"
	LolLootItemOwnershipStatusRENTAL LolLootItemOwnershipStatus = "RENTAL"

	// LolLootItemOwnershipStatusOWNED captures enum value "OWNED"
	LolLootItemOwnershipStatusOWNED LolLootItemOwnershipStatus = "OWNED"
)

// for schema
var lolLootItemOwnershipStatusEnum []interface{}

func init() {
	var res []LolLootItemOwnershipStatus
	if err := json.Unmarshal([]byte(`["NONE","FREE","RENTAL","OWNED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLootItemOwnershipStatusEnum = append(lolLootItemOwnershipStatusEnum, v)
	}
}

func (m LolLootItemOwnershipStatus) validateLolLootItemOwnershipStatusEnum(path, location string, value LolLootItemOwnershipStatus) error {
	if err := validate.Enum(path, location, value, lolLootItemOwnershipStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol loot item ownership status
func (m LolLootItemOwnershipStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLootItemOwnershipStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
