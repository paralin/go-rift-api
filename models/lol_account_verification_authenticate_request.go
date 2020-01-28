// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolAccountVerificationAuthenticateRequest lol account verification authenticate request
// swagger:model LolAccountVerificationAuthenticateRequest
type LolAccountVerificationAuthenticateRequest struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this lol account verification authenticate request
func (m *LolAccountVerificationAuthenticateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolAccountVerificationAuthenticateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolAccountVerificationAuthenticateRequest) UnmarshalBinary(b []byte) error {
	var res LolAccountVerificationAuthenticateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}