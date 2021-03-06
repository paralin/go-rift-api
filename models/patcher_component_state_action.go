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

// PatcherComponentStateAction patcher component state action
// swagger:model PatcherComponentStateAction
type PatcherComponentStateAction string

const (

	// PatcherComponentStateActionIdle captures enum value "Idle"
	PatcherComponentStateActionIdle PatcherComponentStateAction = "Idle"

	// PatcherComponentStateActionCheckingForUpdates captures enum value "CheckingForUpdates"
	PatcherComponentStateActionCheckingForUpdates PatcherComponentStateAction = "CheckingForUpdates"

	// PatcherComponentStateActionPatching captures enum value "Patching"
	PatcherComponentStateActionPatching PatcherComponentStateAction = "Patching"

	// PatcherComponentStateActionRepairing captures enum value "Repairing"
	PatcherComponentStateActionRepairing PatcherComponentStateAction = "Repairing"

	// PatcherComponentStateActionMigrating captures enum value "Migrating"
	PatcherComponentStateActionMigrating PatcherComponentStateAction = "Migrating"
)

// for schema
var patcherComponentStateActionEnum []interface{}

func init() {
	var res []PatcherComponentStateAction
	if err := json.Unmarshal([]byte(`["Idle","CheckingForUpdates","Patching","Repairing","Migrating"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patcherComponentStateActionEnum = append(patcherComponentStateActionEnum, v)
	}
}

func (m PatcherComponentStateAction) validatePatcherComponentStateActionEnum(path, location string, value PatcherComponentStateAction) error {
	if err := validate.Enum(path, location, value, patcherComponentStateActionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this patcher component state action
func (m PatcherComponentStateAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePatcherComponentStateActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
