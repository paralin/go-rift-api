// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolStatstonesStatstoneRekindledVignette lol statstones statstone rekindled vignette
// swagger:model LolStatstonesStatstoneRekindledVignette
type LolStatstonesStatstoneRekindledVignette struct {

	// portrait path
	PortraitPath string `json:"portraitPath,omitempty"`

	// statstone
	Statstone *LolStatstonesStatstoneCompletion `json:"statstone,omitempty"`
}

// Validate validates this lol statstones statstone rekindled vignette
func (m *LolStatstonesStatstoneRekindledVignette) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatstone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolStatstonesStatstoneRekindledVignette) validateStatstone(formats strfmt.Registry) error {

	if swag.IsZero(m.Statstone) { // not required
		return nil
	}

	if m.Statstone != nil {
		if err := m.Statstone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statstone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesStatstoneRekindledVignette) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesStatstoneRekindledVignette) UnmarshalBinary(b []byte) error {
	var res LolStatstonesStatstoneRekindledVignette
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}