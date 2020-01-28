// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRsoAuthAuthError lol rso auth auth error
// swagger:model LolRsoAuthAuthError
type LolRsoAuthAuthError struct {

	// error
	Error string `json:"error,omitempty"`

	// error description
	ErrorDescription string `json:"errorDescription,omitempty"`
}

// Validate validates this lol rso auth auth error
func (m *LolRsoAuthAuthError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRsoAuthAuthError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRsoAuthAuthError) UnmarshalBinary(b []byte) error {
	var res LolRsoAuthAuthError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
