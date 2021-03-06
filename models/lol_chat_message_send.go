// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatMessageSend lol chat message send
// swagger:model LolChatMessageSend
type LolChatMessageSend struct {

	// message
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol chat message send
func (m *LolChatMessageSend) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatMessageSend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatMessageSend) UnmarshalBinary(b []byte) error {
	var res LolChatMessageSend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
