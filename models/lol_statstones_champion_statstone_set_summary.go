// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStatstonesChampionStatstoneSetSummary lol statstones champion statstone set summary
// swagger:model LolStatstonesChampionStatstoneSetSummary
type LolStatstonesChampionStatstoneSetSummary struct {

	// milestones passed
	MilestonesPassed int32 `json:"milestonesPassed,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// stones available
	StonesAvailable int32 `json:"stonesAvailable,omitempty"`

	// stones illuminated
	StonesIlluminated int32 `json:"stonesIlluminated,omitempty"`

	// stones owned
	StonesOwned int32 `json:"stonesOwned,omitempty"`
}

// Validate validates this lol statstones champion statstone set summary
func (m *LolStatstonesChampionStatstoneSetSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesChampionStatstoneSetSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesChampionStatstoneSetSummary) UnmarshalBinary(b []byte) error {
	var res LolStatstonesChampionStatstoneSetSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}