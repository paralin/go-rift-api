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

// ClashRewardConfigClient clash reward config client
// swagger:model ClashRewardConfigClient
type ClashRewardConfigClient struct {

	// entries
	Entries []*ClashRewardConfigEntry `json:"entries"`

	// grant to sub
	GrantToSub bool `json:"grantToSub,omitempty"`

	// key def
	KeyDef []ClashRewardKeyType `json:"keyDef"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this clash reward config client
func (m *ClashRewardConfigClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyDef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClashRewardConfigClient) validateEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.Entries) { // not required
		return nil
	}

	for i := 0; i < len(m.Entries); i++ {
		if swag.IsZero(m.Entries[i]) { // not required
			continue
		}

		if m.Entries[i] != nil {
			if err := m.Entries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClashRewardConfigClient) validateKeyDef(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyDef) { // not required
		return nil
	}

	for i := 0; i < len(m.KeyDef); i++ {

		if err := m.KeyDef[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyDef" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClashRewardConfigClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClashRewardConfigClient) UnmarshalBinary(b []byte) error {
	var res ClashRewardConfigClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}