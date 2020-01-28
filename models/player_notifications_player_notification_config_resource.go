// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PlayerNotificationsPlayerNotificationConfigResource player notifications player notification config resource
// swagger:model PlayerNotificationsPlayerNotificationConfigResource
type PlayerNotificationsPlayerNotificationConfigResource struct {

	// expiration check frequency
	ExpirationCheckFrequency int64 `json:"ExpirationCheckFrequency,omitempty"`
}

// Validate validates this player notifications player notification config resource
func (m *PlayerNotificationsPlayerNotificationConfigResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlayerNotificationsPlayerNotificationConfigResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerNotificationsPlayerNotificationConfigResource) UnmarshalBinary(b []byte) error {
	var res PlayerNotificationsPlayerNotificationConfigResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}