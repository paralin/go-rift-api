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

// LolLoadoutsFrontendInventoryResponse lol loadouts frontend inventory response
// swagger:model LolLoadoutsFrontendInventoryResponse
type LolLoadoutsFrontendInventoryResponse struct {

	// entitlements
	Entitlements []*LolLoadoutsItemKey `json:"entitlements"`
}

// Validate validates this lol loadouts frontend inventory response
func (m *LolLoadoutsFrontendInventoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLoadoutsFrontendInventoryResponse) validateEntitlements(formats strfmt.Registry) error {

	if swag.IsZero(m.Entitlements) { // not required
		return nil
	}

	for i := 0; i < len(m.Entitlements); i++ {
		if swag.IsZero(m.Entitlements[i]) { // not required
			continue
		}

		if m.Entitlements[i] != nil {
			if err := m.Entitlements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLoadoutsFrontendInventoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoadoutsFrontendInventoryResponse) UnmarshalBinary(b []byte) error {
	var res LolLoadoutsFrontendInventoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
