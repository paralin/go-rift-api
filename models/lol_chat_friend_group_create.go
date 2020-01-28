// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatFriendGroupCreate lol chat friend group create
// swagger:model LolChatFriendGroupCreate
type LolChatFriendGroupCreate struct {

	// collapsed
	Collapsed bool `json:"collapsed,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this lol chat friend group create
func (m *LolChatFriendGroupCreate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatFriendGroupCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatFriendGroupCreate) UnmarshalBinary(b []byte) error {
	var res LolChatFriendGroupCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
