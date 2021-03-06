// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRegaliaRegaliaPreferences lol regalia regalia preferences
// swagger:model LolRegaliaRegaliaPreferences
type LolRegaliaRegaliaPreferences struct {

	// preferred banner type
	PreferredBannerType string `json:"preferredBannerType,omitempty"`

	// preferred crest type
	PreferredCrestType string `json:"preferredCrestType,omitempty"`
}

// Validate validates this lol regalia regalia preferences
func (m *LolRegaliaRegaliaPreferences) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRegaliaRegaliaPreferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRegaliaRegaliaPreferences) UnmarshalBinary(b []byte) error {
	var res LolRegaliaRegaliaPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
