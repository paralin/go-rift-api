// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolCollectionsInventoryItem lol collections inventory item
// swagger:model LolCollectionsInventoryItem
type LolCollectionsInventoryItem struct {

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// ownership type
	OwnershipType LolCollectionsItemOwnershipType `json:"ownershipType,omitempty"`

	// purchase date
	PurchaseDate string `json:"purchaseDate,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this lol collections inventory item
func (m *LolCollectionsInventoryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnershipType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCollectionsInventoryItem) validateOwnershipType(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolCollectionsInventoryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCollectionsInventoryItem) UnmarshalBinary(b []byte) error {
	var res LolCollectionsInventoryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
