// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolCatalogCatalogPluginPrice lol catalog catalog plugin price
// swagger:model LolCatalogCatalogPluginPrice
type LolCatalogCatalogPluginPrice struct {

	// cost
	Cost int64 `json:"cost,omitempty"`

	// cost type
	CostType string `json:"costType,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// sale
	Sale *LolCatalogCatalogPluginSale `json:"sale,omitempty"`
}

// Validate validates this lol catalog catalog plugin price
func (m *LolCatalogCatalogPluginPrice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCatalogCatalogPluginPrice) validateSale(formats strfmt.Registry) error {

	if swag.IsZero(m.Sale) { // not required
		return nil
	}

	if m.Sale != nil {
		if err := m.Sale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sale")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCatalogCatalogPluginPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogCatalogPluginPrice) UnmarshalBinary(b []byte) error {
	var res LolCatalogCatalogPluginPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
