// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RankedScoutingTopChampionDTO ranked scouting top champion d t o
// swagger:model RankedScoutingTopChampionDTO
type RankedScoutingTopChampionDTO struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// game count
	GameCount int32 `json:"gameCount,omitempty"`

	// kda
	Kda float32 `json:"kda,omitempty"`

	// rank
	Rank int32 `json:"rank,omitempty"`

	// win count
	WinCount int32 `json:"winCount,omitempty"`
}

// Validate validates this ranked scouting top champion d t o
func (m *RankedScoutingTopChampionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RankedScoutingTopChampionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RankedScoutingTopChampionDTO) UnmarshalBinary(b []byte) error {
	var res RankedScoutingTopChampionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
