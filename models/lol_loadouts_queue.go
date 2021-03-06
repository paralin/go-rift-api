// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLoadoutsQueue lol loadouts queue
// swagger:model LolLoadoutsQueue
type LolLoadoutsQueue struct {

	// is team builder managed
	IsTeamBuilderManaged bool `json:"isTeamBuilderManaged,omitempty"`
}

// Validate validates this lol loadouts queue
func (m *LolLoadoutsQueue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLoadoutsQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoadoutsQueue) UnmarshalBinary(b []byte) error {
	var res LolLoadoutsQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
