// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolItemSetsItemSetItem lol item sets item set item
// swagger:model LolItemSetsItemSetItem
type LolItemSetsItemSetItem struct {

	// count
	Count int64 `json:"count,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this lol item sets item set item
func (m *LolItemSetsItemSetItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolItemSetsItemSetItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolItemSetsItemSetItem) UnmarshalBinary(b []byte) error {
	var res LolItemSetsItemSetItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
