// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEmailVerificationRegionLocale lol email verification region locale
// swagger:model LolEmailVerificationRegionLocale
type LolEmailVerificationRegionLocale struct {

	// locale
	Locale string `json:"locale,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// web language
	WebLanguage string `json:"webLanguage,omitempty"`

	// web region
	WebRegion string `json:"webRegion,omitempty"`
}

// Validate validates this lol email verification region locale
func (m *LolEmailVerificationRegionLocale) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEmailVerificationRegionLocale) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEmailVerificationRegionLocale) UnmarshalBinary(b []byte) error {
	var res LolEmailVerificationRegionLocale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}