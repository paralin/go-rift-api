// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRsoAuthRSOPlayerCredentials lol rso auth r s o player credentials
// swagger:model LolRsoAuthRSOPlayerCredentials
type LolRsoAuthRSOPlayerCredentials struct {

	// password
	Password string `json:"password,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this lol rso auth r s o player credentials
func (m *LolRsoAuthRSOPlayerCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRsoAuthRSOPlayerCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRsoAuthRSOPlayerCredentials) UnmarshalBinary(b []byte) error {
	var res LolRsoAuthRSOPlayerCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}