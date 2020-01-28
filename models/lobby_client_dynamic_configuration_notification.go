// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LobbyClientDynamicConfigurationNotification lobby client dynamic configuration notification
// swagger:model LobbyClientDynamicConfigurationNotification
type LobbyClientDynamicConfigurationNotification struct {

	// configs
	Configs string `json:"configs,omitempty"`

	// delta
	Delta bool `json:"delta,omitempty"`
}

// Validate validates this lobby client dynamic configuration notification
func (m *LobbyClientDynamicConfigurationNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LobbyClientDynamicConfigurationNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LobbyClientDynamicConfigurationNotification) UnmarshalBinary(b []byte) error {
	var res LobbyClientDynamicConfigurationNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
