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

// LolCatalogCatalogPluginItem lol catalog catalog plugin item
// swagger:model LolCatalogCatalogPluginItem
type LolCatalogCatalogPluginItem struct {

	// description
	Description string `json:"description,omitempty"`

	// image path
	ImagePath string `json:"imagePath,omitempty"`

	// inactive date
	InactiveDate int64 `json:"inactiveDate,omitempty"`

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// item instance Id
	ItemInstanceID string `json:"itemInstanceId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owned
	Owned bool `json:"owned,omitempty"`

	// ownership type
	OwnershipType LolCatalogInventoryOwnership `json:"ownershipType,omitempty"`

	// prices
	Prices []*LolCatalogCatalogPluginPrice `json:"prices"`

	// purchase date
	PurchaseDate int64 `json:"purchaseDate,omitempty"`

	// release date
	ReleaseDate int64 `json:"releaseDate,omitempty"`

	// sub inventory type
	SubInventoryType string `json:"subInventoryType,omitempty"`

	// sub title
	SubTitle string `json:"subTitle,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this lol catalog catalog plugin item
func (m *LolCatalogCatalogPluginItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnershipType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCatalogCatalogPluginItem) validateOwnershipType(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnershipType) { // not required
		return nil
	}

	if err := m.OwnershipType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ownershipType")
		}
		return err
	}

	return nil
}

func (m *LolCatalogCatalogPluginItem) validatePrices(formats strfmt.Registry) error {

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
func (m *LolCatalogCatalogPluginItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogCatalogPluginItem) UnmarshalBinary(b []byte) error {
	var res LolCatalogCatalogPluginItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
