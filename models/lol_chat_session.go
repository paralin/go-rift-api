// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolChatSession lol chat session
// swagger:model LolChatSession
type LolChatSession struct {

	// game name
	GameName string `json:"game_name,omitempty"`

	// game tag
	GameTag string `json:"game_tag,omitempty"`

	// loaded
	Loaded bool `json:"loaded,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pid
	Pid string `json:"pid,omitempty"`

	// resource
	Resource string `json:"resource,omitempty"`

	// state
	State LolChatChatSessionState `json:"state,omitempty"`
}

// Validate validates this lol chat session
func (m *LolChatSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChatSession) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChatSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatSession) UnmarshalBinary(b []byte) error {
	var res LolChatSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
