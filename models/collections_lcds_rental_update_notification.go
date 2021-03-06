// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CollectionsLcdsRentalUpdateNotification collections lcds rental update notification
// swagger:model CollectionsLcdsRentalUpdateNotification
type CollectionsLcdsRentalUpdateNotification struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`
}

// Validate validates this collections lcds rental update notification
func (m *CollectionsLcdsRentalUpdateNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectionsLcdsRentalUpdateNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectionsLcdsRentalUpdateNotification) UnmarshalBinary(b []byte) error {
	var res CollectionsLcdsRentalUpdateNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
