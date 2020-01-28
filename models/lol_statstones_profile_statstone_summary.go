// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStatstonesProfileStatstoneSummary lol statstones profile statstone summary
// swagger:model LolStatstonesProfileStatstoneSummary
type LolStatstonesProfileStatstoneSummary struct {

	// category
	Category string `json:"category,omitempty"`

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this lol statstones profile statstone summary
func (m *LolStatstonesProfileStatstoneSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesProfileStatstoneSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesProfileStatstoneSummary) UnmarshalBinary(b []byte) error {
	var res LolStatstonesProfileStatstoneSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
