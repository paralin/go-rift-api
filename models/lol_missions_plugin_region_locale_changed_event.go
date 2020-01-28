// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolMissionsPluginRegionLocaleChangedEvent lol missions plugin region locale changed event
// swagger:model LolMissionsPluginRegionLocaleChangedEvent
type LolMissionsPluginRegionLocaleChangedEvent struct {

	// locale
	Locale string `json:"locale,omitempty"`
}

// Validate validates this lol missions plugin region locale changed event
func (m *LolMissionsPluginRegionLocaleChangedEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolMissionsPluginRegionLocaleChangedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMissionsPluginRegionLocaleChangedEvent) UnmarshalBinary(b []byte) error {
	var res LolMissionsPluginRegionLocaleChangedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
