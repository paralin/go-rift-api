// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLoginRegionStatus lol login region status
// swagger:model LolLoginRegionStatus
type LolLoginRegionStatus struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// is l q fallback allowed
	IsLQFallbackAllowed bool `json:"isLQFallbackAllowed,omitempty"`

	// is user info enabled
	IsUserInfoEnabled bool `json:"isUserInfoEnabled,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`
}

// Validate validates this lol login region status
func (m *LolLoginRegionStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLoginRegionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoginRegionStatus) UnmarshalBinary(b []byte) error {
	var res LolLoginRegionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
