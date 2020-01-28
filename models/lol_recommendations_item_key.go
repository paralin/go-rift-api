// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRecommendationsItemKey lol recommendations item key
// swagger:model LolRecommendationsItemKey
type LolRecommendationsItemKey struct {

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`
}

// Validate validates this lol recommendations item key
func (m *LolRecommendationsItemKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRecommendationsItemKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRecommendationsItemKey) UnmarshalBinary(b []byte) error {
	var res LolRecommendationsItemKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
