// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolPurchaseWidgetItemDefinition lol purchase widget item definition
// swagger:model LolPurchaseWidgetItemDefinition
type LolPurchaseWidgetItemDefinition struct {

	// assets
	Assets *LolPurchaseWidgetCatalogPluginItemAssets `json:"assets,omitempty"`

	// bundled item price
	BundledItemPrice *LolPurchaseWidgetBundledItemPricingInfo `json:"bundledItemPrice,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owned
	Owned bool `json:"owned,omitempty"`

	// sub inventory type
	SubInventoryType string `json:"subInventoryType,omitempty"`

	// sub title
	SubTitle string `json:"subTitle,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this lol purchase widget item definition
func (m *LolPurchaseWidgetItemDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundledItemPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPurchaseWidgetItemDefinition) validateAssets(formats strfmt.Registry) error {

	if swag.IsZero(m.Assets) { // not required
		return nil
	}

	if m.Assets != nil {
		if err := m.Assets.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assets")
			}
			return err
		}
	}

	return nil
}

func (m *LolPurchaseWidgetItemDefinition) validateBundledItemPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.BundledItemPrice) { // not required
		return nil
	}

	if m.BundledItemPrice != nil {
		if err := m.BundledItemPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundledItemPrice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPurchaseWidgetItemDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPurchaseWidgetItemDefinition) UnmarshalBinary(b []byte) error {
	var res LolPurchaseWidgetItemDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}