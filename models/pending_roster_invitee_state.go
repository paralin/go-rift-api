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

// PendingRosterInviteeState pending roster invitee state
// swagger:model PendingRosterInviteeState
type PendingRosterInviteeState string

const (

	// PendingRosterInviteeStateSUGGESTED captures enum value "SUGGESTED"
	PendingRosterInviteeStateSUGGESTED PendingRosterInviteeState = "SUGGESTED"

	// PendingRosterInviteeStatePENDING captures enum value "PENDING"
	PendingRosterInviteeStatePENDING PendingRosterInviteeState = "PENDING"

	// PendingRosterInviteeStateDECLINED captures enum value "DECLINED"
	PendingRosterInviteeStateDECLINED PendingRosterInviteeState = "DECLINED"

	// PendingRosterInviteeStateREVOKED captures enum value "REVOKED"
	PendingRosterInviteeStateREVOKED PendingRosterInviteeState = "REVOKED"
)

// for schema
var pendingRosterInviteeStateEnum []interface{}

func init() {
	var res []PendingRosterInviteeState
	if err := json.Unmarshal([]byte(`["SUGGESTED","PENDING","DECLINED","REVOKED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pendingRosterInviteeStateEnum = append(pendingRosterInviteeStateEnum, v)
	}
}

func (m PendingRosterInviteeState) validatePendingRosterInviteeStateEnum(path, location string, value PendingRosterInviteeState) error {
	if err := validate.Enum(path, location, value, pendingRosterInviteeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this pending roster invitee state
func (m PendingRosterInviteeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePendingRosterInviteeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
