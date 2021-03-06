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

// LolStatstonesStatstoneSetCompleteVignette lol statstones statstone set complete vignette
// swagger:model LolStatstonesStatstoneSetCompleteVignette
type LolStatstonesStatstoneSetCompleteVignette struct {

	// statstones
	Statstones []*LolStatstonesStatstoneCompletion `json:"statstones"`
}

// Validate validates this lol statstones statstone set complete vignette
func (m *LolStatstonesStatstoneSetCompleteVignette) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatstones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolStatstonesStatstoneSetCompleteVignette) validateStatstones(formats strfmt.Registry) error {

	if swag.IsZero(m.Statstones) { // not required
		return nil
	}

	for i := 0; i < len(m.Statstones); i++ {
		if swag.IsZero(m.Statstones[i]) { // not required
			continue
		}

		if m.Statstones[i] != nil {
			if err := m.Statstones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statstones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesStatstoneSetCompleteVignette) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesStatstoneSetCompleteVignette) UnmarshalBinary(b []byte) error {
	var res LolStatstonesStatstoneSetCompleteVignette
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
