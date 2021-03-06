// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolInventoryLoyaltyStatusNotification lol inventory loyalty status notification
// swagger:model LolInventoryLoyaltyStatusNotification
type LolInventoryLoyaltyStatusNotification struct {

	// rewards
	Rewards *LolInventoryLoyaltyRewards `json:"rewards,omitempty"`

	// status
	Status LolInventoryLoyaltyStatus `json:"status,omitempty"`
}

// Validate validates this lol inventory loyalty status notification
func (m *LolInventoryLoyaltyStatusNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolInventoryLoyaltyStatusNotification) validateRewards(formats strfmt.Registry) error {

	if swag.IsZero(m.Rewards) { // not required
		return nil
	}

	if m.Rewards != nil {
		if err := m.Rewards.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rewards")
			}
			return err
		}
	}

	return nil
}

func (m *LolInventoryLoyaltyStatusNotification) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolInventoryLoyaltyStatusNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolInventoryLoyaltyStatusNotification) UnmarshalBinary(b []byte) error {
	var res LolInventoryLoyaltyStatusNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
