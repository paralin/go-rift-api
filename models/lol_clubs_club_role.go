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

// LolClubsClubRole lol clubs club role
// swagger:model LolClubsClubRole
type LolClubsClubRole string

const (

	// LolClubsClubRoleUNKNOWN captures enum value "UNKNOWN"
	LolClubsClubRoleUNKNOWN LolClubsClubRole = "UNKNOWN"

	// LolClubsClubRoleOWNER captures enum value "OWNER"
	LolClubsClubRoleOWNER LolClubsClubRole = "OWNER"

	// LolClubsClubRoleOFFICER captures enum value "OFFICER"
	LolClubsClubRoleOFFICER LolClubsClubRole = "OFFICER"

	// LolClubsClubRoleMEMBER captures enum value "MEMBER"
	LolClubsClubRoleMEMBER LolClubsClubRole = "MEMBER"

	// LolClubsClubRoleINVITEE captures enum value "INVITEE"
	LolClubsClubRoleINVITEE LolClubsClubRole = "INVITEE"

	// LolClubsClubRoleNOMINEE captures enum value "NOMINEE"
	LolClubsClubRoleNOMINEE LolClubsClubRole = "NOMINEE"
)

// for schema
var lolClubsClubRoleEnum []interface{}

func init() {
	var res []LolClubsClubRole
	if err := json.Unmarshal([]byte(`["UNKNOWN","OWNER","OFFICER","MEMBER","INVITEE","NOMINEE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolClubsClubRoleEnum = append(lolClubsClubRoleEnum, v)
	}
}

func (m LolClubsClubRole) validateLolClubsClubRoleEnum(path, location string, value LolClubsClubRole) error {
	if err := validate.Enum(path, location, value, lolClubsClubRoleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol clubs club role
func (m LolClubsClubRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolClubsClubRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}