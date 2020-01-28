// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VoiceChatDeviceResource voice chat device resource
// swagger:model VoiceChatDeviceResource
type VoiceChatDeviceResource struct {

	// handle
	Handle string `json:"handle,omitempty"`

	// is current device
	IsCurrentDevice bool `json:"is_current_device,omitempty"`

	// is default
	IsDefault bool `json:"is_default,omitempty"`

	// is effective device
	IsEffectiveDevice bool `json:"is_effective_device,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this voice chat device resource
func (m *VoiceChatDeviceResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VoiceChatDeviceResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoiceChatDeviceResource) UnmarshalBinary(b []byte) error {
	var res VoiceChatDeviceResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
