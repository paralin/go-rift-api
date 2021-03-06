// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClashGameflowAvailability lol clash gameflow availability
// swagger:model LolClashGameflowAvailability
type LolClashGameflowAvailability struct {

	// is available
	IsAvailable bool `json:"isAvailable,omitempty"`
}

// Validate validates this lol clash gameflow availability
func (m *LolClashGameflowAvailability) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClashGameflowAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashGameflowAvailability) UnmarshalBinary(b []byte) error {
	var res LolClashGameflowAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
