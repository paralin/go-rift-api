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

// LolCatalogBundled lol catalog bundled
// swagger:model LolCatalogBundled
type LolCatalogBundled struct {

	// flexible
	Flexible bool `json:"flexible,omitempty"`

	// items
	Items []*LolCatalogBundledItem `json:"items"`

	// minimum prices
	MinimumPrices []*LolCatalogBundledItemCost `json:"minimumPrices"`
}

// Validate validates this lol catalog bundled
func (m *LolCatalogBundled) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinimumPrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCatalogBundled) validateItems(formats strfmt.Registry) error {

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

func (m *LolCatalogBundled) validateMinimumPrices(formats strfmt.Registry) error {

	if swag.IsZero(m.MinimumPrices) { // not required
		return nil
	}

	for i := 0; i < len(m.MinimumPrices); i++ {
		if swag.IsZero(m.MinimumPrices[i]) { // not required
			continue
		}

		if m.MinimumPrices[i] != nil {
			if err := m.MinimumPrices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("minimumPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCatalogBundled) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogBundled) UnmarshalBinary(b []byte) error {
	var res LolCatalogBundled
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
