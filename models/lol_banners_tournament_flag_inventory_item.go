// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolBannersTournamentFlagInventoryItem lol banners tournament flag inventory item
// swagger:model LolBannersTournamentFlagInventoryItem
type LolBannersTournamentFlagInventoryItem struct {

	// item Id
	ItemID int32 `json:"itemId,omitempty"`

	// payload
	Payload *LolBannersCapClashFlagEntitlementPayload `json:"payload,omitempty"`

	// purchase date
	PurchaseDate string `json:"purchaseDate,omitempty"`
}

// Validate validates this lol banners tournament flag inventory item
func (m *LolBannersTournamentFlagInventoryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolBannersTournamentFlagInventoryItem) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolBannersTournamentFlagInventoryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolBannersTournamentFlagInventoryItem) UnmarshalBinary(b []byte) error {
	var res LolBannersTournamentFlagInventoryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}