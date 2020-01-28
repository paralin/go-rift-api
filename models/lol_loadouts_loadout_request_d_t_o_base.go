// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLoadoutsLoadoutRequestDTOBase lol loadouts loadout request d t o base
// swagger:model LolLoadoutsLoadoutRequestDTOBase
type LolLoadoutsLoadoutRequestDTOBase struct {

	// service to jwts map
	ServiceToJwtsMap map[string]interface{} `json:"serviceToJwtsMap,omitempty"`
}

// Validate validates this lol loadouts loadout request d t o base
func (m *LolLoadoutsLoadoutRequestDTOBase) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLoadoutsLoadoutRequestDTOBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoadoutsLoadoutRequestDTOBase) UnmarshalBinary(b []byte) error {
	var res LolLoadoutsLoadoutRequestDTOBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
