// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SeriesMediaDTO series media d t o
// swagger:model SeriesMediaDTO
type SeriesMediaDTO struct {

	// accent color
	AccentColor string `json:"accentColor,omitempty"`

	// background image large Url
	BackgroundImageLargeURL string `json:"backgroundImageLargeUrl,omitempty"`

	// background image small Url
	BackgroundImageSmallURL string `json:"backgroundImageSmallUrl,omitempty"`

	// background Url
	BackgroundURL string `json:"backgroundUrl,omitempty"`

	// tracker icon Url
	TrackerIconURL string `json:"trackerIconUrl,omitempty"`
}

// Validate validates this series media d t o
func (m *SeriesMediaDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SeriesMediaDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeriesMediaDTO) UnmarshalBinary(b []byte) error {
	var res SeriesMediaDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
