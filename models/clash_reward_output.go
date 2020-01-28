// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClashRewardOutput clash reward output
// swagger:model ClashRewardOutput
type ClashRewardOutput struct {

	// alternative
	Alternative *ClashRewardDefinition `json:"alternative,omitempty"`

	// grant
	Grant ClashRewardTime `json:"grant,omitempty"`

	// primary
	Primary *ClashRewardDefinition `json:"primary,omitempty"`

	// show
	Show ClashRewardTime `json:"show,omitempty"`
}

// Validate validates this clash reward output
func (m *ClashRewardOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClashRewardOutput) validateAlternative(formats strfmt.Registry) error {

	if swag.IsZero(m.Alternative) { // not required
		return nil
	}

	if m.Alternative != nil {
		if err := m.Alternative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alternative")
			}
			return err
		}
	}

	return nil
}

func (m *ClashRewardOutput) validateGrant(formats strfmt.Registry) error {

	if swag.IsZero(m.Grant) { // not required
		return nil
	}

	if err := m.Grant.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grant")
		}
		return err
	}

	return nil
}

func (m *ClashRewardOutput) validatePrimary(formats strfmt.Registry) error {

	if swag.IsZero(m.Primary) { // not required
		return nil
	}

	if m.Primary != nil {
		if err := m.Primary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

func (m *ClashRewardOutput) validateShow(formats strfmt.Registry) error {

	if swag.IsZero(m.Show) { // not required
		return nil
	}

	if err := m.Show.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("show")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClashRewardOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClashRewardOutput) UnmarshalBinary(b []byte) error {
	var res ClashRewardOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}