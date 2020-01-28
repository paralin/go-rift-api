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

// LolClashPlayerRewards lol clash player rewards
// swagger:model LolClashPlayerRewards
type LolClashPlayerRewards struct {

	// season vp
	SeasonVp int32 `json:"seasonVp,omitempty"`

	// theme vp
	ThemeVp []*LolClashThemeVp `json:"themeVp"`
}

// Validate validates this lol clash player rewards
func (m *LolClashPlayerRewards) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThemeVp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashPlayerRewards) validateThemeVp(formats strfmt.Registry) error {

	if swag.IsZero(m.ThemeVp) { // not required
		return nil
	}

	for i := 0; i < len(m.ThemeVp); i++ {
		if swag.IsZero(m.ThemeVp[i]) { // not required
			continue
		}

		if m.ThemeVp[i] != nil {
			if err := m.ThemeVp[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("themeVp" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashPlayerRewards) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashPlayerRewards) UnmarshalBinary(b []byte) error {
	var res LolClashPlayerRewards
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
