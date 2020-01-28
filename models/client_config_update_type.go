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

// ClientConfigUpdateType client config update type
// swagger:model ClientConfigUpdateType
type ClientConfigUpdateType string

const (

	// ClientConfigUpdateTypeCreate captures enum value "Create"
	ClientConfigUpdateTypeCreate ClientConfigUpdateType = "Create"

	// ClientConfigUpdateTypeUpdate captures enum value "Update"
	ClientConfigUpdateTypeUpdate ClientConfigUpdateType = "Update"

	// ClientConfigUpdateTypeDelete captures enum value "Delete"
	ClientConfigUpdateTypeDelete ClientConfigUpdateType = "Delete"
)

// for schema
var clientConfigUpdateTypeEnum []interface{}

func init() {
	var res []ClientConfigUpdateType
	if err := json.Unmarshal([]byte(`["Create","Update","Delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientConfigUpdateTypeEnum = append(clientConfigUpdateTypeEnum, v)
	}
}

func (m ClientConfigUpdateType) validateClientConfigUpdateTypeEnum(path, location string, value ClientConfigUpdateType) error {
	if err := validate.Enum(path, location, value, clientConfigUpdateTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this client config update type
func (m ClientConfigUpdateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClientConfigUpdateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
