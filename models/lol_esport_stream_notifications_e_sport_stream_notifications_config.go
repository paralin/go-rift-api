// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEsportStreamNotificationsESportStreamNotificationsConfig lol esport stream notifications e sport stream notifications config
// swagger:model LolEsportStreamNotificationsESportStreamNotificationsConfig
type LolEsportStreamNotificationsESportStreamNotificationsConfig struct {

	// beapp failure long poll minutes
	BeappFailureLongPollMinutes int64 `json:"beappFailureLongPollMinutes,omitempty"`

	// notifications asset magick URL
	NotificationsAssetMagickURL string `json:"notificationsAssetMagickURL,omitempty"`

	// notifications enabled
	NotificationsEnabled bool `json:"notificationsEnabled,omitempty"`

	// notifications long poll minutes
	NotificationsLongPollMinutes int64 `json:"notificationsLongPollMinutes,omitempty"`

	// notifications service endpoint
	NotificationsServiceEndpoint string `json:"notificationsServiceEndpoint,omitempty"`

	// notifications service endpoint v2
	NotificationsServiceEndpointV2 string `json:"notificationsServiceEndpointV2,omitempty"`

	// notifications short poll minutes
	NotificationsShortPollMinutes int64 `json:"notificationsShortPollMinutes,omitempty"`

	// notifications stream group slug
	NotificationsStreamGroupSlug string `json:"notificationsStreamGroupSlug,omitempty"`

	// notifications stream URL
	NotificationsStreamURL string `json:"notificationsStreamURL,omitempty"`

	// use service endpoint v2
	UseServiceEndpointV2 bool `json:"useServiceEndpointV2,omitempty"`
}

// Validate validates this lol esport stream notifications e sport stream notifications config
func (m *LolEsportStreamNotificationsESportStreamNotificationsConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEsportStreamNotificationsESportStreamNotificationsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEsportStreamNotificationsESportStreamNotificationsConfig) UnmarshalBinary(b []byte) error {
	var res LolEsportStreamNotificationsESportStreamNotificationsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
