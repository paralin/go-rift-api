// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolMissionsMissionsSettingsDataResource lol missions missions settings data resource
// swagger:model LolMissionsMissionsSettingsDataResource
type LolMissionsMissionsSettingsDataResource struct {

	// selected series
	SelectedSeries string `json:"selected_series,omitempty"`
}

// Validate validates this lol missions missions settings data resource
func (m *LolMissionsMissionsSettingsDataResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolMissionsMissionsSettingsDataResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMissionsMissionsSettingsDataResource) UnmarshalBinary(b []byte) error {
	var res LolMissionsMissionsSettingsDataResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
