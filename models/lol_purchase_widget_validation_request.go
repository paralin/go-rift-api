// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolPurchaseWidgetValidationRequest lol purchase widget validation request
// swagger:model LolPurchaseWidgetValidationRequest
type LolPurchaseWidgetValidationRequest struct {

	// items
	Items []*LolPurchaseWidgetValidationRequestItem `json:"items"`

	// owned items
	OwnedItems []*LolPurchaseWidgetItemOwnership `json:"ownedItems"`
}

// Validate validates this lol purchase widget validation request
func (m *LolPurchaseWidgetValidationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnedItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPurchaseWidgetValidationRequest) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolPurchaseWidgetValidationRequest) validateOwnedItems(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnedItems) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnedItems); i++ {
		if swag.IsZero(m.OwnedItems[i]) { // not required
			continue
		}

		if m.OwnedItems[i] != nil {
			if err := m.OwnedItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ownedItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPurchaseWidgetValidationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPurchaseWidgetValidationRequest) UnmarshalBinary(b []byte) error {
	var res LolPurchaseWidgetValidationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}