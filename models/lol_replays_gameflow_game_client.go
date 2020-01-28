// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolReplaysGameflowGameClient lol replays gameflow game client
// swagger:model LolReplaysGameflowGameClient
type LolReplaysGameflowGameClient struct {

	// running
	Running bool `json:"running,omitempty"`
}

// Validate validates this lol replays gameflow game client
func (m *LolReplaysGameflowGameClient) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolReplaysGameflowGameClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolReplaysGameflowGameClient) UnmarshalBinary(b []byte) error {
	var res LolReplaysGameflowGameClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}