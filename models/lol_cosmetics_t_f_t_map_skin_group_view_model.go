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

// LolCosmeticsTFTMapSkinGroupViewModel lol cosmetics t f t map skin group view model
// swagger:model LolCosmeticsTFTMapSkinGroupViewModel
type LolCosmeticsTFTMapSkinGroupViewModel struct {

	// group Id
	GroupID int32 `json:"groupId,omitempty"`

	// group name
	GroupName string `json:"groupName,omitempty"`

	// items
	Items []*LolCosmeticsCosmeticsTFTMapSkinViewModel `json:"items"`

	// num available
	NumAvailable int32 `json:"numAvailable,omitempty"`

	// num owned
	NumOwned int32 `json:"numOwned,omitempty"`
}

// Validate validates this lol cosmetics t f t map skin group view model
func (m *LolCosmeticsTFTMapSkinGroupViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCosmeticsTFTMapSkinGroupViewModel) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCosmeticsTFTMapSkinGroupViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCosmeticsTFTMapSkinGroupViewModel) UnmarshalBinary(b []byte) error {
	var res LolCosmeticsTFTMapSkinGroupViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
