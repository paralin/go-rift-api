// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolChatConfigStatus lol chat config status
// swagger:model LolChatConfigStatus
type LolChatConfigStatus struct {

	// readiness
	Readiness LolChatConfigReadinessEnum `json:"readiness,omitempty"`
}

// Validate validates this lol chat config status
func (m *LolChatConfigStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReadiness(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChatConfigStatus) validateReadiness(formats strfmt.Registry) error {

	if swag.IsZero(m.Readiness) { // not required
		return nil
	}

	if err := m.Readiness.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("readiness")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChatConfigStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatConfigStatus) UnmarshalBinary(b []byte) error {
	var res LolChatConfigStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
