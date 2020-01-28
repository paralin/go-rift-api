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

// LolClashRosterMemberState lol clash roster member state
// swagger:model LolClashRosterMemberState
type LolClashRosterMemberState string

const (

	// LolClashRosterMemberStatePENDING captures enum value "PENDING"
	LolClashRosterMemberStatePENDING LolClashRosterMemberState = "PENDING"

	// LolClashRosterMemberStateNOTREADY captures enum value "NOT_READY"
	LolClashRosterMemberStateNOTREADY LolClashRosterMemberState = "NOT_READY"

	// LolClashRosterMemberStateFORCEDNOTREADY captures enum value "FORCED_NOT_READY"
	LolClashRosterMemberStateFORCEDNOTREADY LolClashRosterMemberState = "FORCED_NOT_READY"

	// LolClashRosterMemberStateREADY captures enum value "READY"
	LolClashRosterMemberStateREADY LolClashRosterMemberState = "READY"
)

// for schema
var lolClashRosterMemberStateEnum []interface{}

func init() {
	var res []LolClashRosterMemberState
	if err := json.Unmarshal([]byte(`["PENDING","NOT_READY","FORCED_NOT_READY","READY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolClashRosterMemberStateEnum = append(lolClashRosterMemberStateEnum, v)
	}
}

func (m LolClashRosterMemberState) validateLolClashRosterMemberStateEnum(path, location string, value LolClashRosterMemberState) error {
	if err := validate.Enum(path, location, value, lolClashRosterMemberStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol clash roster member state
func (m LolClashRosterMemberState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolClashRosterMemberStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}