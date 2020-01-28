// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolInventoryInventoryItemDTO lol inventory inventory item d t o
// swagger:model LolInventoryInventoryItemDTO
type LolInventoryInventoryItemDTO struct {

	// entitlement Id
	EntitlementID string `json:"entitlementId,omitempty"`

	// entitlement type Id
	EntitlementTypeID string `json:"entitlementTypeId,omitempty"`

	// expiration date
	ExpirationDate string `json:"expirationDate,omitempty"`

	// f2p
	F2p bool `json:"f2p,omitempty"`

	// instance Id
	InstanceID string `json:"instanceId,omitempty"`

	// instance type Id
	InstanceTypeID string `json:"instanceTypeId,omitempty"`

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// loyalty
	Loyalty bool `json:"loyalty,omitempty"`

	// lsb
	Lsb bool `json:"lsb,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`

	// purchase date
	PurchaseDate string `json:"purchaseDate,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// rental
	Rental bool `json:"rental,omitempty"`

	// used in game date
	UsedInGameDate string `json:"usedInGameDate,omitempty"`

	// wins
	Wins int64 `json:"wins,omitempty"`
}

// Validate validates this lol inventory inventory item d t o
func (m *LolInventoryInventoryItemDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolInventoryInventoryItemDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolInventoryInventoryItemDTO) UnmarshalBinary(b []byte) error {
	var res LolInventoryInventoryItemDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}