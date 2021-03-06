// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatConversationUpdate lol chat conversation update
// swagger:model LolChatConversationUpdate
type LolChatConversationUpdate struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// hidden
	Hidden bool `json:"hidden,omitempty"`

	// muted
	Muted bool `json:"muted,omitempty"`
}

// Validate validates this lol chat conversation update
func (m *LolChatConversationUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatConversationUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatConversationUpdate) UnmarshalBinary(b []byte) error {
	var res LolChatConversationUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
