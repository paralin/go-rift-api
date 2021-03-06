// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathIds lol npe tutorial path ids
// swagger:model LolNpeTutorialPathIds
type LolNpeTutorialPathIds struct {

	// mission ids
	MissionIds []string `json:"missionIds"`

	// series ids
	SeriesIds []string `json:"seriesIds"`
}

// Validate validates this lol npe tutorial path ids
func (m *LolNpeTutorialPathIds) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathIds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathIds) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathIds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
