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

// PlayerLootResultDTO player loot result d t o
// swagger:model PlayerLootResultDTO
type PlayerLootResultDTO struct {

	// added
	Added []*PlayerLootDTO `json:"added"`

	// details
	Details string `json:"details,omitempty"`

	// redeemed
	Redeemed []string `json:"redeemed"`

	// removed
	Removed []*PlayerLootDTO `json:"removed"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this player loot result d t o
func (m *PlayerLootResultDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoved(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerLootResultDTO) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	for i := 0; i < len(m.Added); i++ {
		if swag.IsZero(m.Added[i]) { // not required
			continue
		}

		if m.Added[i] != nil {
			if err := m.Added[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("added" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlayerLootResultDTO) validateRemoved(formats strfmt.Registry) error {

	if swag.IsZero(m.Removed) { // not required
		return nil
	}

	for i := 0; i < len(m.Removed); i++ {
		if swag.IsZero(m.Removed[i]) { // not required
			continue
		}

		if m.Removed[i] != nil {
			if err := m.Removed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("removed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayerLootResultDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerLootResultDTO) UnmarshalBinary(b []byte) error {
	var res PlayerLootResultDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
