// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LolBannersLoadout lol banners loadout
// swagger:model LolBannersLoadout
type LolBannersLoadout struct {

	// id
	ID string `json:"id,omitempty"`

	// loadout
	Loadout map[string]LolBannersLoadoutsSlot `json:"loadout,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`
}

// Validate validates this lol banners loadout
func (m *LolBannersLoadout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolBannersLoadout) validateLoadout(formats strfmt.Registry) error {

	if swag.IsZero(m.Loadout) { // not required
		return nil
	}

	for k := range m.Loadout {

		if err := validate.Required("loadout"+"."+k, "body", m.Loadout[k]); err != nil {
			return err
		}
		if val, ok := m.Loadout[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolBannersLoadout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolBannersLoadout) UnmarshalBinary(b []byte) error {
	var res LolBannersLoadout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
