// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPatchChunkingPatcherEnvironment lol patch chunking patcher environment
// swagger:model LolPatchChunkingPatcherEnvironment
type LolPatchChunkingPatcherEnvironment struct {

	// client patcher available
	ClientPatcherAvailable bool `json:"client_patcher_available,omitempty"`

	// client patcher enabled
	ClientPatcherEnabled bool `json:"client_patcher_enabled,omitempty"`

	// game patcher available
	GamePatcherAvailable bool `json:"game_patcher_available,omitempty"`

	// game patcher enabled
	GamePatcherEnabled bool `json:"game_patcher_enabled,omitempty"`
}

// Validate validates this lol patch chunking patcher environment
func (m *LolPatchChunkingPatcherEnvironment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPatchChunkingPatcherEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPatchChunkingPatcherEnvironment) UnmarshalBinary(b []byte) error {
	var res LolPatchChunkingPatcherEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
