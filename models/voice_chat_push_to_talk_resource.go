// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VoiceChatPushToTalkResource voice chat push to talk resource
// swagger:model VoiceChatPushToTalkResource
type VoiceChatPushToTalkResource struct {

	// ptt enabled
	PttEnabled bool `json:"pttEnabled,omitempty"`

	// ptt key binding
	PttKeyBinding string `json:"pttKeyBinding,omitempty"`
}

// Validate validates this voice chat push to talk resource
func (m *VoiceChatPushToTalkResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VoiceChatPushToTalkResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoiceChatPushToTalkResource) UnmarshalBinary(b []byte) error {
	var res VoiceChatPushToTalkResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}