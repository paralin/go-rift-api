// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PluginRegionLocaleChangedEvent plugin region locale changed event
// swagger:model PluginRegionLocaleChangedEvent
type PluginRegionLocaleChangedEvent struct {

	// locale
	Locale string `json:"locale,omitempty"`

	// region
	Region string `json:"region,omitempty"`
}

// Validate validates this plugin region locale changed event
func (m *PluginRegionLocaleChangedEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginRegionLocaleChangedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginRegionLocaleChangedEvent) UnmarshalBinary(b []byte) error {
	var res PluginRegionLocaleChangedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
