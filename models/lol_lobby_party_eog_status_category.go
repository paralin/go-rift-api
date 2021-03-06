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

// LolLobbyPartyEogStatusCategory lol lobby party eog status category
// swagger:model LolLobbyPartyEogStatusCategory
type LolLobbyPartyEogStatusCategory string

const (

	// LolLobbyPartyEogStatusCategoryKLeft captures enum value "kLeft"
	LolLobbyPartyEogStatusCategoryKLeft LolLobbyPartyEogStatusCategory = "kLeft"

	// LolLobbyPartyEogStatusCategoryKPlayAgain captures enum value "kPlayAgain"
	LolLobbyPartyEogStatusCategoryKPlayAgain LolLobbyPartyEogStatusCategory = "kPlayAgain"

	// LolLobbyPartyEogStatusCategoryKOnEog captures enum value "kOnEog"
	LolLobbyPartyEogStatusCategoryKOnEog LolLobbyPartyEogStatusCategory = "kOnEog"
)

// for schema
var lolLobbyPartyEogStatusCategoryEnum []interface{}

func init() {
	var res []LolLobbyPartyEogStatusCategory
	if err := json.Unmarshal([]byte(`["kLeft","kPlayAgain","kOnEog"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLobbyPartyEogStatusCategoryEnum = append(lolLobbyPartyEogStatusCategoryEnum, v)
	}
}

func (m LolLobbyPartyEogStatusCategory) validateLolLobbyPartyEogStatusCategoryEnum(path, location string, value LolLobbyPartyEogStatusCategory) error {
	if err := validate.Enum(path, location, value, lolLobbyPartyEogStatusCategoryEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol lobby party eog status category
func (m LolLobbyPartyEogStatusCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLobbyPartyEogStatusCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
