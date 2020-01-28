// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GameQueuesLcdsMatchingThrottleConfig game queues lcds matching throttle config
// swagger:model GameQueuesLcdsMatchingThrottleConfig
type GameQueuesLcdsMatchingThrottleConfig struct {

	// cache name
	CacheName string `json:"cacheName,omitempty"`

	// limit
	Limit int64 `json:"limit,omitempty"`
}

// Validate validates this game queues lcds matching throttle config
func (m *GameQueuesLcdsMatchingThrottleConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GameQueuesLcdsMatchingThrottleConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameQueuesLcdsMatchingThrottleConfig) UnmarshalBinary(b []byte) error {
	var res GameQueuesLcdsMatchingThrottleConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
