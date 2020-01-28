// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolSummonerLevelProgression lol summoner level progression
// swagger:model LolSummonerLevelProgression
type LolSummonerLevelProgression struct {

	// final level boundary
	FinalLevelBoundary int64 `json:"finalLevelBoundary,omitempty"`

	// final xp
	FinalXp int64 `json:"finalXp,omitempty"`

	// initial level boundary
	InitialLevelBoundary int64 `json:"initialLevelBoundary,omitempty"`

	// initial xp
	InitialXp int64 `json:"initialXp,omitempty"`
}

// Validate validates this lol summoner level progression
func (m *LolSummonerLevelProgression) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolSummonerLevelProgression) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSummonerLevelProgression) UnmarshalBinary(b []byte) error {
	var res LolSummonerLevelProgression
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}