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

// PluginResourceEventType plugin resource event type
// swagger:model PluginResourceEventType
type PluginResourceEventType string

const (

	// PluginResourceEventTypeCreate captures enum value "Create"
	PluginResourceEventTypeCreate PluginResourceEventType = "Create"

	// PluginResourceEventTypeUpdate captures enum value "Update"
	PluginResourceEventTypeUpdate PluginResourceEventType = "Update"

	// PluginResourceEventTypeDelete captures enum value "Delete"
	PluginResourceEventTypeDelete PluginResourceEventType = "Delete"
)

// for schema
var pluginResourceEventTypeEnum []interface{}

func init() {
	var res []PluginResourceEventType
	if err := json.Unmarshal([]byte(`["Create","Update","Delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pluginResourceEventTypeEnum = append(pluginResourceEventTypeEnum, v)
	}
}

func (m PluginResourceEventType) validatePluginResourceEventTypeEnum(path, location string, value PluginResourceEventType) error {
	if err := validate.Enum(path, location, value, pluginResourceEventTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this plugin resource event type
func (m PluginResourceEventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePluginResourceEventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}