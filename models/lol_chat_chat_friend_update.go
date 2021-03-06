// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatChatFriendUpdate lol chat chat friend update
// swagger:model LolChatChatFriendUpdate
type LolChatChatFriendUpdate struct {

	// group
	Group string `json:"group,omitempty"`

	// note
	Note string `json:"note,omitempty"`

	// pid
	Pid string `json:"pid,omitempty"`
}

// Validate validates this lol chat chat friend update
func (m *LolChatChatFriendUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatChatFriendUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatChatFriendUpdate) UnmarshalBinary(b []byte) error {
	var res LolChatChatFriendUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
