// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEmailVerificationAccessToken lol email verification access token
// swagger:model LolEmailVerificationAccessToken
type LolEmailVerificationAccessToken struct {

	// expiry
	Expiry int64 `json:"expiry,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this lol email verification access token
func (m *LolEmailVerificationAccessToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEmailVerificationAccessToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEmailVerificationAccessToken) UnmarshalBinary(b []byte) error {
	var res LolEmailVerificationAccessToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}