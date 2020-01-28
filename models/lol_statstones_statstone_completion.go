// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStatstonesStatstoneCompletion lol statstones statstone completion
// swagger:model LolStatstonesStatstoneCompletion
type LolStatstonesStatstoneCompletion struct {

	// category
	Category string `json:"category,omitempty"`

	// is epic
	IsEpic bool `json:"isEpic,omitempty"`

	// statstone name
	StatstoneName string `json:"statstoneName,omitempty"`
}

// Validate validates this lol statstones statstone completion
func (m *LolStatstonesStatstoneCompletion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesStatstoneCompletion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesStatstoneCompletion) UnmarshalBinary(b []byte) error {
	var res LolStatstonesStatstoneCompletion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}