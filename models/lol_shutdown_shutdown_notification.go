// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolShutdownShutdownNotification lol shutdown shutdown notification
// swagger:model LolShutdownShutdownNotification
type LolShutdownShutdownNotification struct {

	// additional info
	AdditionalInfo string `json:"additionalInfo,omitempty"`

	// countdown
	Countdown float32 `json:"countdown,omitempty"`

	// reason
	Reason LolShutdownShutdownReason `json:"reason,omitempty"`
}

// Validate validates this lol shutdown shutdown notification
func (m *LolShutdownShutdownNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolShutdownShutdownNotification) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := m.Reason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolShutdownShutdownNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolShutdownShutdownNotification) UnmarshalBinary(b []byte) error {
	var res LolShutdownShutdownNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
