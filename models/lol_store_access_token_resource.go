// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStoreAccessTokenResource lol store access token resource
// swagger:model LolStoreAccessTokenResource
type LolStoreAccessTokenResource struct {

	// expiry
	Expiry int64 `json:"expiry,omitempty"`

	// scopes
	Scopes []string `json:"scopes"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this lol store access token resource
func (m *LolStoreAccessTokenResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreAccessTokenResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreAccessTokenResource) UnmarshalBinary(b []byte) error {
	var res LolStoreAccessTokenResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}