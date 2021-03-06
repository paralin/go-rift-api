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

// ClashRewardConfigEntry clash reward config entry
// swagger:model ClashRewardConfigEntry
type ClashRewardConfigEntry struct {

	// key
	Key string `json:"key,omitempty"`

	// vals
	Vals []*ClashRewardOutput `json:"vals"`
}

// Validate validates this clash reward config entry
func (m *ClashRewardConfigEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClashRewardConfigEntry) validateVals(formats strfmt.Registry) error {

	if swag.IsZero(m.Vals) { // not required
		return nil
	}

	for i := 0; i < len(m.Vals); i++ {
		if swag.IsZero(m.Vals[i]) { // not required
			continue
		}

		if m.Vals[i] != nil {
			if err := m.Vals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClashRewardConfigEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClashRewardConfigEntry) UnmarshalBinary(b []byte) error {
	var res ClashRewardConfigEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
