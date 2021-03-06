// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolChatMultiGamePresence lol chat multi game presence
// swagger:model LolChatMultiGamePresence
type LolChatMultiGamePresence struct {

	// actor
	Actor string `json:"actor,omitempty"`

	// basic
	Basic string `json:"basic,omitempty"`

	// details
	Details string `json:"details,omitempty"`

	// game name
	GameName string `json:"game_name,omitempty"`

	// game tag
	GameTag string `json:"game_tag,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// msg
	Msg string `json:"msg,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// patchline
	Patchline string `json:"patchline,omitempty"`

	// pid
	Pid string `json:"pid,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// private
	Private string `json:"private,omitempty"`

	// private jwt
	PrivateJwt string `json:"privateJwt,omitempty"`

	// product
	Product string `json:"product,omitempty"`

	// resource
	Resource string `json:"resource,omitempty"`

	// state
	State LolChatAccountState `json:"state,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`
}

// Validate validates this lol chat multi game presence
func (m *LolChatMultiGamePresence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChatMultiGamePresence) validateState(formats strfmt.Registry) error {

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
func (m *LolChatMultiGamePresence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatMultiGamePresence) UnmarshalBinary(b []byte) error {
	var res LolChatMultiGamePresence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
