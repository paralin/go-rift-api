// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolServiceStatusBroadcastMessage lol service status broadcast message
// swagger:model LolServiceStatusBroadcastMessage
type LolServiceStatusBroadcastMessage struct {

	// content
	Content string `json:"content,omitempty"`

	// message key
	MessageKey string `json:"messageKey,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`
}

// Validate validates this lol service status broadcast message
func (m *LolServiceStatusBroadcastMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolServiceStatusBroadcastMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolServiceStatusBroadcastMessage) UnmarshalBinary(b []byte) error {
	var res LolServiceStatusBroadcastMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
