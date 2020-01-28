// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPremadeVoiceLocalSettingsCategoryDataResource lol premade voice local settings category data resource
// swagger:model LolPremadeVoiceLocalSettingsCategoryDataResource
type LolPremadeVoiceLocalSettingsCategoryDataResource struct {

	// current capture device handle
	CurrentCaptureDeviceHandle string `json:"currentCaptureDeviceHandle,omitempty"`

	// input volume
	InputVolume int32 `json:"inputVolume,omitempty"`

	// vad sensitivity
	VadSensitivity int32 `json:"vadSensitivity,omitempty"`
}

// Validate validates this lol premade voice local settings category data resource
func (m *LolPremadeVoiceLocalSettingsCategoryDataResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPremadeVoiceLocalSettingsCategoryDataResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPremadeVoiceLocalSettingsCategoryDataResource) UnmarshalBinary(b []byte) error {
	var res LolPremadeVoiceLocalSettingsCategoryDataResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}