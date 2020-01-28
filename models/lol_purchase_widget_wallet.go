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

// LolPurchaseWidgetWallet lol purchase widget wallet
// swagger:model LolPurchaseWidgetWallet
type LolPurchaseWidgetWallet struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// balances
	Balances []*LolPurchaseWidgetBalance `json:"balances"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this lol purchase widget wallet
func (m *LolPurchaseWidgetWallet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPurchaseWidgetWallet) validateBalances(formats strfmt.Registry) error {

	if swag.IsZero(m.Balances) { // not required
		return nil
	}

	for i := 0; i < len(m.Balances); i++ {
		if swag.IsZero(m.Balances[i]) { // not required
			continue
		}

		if m.Balances[i] != nil {
			if err := m.Balances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("balances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPurchaseWidgetWallet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPurchaseWidgetWallet) UnmarshalBinary(b []byte) error {
	var res LolPurchaseWidgetWallet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
