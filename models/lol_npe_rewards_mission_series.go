// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeRewardsMissionSeries lol npe rewards mission series
// swagger:model LolNpeRewardsMissionSeries
type LolNpeRewardsMissionSeries struct {

	// id
	ID string `json:"id,omitempty"`

	// internal name
	InternalName string `json:"internalName,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this lol npe rewards mission series
func (m *LolNpeRewardsMissionSeries) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeRewardsMissionSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeRewardsMissionSeries) UnmarshalBinary(b []byte) error {
	var res LolNpeRewardsMissionSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
