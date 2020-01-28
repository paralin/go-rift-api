// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLoadoutsLoadout lol loadouts loadout
// swagger:model LolLoadoutsLoadout
type LolLoadoutsLoadout struct {

	// id
	ID int32 `json:"id,omitempty"`

	// items
	Items map[string]interface{} `json:"items,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this lol loadouts loadout
func (m *LolLoadoutsLoadout) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLoadoutsLoadout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoadoutsLoadout) UnmarshalBinary(b []byte) error {
	var res LolLoadoutsLoadout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}