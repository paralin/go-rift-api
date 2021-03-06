// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatErrorResource lol chat error resource
// swagger:model LolChatErrorResource
type LolChatErrorResource struct {

	// code
	Code int64 `json:"code,omitempty"`

	// from
	From string `json:"from,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this lol chat error resource
func (m *LolChatErrorResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatErrorResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatErrorResource) UnmarshalBinary(b []byte) error {
	var res LolChatErrorResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
