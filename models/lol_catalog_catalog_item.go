// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LolCatalogCatalogItem lol catalog catalog item
// swagger:model LolCatalogCatalogItem
type LolCatalogCatalogItem struct {

	// active
	Active bool `json:"active,omitempty"`

	// bundled
	Bundled *LolCatalogBundled `json:"bundled,omitempty"`

	// icon Url
	IconURL string `json:"iconUrl,omitempty"`

	// inactive date
	InactiveDate string `json:"inactiveDate,omitempty"`

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// item instance Id
	ItemInstanceID string `json:"itemInstanceId,omitempty"`

	// item requirements
	ItemRequirements []*LolCatalogItemKey `json:"itemRequirements"`

	// localizations
	Localizations map[string]LolCatalogItemLocalization `json:"localizations,omitempty"`

	// prices
	Prices []*LolCatalogItemCost `json:"prices"`

	// release date
	ReleaseDate string `json:"releaseDate,omitempty"`

	// sale
	Sale *LolCatalogSale `json:"sale,omitempty"`

	// sub inventory type
	SubInventoryType string `json:"subInventoryType,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this lol catalog catalog item
func (m *LolCatalogCatalogItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBundled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemRequirements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCatalogCatalogItem) validateBundled(formats strfmt.Registry) error {

	if swag.IsZero(m.Bundled) { // not required
		return nil
	}

	if m.Bundled != nil {
		if err := m.Bundled.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundled")
			}
			return err
		}
	}

	return nil
}

func (m *LolCatalogCatalogItem) validateItemRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemRequirements) { // not required
		return nil
	}

	for i := 0; i < len(m.ItemRequirements); i++ {
		if swag.IsZero(m.ItemRequirements[i]) { // not required
			continue
		}

		if m.ItemRequirements[i] != nil {
			if err := m.ItemRequirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemRequirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolCatalogCatalogItem) validateLocalizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Localizations) { // not required
		return nil
	}

	for k := range m.Localizations {

		if err := validate.Required("localizations"+"."+k, "body", m.Localizations[k]); err != nil {
			return err
		}
		if val, ok := m.Localizations[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *LolCatalogCatalogItem) validatePrices(formats strfmt.Registry) error {

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

func (m *LolCatalogCatalogItem) validateSale(formats strfmt.Registry) error {

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
func (m *LolCatalogCatalogItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogCatalogItem) UnmarshalBinary(b []byte) error {
	var res LolCatalogCatalogItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
