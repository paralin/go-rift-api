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

// LolLobbyLobbyPartyRewardType lol lobby lobby party reward type
// swagger:model LolLobbyLobbyPartyRewardType
type LolLobbyLobbyPartyRewardType string

const (

	// LolLobbyLobbyPartyRewardTypeIP captures enum value "Ip"
	LolLobbyLobbyPartyRewardTypeIP LolLobbyLobbyPartyRewardType = "Ip"

	// LolLobbyLobbyPartyRewardTypeIcon captures enum value "Icon"
	LolLobbyLobbyPartyRewardTypeIcon LolLobbyLobbyPartyRewardType = "Icon"

	// LolLobbyLobbyPartyRewardTypeNone captures enum value "None"
	LolLobbyLobbyPartyRewardTypeNone LolLobbyLobbyPartyRewardType = "None"
)

// for schema
var lolLobbyLobbyPartyRewardTypeEnum []interface{}

func init() {
	var res []LolLobbyLobbyPartyRewardType
	if err := json.Unmarshal([]byte(`["Ip","Icon","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLobbyLobbyPartyRewardTypeEnum = append(lolLobbyLobbyPartyRewardTypeEnum, v)
	}
}

func (m LolLobbyLobbyPartyRewardType) validateLolLobbyLobbyPartyRewardTypeEnum(path, location string, value LolLobbyLobbyPartyRewardType) error {
	if err := validate.Enum(path, location, value, lolLobbyLobbyPartyRewardTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol lobby lobby party reward type
func (m LolLobbyLobbyPartyRewardType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLobbyLobbyPartyRewardTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
