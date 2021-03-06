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

// Role role
// swagger:model Role
type Role string

const (

	// RoleCAPTAIN captures enum value "CAPTAIN"
	RoleCAPTAIN Role = "CAPTAIN"

	// RoleMEMBER captures enum value "MEMBER"
	RoleMEMBER Role = "MEMBER"

	// RoleNONE captures enum value "NONE"
	RoleNONE Role = "NONE"
)

// for schema
var roleEnum []interface{}

func init() {
	var res []Role
	if err := json.Unmarshal([]byte(`["CAPTAIN","MEMBER","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roleEnum = append(roleEnum, v)
	}
}

func (m Role) validateRoleEnum(path, location string, value Role) error {
	if err := validate.Enum(path, location, value, roleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this role
func (m Role) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
