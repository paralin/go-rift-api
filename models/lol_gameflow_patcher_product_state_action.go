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

// LolGameflowPatcherProductStateAction lol gameflow patcher product state action
// swagger:model LolGameflowPatcherProductStateAction
type LolGameflowPatcherProductStateAction string

const (

	// LolGameflowPatcherProductStateActionIdle captures enum value "Idle"
	LolGameflowPatcherProductStateActionIdle LolGameflowPatcherProductStateAction = "Idle"

	// LolGameflowPatcherProductStateActionCheckingForUpdates captures enum value "CheckingForUpdates"
	LolGameflowPatcherProductStateActionCheckingForUpdates LolGameflowPatcherProductStateAction = "CheckingForUpdates"

	// LolGameflowPatcherProductStateActionPatching captures enum value "Patching"
	LolGameflowPatcherProductStateActionPatching LolGameflowPatcherProductStateAction = "Patching"

	// LolGameflowPatcherProductStateActionRepairing captures enum value "Repairing"
	LolGameflowPatcherProductStateActionRepairing LolGameflowPatcherProductStateAction = "Repairing"

	// LolGameflowPatcherProductStateActionMigrating captures enum value "Migrating"
	LolGameflowPatcherProductStateActionMigrating LolGameflowPatcherProductStateAction = "Migrating"
)

// for schema
var lolGameflowPatcherProductStateActionEnum []interface{}

func init() {
	var res []LolGameflowPatcherProductStateAction
	if err := json.Unmarshal([]byte(`["Idle","CheckingForUpdates","Patching","Repairing","Migrating"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolGameflowPatcherProductStateActionEnum = append(lolGameflowPatcherProductStateActionEnum, v)
	}
}

func (m LolGameflowPatcherProductStateAction) validateLolGameflowPatcherProductStateActionEnum(path, location string, value LolGameflowPatcherProductStateAction) error {
	if err := validate.Enum(path, location, value, lolGameflowPatcherProductStateActionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol gameflow patcher product state action
func (m LolGameflowPatcherProductStateAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolGameflowPatcherProductStateActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
