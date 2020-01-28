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

// LolLootRedeemableStatus lol loot redeemable status
// swagger:model LolLootRedeemableStatus
type LolLootRedeemableStatus string

const (

	// LolLootRedeemableStatusUNKNOWN captures enum value "UNKNOWN"
	LolLootRedeemableStatusUNKNOWN LolLootRedeemableStatus = "UNKNOWN"

	// LolLootRedeemableStatusREDEEMABLE captures enum value "REDEEMABLE"
	LolLootRedeemableStatusREDEEMABLE LolLootRedeemableStatus = "REDEEMABLE"

	// LolLootRedeemableStatusREDEEMABLERENTAL captures enum value "REDEEMABLE_RENTAL"
	LolLootRedeemableStatusREDEEMABLERENTAL LolLootRedeemableStatus = "REDEEMABLE_RENTAL"

	// LolLootRedeemableStatusNOTREDEEMABLE captures enum value "NOT_REDEEMABLE"
	LolLootRedeemableStatusNOTREDEEMABLE LolLootRedeemableStatus = "NOT_REDEEMABLE"

	// LolLootRedeemableStatusNOTREDEEMABLERENTAL captures enum value "NOT_REDEEMABLE_RENTAL"
	LolLootRedeemableStatusNOTREDEEMABLERENTAL LolLootRedeemableStatus = "NOT_REDEEMABLE_RENTAL"

	// LolLootRedeemableStatusALREADYOWNED captures enum value "ALREADY_OWNED"
	LolLootRedeemableStatusALREADYOWNED LolLootRedeemableStatus = "ALREADY_OWNED"

	// LolLootRedeemableStatusALREADYRENTED captures enum value "ALREADY_RENTED"
	LolLootRedeemableStatusALREADYRENTED LolLootRedeemableStatus = "ALREADY_RENTED"

	// LolLootRedeemableStatusCHAMPIONNOTOWNED captures enum value "CHAMPION_NOT_OWNED"
	LolLootRedeemableStatusCHAMPIONNOTOWNED LolLootRedeemableStatus = "CHAMPION_NOT_OWNED"

	// LolLootRedeemableStatusSKINNOTOWNED captures enum value "SKIN_NOT_OWNED"
	LolLootRedeemableStatusSKINNOTOWNED LolLootRedeemableStatus = "SKIN_NOT_OWNED"
)

// for schema
var lolLootRedeemableStatusEnum []interface{}

func init() {
	var res []LolLootRedeemableStatus
	if err := json.Unmarshal([]byte(`["UNKNOWN","REDEEMABLE","REDEEMABLE_RENTAL","NOT_REDEEMABLE","NOT_REDEEMABLE_RENTAL","ALREADY_OWNED","ALREADY_RENTED","CHAMPION_NOT_OWNED","SKIN_NOT_OWNED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLootRedeemableStatusEnum = append(lolLootRedeemableStatusEnum, v)
	}
}

func (m LolLootRedeemableStatus) validateLolLootRedeemableStatusEnum(path, location string, value LolLootRedeemableStatus) error {
	if err := validate.Enum(path, location, value, lolLootRedeemableStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol loot redeemable status
func (m LolLootRedeemableStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLootRedeemableStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}