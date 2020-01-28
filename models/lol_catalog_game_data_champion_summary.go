// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolCatalogGameDataChampionSummary lol catalog game data champion summary
// swagger:model LolCatalogGameDataChampionSummary
type LolCatalogGameDataChampionSummary struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// skins
	Skins []*LolCatalogGameDataChampionSkin `json:"skins"`

	// square portrait path
	SquarePortraitPath string `json:"squarePortraitPath,omitempty"`
}

// Validate validates this lol catalog game data champion summary
func (m *LolCatalogGameDataChampionSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSkins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCatalogGameDataChampionSummary) validateSkins(formats strfmt.Registry) error {

	if swag.IsZero(m.Skins) { // not required
		return nil
	}

	for i := 0; i < len(m.Skins); i++ {
		if swag.IsZero(m.Skins[i]) { // not required
			continue
		}

		if m.Skins[i] != nil {
			if err := m.Skins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCatalogGameDataChampionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogGameDataChampionSummary) UnmarshalBinary(b []byte) error {
	var res LolCatalogGameDataChampionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
