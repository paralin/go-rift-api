// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LCDSLoyaltyStateChangeNotification l c d s loyalty state change notification
// swagger:model LCDSLoyaltyStateChangeNotification
type LCDSLoyaltyStateChangeNotification struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// notification category
	NotificationCategory LCDSLoyaltyStateChangeNotificationCategory `json:"notificationCategory,omitempty"`

	// rewards
	Rewards *LCDSLoyaltyRewards `json:"rewards,omitempty"`
}

// Validate validates this l c d s loyalty state change notification
func (m *LCDSLoyaltyStateChangeNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LCDSLoyaltyStateChangeNotification) validateNotificationCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.NotificationCategory) { // not required
		return nil
	}

	if err := m.NotificationCategory.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notificationCategory")
		}
		return err
	}

	return nil
}

func (m *LCDSLoyaltyStateChangeNotification) validateRewards(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LCDSLoyaltyStateChangeNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LCDSLoyaltyStateChangeNotification) UnmarshalBinary(b []byte) error {
	var res LCDSLoyaltyStateChangeNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}