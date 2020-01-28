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

// PatcherComponentStateWorkType patcher component state work type
// swagger:model PatcherComponentStateWorkType
type PatcherComponentStateWorkType string

const (

	// PatcherComponentStateWorkTypeScanning captures enum value "Scanning"
	PatcherComponentStateWorkTypeScanning PatcherComponentStateWorkType = "Scanning"

	// PatcherComponentStateWorkTypeNetwork captures enum value "Network"
	PatcherComponentStateWorkTypeNetwork PatcherComponentStateWorkType = "Network"

	// PatcherComponentStateWorkTypeDisk captures enum value "Disk"
	PatcherComponentStateWorkTypeDisk PatcherComponentStateWorkType = "Disk"
)

// for schema
var patcherComponentStateWorkTypeEnum []interface{}

func init() {
	var res []PatcherComponentStateWorkType
	if err := json.Unmarshal([]byte(`["Scanning","Network","Disk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patcherComponentStateWorkTypeEnum = append(patcherComponentStateWorkTypeEnum, v)
	}
}

func (m PatcherComponentStateWorkType) validatePatcherComponentStateWorkTypeEnum(path, location string, value PatcherComponentStateWorkType) error {
	if err := validate.Enum(path, location, value, patcherComponentStateWorkTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this patcher component state work type
func (m PatcherComponentStateWorkType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePatcherComponentStateWorkTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}