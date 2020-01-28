// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolReplaysReplaysSettingsData lol replays replays settings data
// swagger:model LolReplaysReplaysSettingsData
type LolReplaysReplaysSettingsData struct {

	// highlights folder path
	HighlightsFolderPath string `json:"highlights-folder-path,omitempty"`

	// replays folder path
	ReplaysFolderPath string `json:"replays-folder-path,omitempty"`
}

// Validate validates this lol replays replays settings data
func (m *LolReplaysReplaysSettingsData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolReplaysReplaysSettingsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolReplaysReplaysSettingsData) UnmarshalBinary(b []byte) error {
	var res LolReplaysReplaysSettingsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}