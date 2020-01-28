// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStoreOrderNotificationResource lol store order notification resource
// swagger:model LolStoreOrderNotificationResource
type LolStoreOrderNotificationResource struct {

	// event type
	EventType string `json:"eventType,omitempty"`

	// event type Id
	EventTypeID string `json:"eventTypeId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this lol store order notification resource
func (m *LolStoreOrderNotificationResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreOrderNotificationResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreOrderNotificationResource) UnmarshalBinary(b []byte) error {
	var res LolStoreOrderNotificationResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}