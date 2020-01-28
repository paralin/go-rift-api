// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClientConfigConfigStatus client config config status
// swagger:model ClientConfigConfigStatus
type ClientConfigConfigStatus struct {

	// readiness
	Readiness ClientConfigConfigReadinessEnum `json:"readiness,omitempty"`

	// update Id
	UpdateID int64 `json:"updateId,omitempty"`
}

// Validate validates this client config config status
func (m *ClientConfigConfigStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReadiness(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientConfigConfigStatus) validateReadiness(formats strfmt.Registry) error {

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
func (m *ClientConfigConfigStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientConfigConfigStatus) UnmarshalBinary(b []byte) error {
	var res ClientConfigConfigStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
