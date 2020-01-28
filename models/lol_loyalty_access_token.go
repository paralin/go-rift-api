// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLoyaltyAccessToken lol loyalty access token
// swagger:model LolLoyaltyAccessToken
type LolLoyaltyAccessToken struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this lol loyalty access token
func (m *LolLoyaltyAccessToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLoyaltyAccessToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoyaltyAccessToken) UnmarshalBinary(b []byte) error {
	var res LolLoyaltyAccessToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}