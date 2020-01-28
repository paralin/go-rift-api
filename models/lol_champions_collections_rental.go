// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampionsCollectionsRental lol champions collections rental
// swagger:model LolChampionsCollectionsRental
type LolChampionsCollectionsRental struct {

	// end date
	EndDate int64 `json:"endDate,omitempty"`

	// purchase date
	PurchaseDate int64 `json:"purchaseDate,omitempty"`

	// rented
	Rented bool `json:"rented,omitempty"`

	// win count remaining
	WinCountRemaining int32 `json:"winCountRemaining,omitempty"`
}

// Validate validates this lol champions collections rental
func (m *LolChampionsCollectionsRental) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampionsCollectionsRental) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampionsCollectionsRental) UnmarshalBinary(b []byte) error {
	var res LolChampionsCollectionsRental
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}