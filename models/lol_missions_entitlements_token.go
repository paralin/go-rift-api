// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolMissionsEntitlementsToken lol missions entitlements token
// swagger:model LolMissionsEntitlementsToken
type LolMissionsEntitlementsToken struct {

	// entitlements
	Entitlements []string `json:"entitlements"`
}

// Validate validates this lol missions entitlements token
func (m *LolMissionsEntitlementsToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolMissionsEntitlementsToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMissionsEntitlementsToken) UnmarshalBinary(b []byte) error {
	var res LolMissionsEntitlementsToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}