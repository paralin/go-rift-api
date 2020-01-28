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

// LolStoreSale lol store sale
// swagger:model LolStoreSale
type LolStoreSale struct {

	// end date
	EndDate string `json:"endDate,omitempty"`

	// prices
	Prices []*LolStoreItemCost `json:"prices"`

	// start date
	StartDate string `json:"startDate,omitempty"`
}

// Validate validates this lol store sale
func (m *LolStoreSale) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolStoreSale) validatePrices(formats strfmt.Registry) error {

	if swag.IsZero(m.Prices) { // not required
		return nil
	}

	for i := 0; i < len(m.Prices); i++ {
		if swag.IsZero(m.Prices[i]) { // not required
			continue
		}

		if m.Prices[i] != nil {
			if err := m.Prices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreSale) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreSale) UnmarshalBinary(b []byte) error {
	var res LolStoreSale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
