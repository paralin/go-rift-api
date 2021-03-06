// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolChatChatMessage lol chat chat message
// swagger:model LolChatChatMessage
type LolChatChatMessage struct {

	// body
	Body string `json:"body,omitempty"`

	// cid
	Cid string `json:"cid,omitempty"`

	// game name
	GameName string `json:"game_name,omitempty"`

	// game tag
	GameTag string `json:"game_tag,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// mid
	Mid string `json:"mid,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pid
	Pid string `json:"pid,omitempty"`

	// read
	Read bool `json:"read,omitempty"`

	// time
	Time string `json:"time,omitempty"`

	// type
	Type LolChatMessageType `json:"type,omitempty"`
}

// Validate validates this lol chat chat message
func (m *LolChatChatMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChatChatMessage) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChatChatMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatChatMessage) UnmarshalBinary(b []byte) error {
	var res LolChatChatMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
