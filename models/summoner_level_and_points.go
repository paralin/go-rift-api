// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SummonerLevelAndPoints summoner level and points
// swagger:model SummonerLevelAndPoints
type SummonerLevelAndPoints struct {

	// exp points
	ExpPoints int64 `json:"expPoints,omitempty"`

	// exp to next level
	ExpToNextLevel int64 `json:"expToNextLevel,omitempty"`

	// sum Id
	SumID int64 `json:"sumId,omitempty"`

	// summoner level
	SummonerLevel int32 `json:"summonerLevel,omitempty"`
}

// Validate validates this summoner level and points
func (m *SummonerLevelAndPoints) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SummonerLevelAndPoints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummonerLevelAndPoints) UnmarshalBinary(b []byte) error {
	var res SummonerLevelAndPoints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
