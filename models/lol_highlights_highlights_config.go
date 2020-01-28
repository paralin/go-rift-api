// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolHighlightsHighlightsConfig lol highlights highlights config
// swagger:model LolHighlightsHighlightsConfig
type LolHighlightsHighlightsConfig struct {

	// invalid highlight name characters
	InvalidHighlightNameCharacters string `json:"invalidHighlightNameCharacters,omitempty"`

	// is highlights enabled
	IsHighlightsEnabled bool `json:"isHighlightsEnabled,omitempty"`
}

// Validate validates this lol highlights highlights config
func (m *LolHighlightsHighlightsConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolHighlightsHighlightsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHighlightsHighlightsConfig) UnmarshalBinary(b []byte) error {
	var res LolHighlightsHighlightsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
