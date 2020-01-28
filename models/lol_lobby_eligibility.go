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

// LolLobbyEligibility lol lobby eligibility
// swagger:model LolLobbyEligibility
type LolLobbyEligibility struct {

	// eligible
	Eligible bool `json:"eligible,omitempty"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`

	// restrictions
	Restrictions []*LolLobbyEligibilityRestriction `json:"restrictions"`
}

// Validate validates this lol lobby eligibility
func (m *LolLobbyEligibility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyEligibility) validateRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.Restrictions) { // not required
		return nil
	}

	for i := 0; i < len(m.Restrictions); i++ {
		if swag.IsZero(m.Restrictions[i]) { // not required
			continue
		}

		if m.Restrictions[i] != nil {
			if err := m.Restrictions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyEligibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyEligibility) UnmarshalBinary(b []byte) error {
	var res LolLobbyEligibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}