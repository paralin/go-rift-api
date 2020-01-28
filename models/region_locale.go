// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RegionLocale region locale
// swagger:model RegionLocale
type RegionLocale struct {

	// locale
	Locale string `json:"locale,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// web language
	WebLanguage string `json:"webLanguage,omitempty"`

	// web region
	WebRegion string `json:"webRegion,omitempty"`
}

// Validate validates this region locale
func (m *RegionLocale) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegionLocale) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegionLocale) UnmarshalBinary(b []byte) error {
	var res RegionLocale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
