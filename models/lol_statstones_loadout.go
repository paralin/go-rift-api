// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStatstonesLoadout lol statstones loadout
// swagger:model LolStatstonesLoadout
type LolStatstonesLoadout struct {

	// id
	ID string `json:"id,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// loadout
	Loadout map[string]interface{} `json:"loadout,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`
}

// Validate validates this lol statstones loadout
func (m *LolStatstonesLoadout) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesLoadout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesLoadout) UnmarshalBinary(b []byte) error {
	var res LolStatstonesLoadout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
