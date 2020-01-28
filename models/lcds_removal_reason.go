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

// LcdsRemovalReason lcds removal reason
// swagger:model LcdsRemovalReason
type LcdsRemovalReason string

const (

	// LcdsRemovalReasonKICKED captures enum value "KICKED"
	LcdsRemovalReasonKICKED LcdsRemovalReason = "KICKED"

	// LcdsRemovalReasonDESTROYED captures enum value "DESTROYED"
	LcdsRemovalReasonDESTROYED LcdsRemovalReason = "DESTROYED"

	// LcdsRemovalReasonPROGRESSED captures enum value "PROGRESSED"
	LcdsRemovalReasonPROGRESSED LcdsRemovalReason = "PROGRESSED"
)

// for schema
var lcdsRemovalReasonEnum []interface{}

func init() {
	var res []LcdsRemovalReason
	if err := json.Unmarshal([]byte(`["KICKED","DESTROYED","PROGRESSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lcdsRemovalReasonEnum = append(lcdsRemovalReasonEnum, v)
	}
}

func (m LcdsRemovalReason) validateLcdsRemovalReasonEnum(path, location string, value LcdsRemovalReason) error {
	if err := validate.Enum(path, location, value, lcdsRemovalReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lcds removal reason
func (m LcdsRemovalReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLcdsRemovalReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
