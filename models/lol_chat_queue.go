// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolChatQueue lol chat queue
// swagger:model LolChatQueue
type LolChatQueue struct {

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// game type config
	GameTypeConfig *LolChatQueueGameTypeConfig `json:"gameTypeConfig,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol chat queue
func (m *LolChatQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameTypeConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChatQueue) validateGameTypeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.GameTypeConfig) { // not required
		return nil
	}

	if m.GameTypeConfig != nil {
		if err := m.GameTypeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameTypeConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChatQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatQueue) UnmarshalBinary(b []byte) error {
	var res LolChatQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
