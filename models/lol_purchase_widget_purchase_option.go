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

// LolPurchaseWidgetPurchaseOption lol purchase widget purchase option
// swagger:model LolPurchaseWidgetPurchaseOption
type LolPurchaseWidgetPurchaseOption struct {

	// price details
	PriceDetails []*LolPurchaseWidgetPriceDetail `json:"priceDetails"`
}

// Validate validates this lol purchase widget purchase option
func (m *LolPurchaseWidgetPurchaseOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPurchaseWidgetPurchaseOption) validatePriceDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.PriceDetails); i++ {
		if swag.IsZero(m.PriceDetails[i]) { // not required
			continue
		}

		if m.PriceDetails[i] != nil {
			if err := m.PriceDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPurchaseWidgetPurchaseOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPurchaseWidgetPurchaseOption) UnmarshalBinary(b []byte) error {
	var res LolPurchaseWidgetPurchaseOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
