// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashTournamentUpdatedNotification lol clash tournament updated notification
// swagger:model LolClashTournamentUpdatedNotification
type LolClashTournamentUpdatedNotification struct {

	// current retry
	CurrentRetry int32 `json:"currentRetry,omitempty"`

	// max retry
	MaxRetry int32 `json:"maxRetry,omitempty"`

	// missing player ids
	MissingPlayerIds []int64 `json:"missingPlayerIds"`

	// notify reason
	NotifyReason LolClashRosterNotifyReason `json:"notifyReason,omitempty"`

	// retry seconds
	RetrySeconds int64 `json:"retrySeconds,omitempty"`
}

// Validate validates this lol clash tournament updated notification
func (m *LolClashTournamentUpdatedNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifyReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashTournamentUpdatedNotification) validateNotifyReason(formats strfmt.Registry) error {

	if swag.IsZero(m.NotifyReason) { // not required
		return nil
	}

	if err := m.NotifyReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notifyReason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashTournamentUpdatedNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashTournamentUpdatedNotification) UnmarshalBinary(b []byte) error {
	var res LolClashTournamentUpdatedNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
