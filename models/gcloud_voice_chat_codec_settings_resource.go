// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GcloudVoiceChatCodecSettingsResource gcloud voice chat codec settings resource
// swagger:model GcloudVoiceChatCodecSettingsResource
type GcloudVoiceChatCodecSettingsResource struct {

	// codec bandwidth
	CodecBandwidth int32 `json:"codecBandwidth,omitempty"`

	// codec bitrate
	CodecBitrate int32 `json:"codecBitrate,omitempty"`

	// codec complexity
	CodecComplexity int32 `json:"codecComplexity,omitempty"`

	// codec vbr mode
	CodecVbrMode int32 `json:"codecVbrMode,omitempty"`
}

// Validate validates this gcloud voice chat codec settings resource
func (m *GcloudVoiceChatCodecSettingsResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcloudVoiceChatCodecSettingsResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcloudVoiceChatCodecSettingsResource) UnmarshalBinary(b []byte) error {
	var res GcloudVoiceChatCodecSettingsResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
