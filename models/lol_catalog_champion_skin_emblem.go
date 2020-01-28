// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolCatalogChampionSkinEmblem lol catalog champion skin emblem
// swagger:model LolCatalogChampionSkinEmblem
type LolCatalogChampionSkinEmblem struct {

	// emblem path
	EmblemPath *LolCatalogChampionSkinEmblemPath `json:"emblemPath,omitempty"`

	// emblem position
	EmblemPosition *LolCatalogChampionSkinEmblemPosition `json:"emblemPosition,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this lol catalog champion skin emblem
func (m *LolCatalogChampionSkinEmblem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmblemPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmblemPosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCatalogChampionSkinEmblem) validateEmblemPath(formats strfmt.Registry) error {

	if swag.IsZero(m.EmblemPath) { // not required
		return nil
	}

	if m.EmblemPath != nil {
		if err := m.EmblemPath.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emblemPath")
			}
			return err
		}
	}

	return nil
}

func (m *LolCatalogChampionSkinEmblem) validateEmblemPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.EmblemPosition) { // not required
		return nil
	}

	if m.EmblemPosition != nil {
		if err := m.EmblemPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emblemPosition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCatalogChampionSkinEmblem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogChampionSkinEmblem) UnmarshalBinary(b []byte) error {
	var res LolCatalogChampionSkinEmblem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
