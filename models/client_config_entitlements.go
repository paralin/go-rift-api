// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClientConfigEntitlements client config entitlements
// swagger:model ClientConfigEntitlements
type ClientConfigEntitlements struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`

	// entitlements
	Entitlements []string `json:"entitlements"`

	// issuer
	Issuer string `json:"issuer,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this client config entitlements
func (m *ClientConfigEntitlements) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientConfigEntitlements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientConfigEntitlements) UnmarshalBinary(b []byte) error {
	var res ClientConfigEntitlements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
