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

// PendingRosterMemberState pending roster member state
// swagger:model PendingRosterMemberState
type PendingRosterMemberState string

const (

	// PendingRosterMemberStateNOTREADY captures enum value "NOT_READY"
	PendingRosterMemberStateNOTREADY PendingRosterMemberState = "NOT_READY"

	// PendingRosterMemberStateFORCEDNOTREADY captures enum value "FORCED_NOT_READY"
	PendingRosterMemberStateFORCEDNOTREADY PendingRosterMemberState = "FORCED_NOT_READY"

	// PendingRosterMemberStateREADY captures enum value "READY"
	PendingRosterMemberStateREADY PendingRosterMemberState = "READY"

	// PendingRosterMemberStateLEFT captures enum value "LEFT"
	PendingRosterMemberStateLEFT PendingRosterMemberState = "LEFT"

	// PendingRosterMemberStateKICK captures enum value "KICK"
	PendingRosterMemberStateKICK PendingRosterMemberState = "KICK"
)

// for schema
var pendingRosterMemberStateEnum []interface{}

func init() {
	var res []PendingRosterMemberState
	if err := json.Unmarshal([]byte(`["NOT_READY","FORCED_NOT_READY","READY","LEFT","KICK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pendingRosterMemberStateEnum = append(pendingRosterMemberStateEnum, v)
	}
}

func (m PendingRosterMemberState) validatePendingRosterMemberStateEnum(path, location string, value PendingRosterMemberState) error {
	if err := validate.Enum(path, location, value, pendingRosterMemberStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this pending roster member state
func (m PendingRosterMemberState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePendingRosterMemberStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}