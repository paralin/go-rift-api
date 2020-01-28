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

// LcdsInvitationState lcds invitation state
// swagger:model LcdsInvitationState
type LcdsInvitationState string

const (

	// LcdsInvitationStateACTIVE captures enum value "ACTIVE"
	LcdsInvitationStateACTIVE LcdsInvitationState = "ACTIVE"

	// LcdsInvitationStateONHOLD captures enum value "ON_HOLD"
	LcdsInvitationStateONHOLD LcdsInvitationState = "ON_HOLD"

	// LcdsInvitationStateREVOKED captures enum value "REVOKED"
	LcdsInvitationStateREVOKED LcdsInvitationState = "REVOKED"
)

// for schema
var lcdsInvitationStateEnum []interface{}

func init() {
	var res []LcdsInvitationState
	if err := json.Unmarshal([]byte(`["ACTIVE","ON_HOLD","REVOKED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lcdsInvitationStateEnum = append(lcdsInvitationStateEnum, v)
	}
}

func (m LcdsInvitationState) validateLcdsInvitationStateEnum(path, location string, value LcdsInvitationState) error {
	if err := validate.Enum(path, location, value, lcdsInvitationStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lcds invitation state
func (m LcdsInvitationState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLcdsInvitationStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
