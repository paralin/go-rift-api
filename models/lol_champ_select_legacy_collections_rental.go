// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectLegacyCollectionsRental lol champ select legacy collections rental
// swagger:model LolChampSelectLegacyCollectionsRental
type LolChampSelectLegacyCollectionsRental struct {

	// rented
	Rented bool `json:"rented,omitempty"`
}

// Validate validates this lol champ select legacy collections rental
func (m *LolChampSelectLegacyCollectionsRental) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectLegacyCollectionsRental) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectLegacyCollectionsRental) UnmarshalBinary(b []byte) error {
	var res LolChampSelectLegacyCollectionsRental
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
