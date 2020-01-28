// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolCatalogItemDetails lol catalog item details
// swagger:model LolCatalogItemDetails
type LolCatalogItemDetails struct {

	// description
	Description string `json:"description,omitempty"`

	// icon Url
	IconURL string `json:"iconUrl,omitempty"`

	// sub title
	SubTitle string `json:"subTitle,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this lol catalog item details
func (m *LolCatalogItemDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolCatalogItemDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogItemDetails) UnmarshalBinary(b []byte) error {
	var res LolCatalogItemDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
