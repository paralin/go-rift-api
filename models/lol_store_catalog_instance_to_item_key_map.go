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

// LolStoreCatalogInstanceToItemKeyMap lol store catalog instance to item key map
// swagger:model LolStoreCatalogInstanceToItemKeyMap
type LolStoreCatalogInstanceToItemKeyMap struct {

	// platform ids
	PlatformIds map[string]LolStoreItemKey `json:"platformIds,omitempty"`
}

// Validate validates this lol store catalog instance to item key map
func (m *LolStoreCatalogInstanceToItemKeyMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatformIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolStoreCatalogInstanceToItemKeyMap) validatePlatformIds(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformIds) { // not required
		return nil
	}

	for k := range m.PlatformIds {

		if err := validate.Required("platformIds"+"."+k, "body", m.PlatformIds[k]); err != nil {
			return err
		}
		if val, ok := m.PlatformIds[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreCatalogInstanceToItemKeyMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreCatalogInstanceToItemKeyMap) UnmarshalBinary(b []byte) error {
	var res LolStoreCatalogInstanceToItemKeyMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
