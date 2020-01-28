// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPreEndOfGameSequenceEvent lol pre end of game sequence event
// swagger:model LolPreEndOfGameSequenceEvent
type LolPreEndOfGameSequenceEvent struct {

	// name
	Name string `json:"name,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`
}

// Validate validates this lol pre end of game sequence event
func (m *LolPreEndOfGameSequenceEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPreEndOfGameSequenceEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPreEndOfGameSequenceEvent) UnmarshalBinary(b []byte) error {
	var res LolPreEndOfGameSequenceEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
