// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectChampSelectChatRoomDetails lol champ select champ select chat room details
// swagger:model LolChampSelectChampSelectChatRoomDetails
type LolChampSelectChampSelectChatRoomDetails struct {

	// chat room name
	ChatRoomName string `json:"chatRoomName,omitempty"`

	// chat room password
	ChatRoomPassword string `json:"chatRoomPassword,omitempty"`
}

// Validate validates this lol champ select champ select chat room details
func (m *LolChampSelectChampSelectChatRoomDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectChampSelectChatRoomDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectChampSelectChatRoomDetails) UnmarshalBinary(b []byte) error {
	var res LolChampSelectChampSelectChatRoomDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}