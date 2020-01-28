// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRiotclientUpgraderGameflowAvailability lol riotclient upgrader gameflow availability
// swagger:model LolRiotclientUpgraderGameflowAvailability
type LolRiotclientUpgraderGameflowAvailability struct {

	// is available
	IsAvailable bool `json:"isAvailable,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this lol riotclient upgrader gameflow availability
func (m *LolRiotclientUpgraderGameflowAvailability) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRiotclientUpgraderGameflowAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRiotclientUpgraderGameflowAvailability) UnmarshalBinary(b []byte) error {
	var res LolRiotclientUpgraderGameflowAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
