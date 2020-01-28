// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPurchaseWidgetItemDetails lol purchase widget item details
// swagger:model LolPurchaseWidgetItemDetails
type LolPurchaseWidgetItemDetails struct {

	// description
	Description string `json:"description,omitempty"`

	// icon Url
	IconURL string `json:"iconUrl,omitempty"`

	// sub title
	SubTitle string `json:"subTitle,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this lol purchase widget item details
func (m *LolPurchaseWidgetItemDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPurchaseWidgetItemDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPurchaseWidgetItemDetails) UnmarshalBinary(b []byte) error {
	var res LolPurchaseWidgetItemDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
