// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolMatchmakingGameflowGameTypeConfig lol matchmaking gameflow game type config
// swagger:model LolMatchmakingGameflowGameTypeConfig
type LolMatchmakingGameflowGameTypeConfig struct {

	// reroll
	Reroll bool `json:"reroll,omitempty"`
}

// Validate validates this lol matchmaking gameflow game type config
func (m *LolMatchmakingGameflowGameTypeConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolMatchmakingGameflowGameTypeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMatchmakingGameflowGameTypeConfig) UnmarshalBinary(b []byte) error {
	var res LolMatchmakingGameflowGameTypeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
