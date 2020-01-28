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

// LolPremadeVoicePartyMemberRoleEnum lol premade voice party member role enum
// swagger:model LolPremadeVoicePartyMemberRoleEnum
type LolPremadeVoicePartyMemberRoleEnum string

const (

	// LolPremadeVoicePartyMemberRoleEnumLEADER captures enum value "LEADER"
	LolPremadeVoicePartyMemberRoleEnumLEADER LolPremadeVoicePartyMemberRoleEnum = "LEADER"

	// LolPremadeVoicePartyMemberRoleEnumMEMBER captures enum value "MEMBER"
	LolPremadeVoicePartyMemberRoleEnumMEMBER LolPremadeVoicePartyMemberRoleEnum = "MEMBER"

	// LolPremadeVoicePartyMemberRoleEnumINVITED captures enum value "INVITED"
	LolPremadeVoicePartyMemberRoleEnumINVITED LolPremadeVoicePartyMemberRoleEnum = "INVITED"

	// LolPremadeVoicePartyMemberRoleEnumHOLD captures enum value "HOLD"
	LolPremadeVoicePartyMemberRoleEnumHOLD LolPremadeVoicePartyMemberRoleEnum = "HOLD"

	// LolPremadeVoicePartyMemberRoleEnumKICKED captures enum value "KICKED"
	LolPremadeVoicePartyMemberRoleEnumKICKED LolPremadeVoicePartyMemberRoleEnum = "KICKED"

	// LolPremadeVoicePartyMemberRoleEnumDECLINED captures enum value "DECLINED"
	LolPremadeVoicePartyMemberRoleEnumDECLINED LolPremadeVoicePartyMemberRoleEnum = "DECLINED"
)

// for schema
var lolPremadeVoicePartyMemberRoleEnumEnum []interface{}

func init() {
	var res []LolPremadeVoicePartyMemberRoleEnum
	if err := json.Unmarshal([]byte(`["LEADER","MEMBER","INVITED","HOLD","KICKED","DECLINED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolPremadeVoicePartyMemberRoleEnumEnum = append(lolPremadeVoicePartyMemberRoleEnumEnum, v)
	}
}

func (m LolPremadeVoicePartyMemberRoleEnum) validateLolPremadeVoicePartyMemberRoleEnumEnum(path, location string, value LolPremadeVoicePartyMemberRoleEnum) error {
	if err := validate.Enum(path, location, value, lolPremadeVoicePartyMemberRoleEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol premade voice party member role enum
func (m LolPremadeVoicePartyMemberRoleEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolPremadeVoicePartyMemberRoleEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}