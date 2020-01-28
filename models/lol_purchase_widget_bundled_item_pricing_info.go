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

// LolPurchaseWidgetBundledItemPricingInfo lol purchase widget bundled item pricing info
// swagger:model LolPurchaseWidgetBundledItemPricingInfo
type LolPurchaseWidgetBundledItemPricingInfo struct {

	// discount prices
	DiscountPrices []*LolPurchaseWidgetDiscountPricingInfo `json:"discountPrices"`

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`
}

// Validate validates this lol purchase widget bundled item pricing info
func (m *LolPurchaseWidgetBundledItemPricingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountPrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPurchaseWidgetBundledItemPricingInfo) validateDiscountPrices(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountPrices) { // not required
		return nil
	}

	for i := 0; i < len(m.DiscountPrices); i++ {
		if swag.IsZero(m.DiscountPrices[i]) { // not required
			continue
		}

		if m.DiscountPrices[i] != nil {
			if err := m.DiscountPrices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discountPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPurchaseWidgetBundledItemPricingInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPurchaseWidgetBundledItemPricingInfo) UnmarshalBinary(b []byte) error {
	var res LolPurchaseWidgetBundledItemPricingInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
