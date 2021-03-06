// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectSettingsResource lol champ select settings resource
// swagger:model LolChampSelectSettingsResource
type LolChampSelectSettingsResource struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// schema version
	SchemaVersion int64 `json:"schemaVersion,omitempty"`
}

// Validate validates this lol champ select settings resource
func (m *LolChampSelectSettingsResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectSettingsResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectSettingsResource) UnmarshalBinary(b []byte) error {
	var res LolChampSelectSettingsResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
