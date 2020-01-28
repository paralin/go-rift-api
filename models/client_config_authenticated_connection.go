// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClientConfigAuthenticatedConnection client config authenticated connection
// swagger:model ClientConfigAuthenticatedConnection
type ClientConfigAuthenticatedConnection struct {

	// auth token
	AuthToken string `json:"authToken,omitempty"`

	// connection Id
	ConnectionID int32 `json:"connectionId,omitempty"`

	// subscribed
	Subscribed bool `json:"subscribed,omitempty"`
}

// Validate validates this client config authenticated connection
func (m *ClientConfigAuthenticatedConnection) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientConfigAuthenticatedConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientConfigAuthenticatedConnection) UnmarshalBinary(b []byte) error {
	var res ClientConfigAuthenticatedConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
