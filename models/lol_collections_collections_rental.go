// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolCollectionsCollectionsRental lol collections collections rental
// swagger:model LolCollectionsCollectionsRental
type LolCollectionsCollectionsRental struct {

	// end date
	EndDate int64 `json:"endDate,omitempty"`

	// purchase date
	PurchaseDate int64 `json:"purchaseDate,omitempty"`

	// rented
	Rented bool `json:"rented,omitempty"`

	// win count remaining
	WinCountRemaining int32 `json:"winCountRemaining,omitempty"`
}

// Validate validates this lol collections collections rental
func (m *LolCollectionsCollectionsRental) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolCollectionsCollectionsRental) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCollectionsCollectionsRental) UnmarshalBinary(b []byte) error {
	var res LolCollectionsCollectionsRental
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
