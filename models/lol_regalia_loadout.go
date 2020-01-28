// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolRegaliaLoadout lol regalia loadout
// swagger:model LolRegaliaLoadout
type LolRegaliaLoadout struct {

	// id
	ID string `json:"id,omitempty"`

	// loadout
	Loadout *LolRegaliaRegaliaLoadout `json:"loadout,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`
}

// Validate validates this lol regalia loadout
func (m *LolRegaliaLoadout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolRegaliaLoadout) validateLoadout(formats strfmt.Registry) error {

	if swag.IsZero(m.Loadout) { // not required
		return nil
	}

	if m.Loadout != nil {
		if err := m.Loadout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadout")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolRegaliaLoadout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRegaliaLoadout) UnmarshalBinary(b []byte) error {
	var res LolRegaliaLoadout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}