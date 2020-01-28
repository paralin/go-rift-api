// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolPurchaseWidgetValidationRequestItem lol purchase widget validation request item
// swagger:model LolPurchaseWidgetValidationRequestItem
type LolPurchaseWidgetValidationRequestItem struct {

	// item key
	ItemKey *LolPurchaseWidgetItemKey `json:"itemKey,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`
}

// Validate validates this lol purchase widget validation request item
func (m *LolPurchaseWidgetValidationRequestItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPurchaseWidgetValidationRequestItem) validateItemKey(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemKey) { // not required
		return nil
	}

	if m.ItemKey != nil {
		if err := m.ItemKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPurchaseWidgetValidationRequestItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPurchaseWidgetValidationRequestItem) UnmarshalBinary(b []byte) error {
	var res LolPurchaseWidgetValidationRequestItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}