// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolQueueEligibilityEligibilityRestriction lol queue eligibility eligibility restriction
// swagger:model LolQueueEligibilityEligibilityRestriction
type LolQueueEligibilityEligibilityRestriction struct {

	// expired timestamp
	ExpiredTimestamp int64 `json:"expiredTimestamp,omitempty"`

	// restriction args
	RestrictionArgs map[string]string `json:"restrictionArgs,omitempty"`

	// restriction code
	RestrictionCode LolQueueEligibilityEligibilityRestrictionCode `json:"restrictionCode,omitempty"`

	// summoner ids
	SummonerIds []int64 `json:"summonerIds"`
}

// Validate validates this lol queue eligibility eligibility restriction
func (m *LolQueueEligibilityEligibilityRestriction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestrictionCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolQueueEligibilityEligibilityRestriction) validateRestrictionCode(formats strfmt.Registry) error {

	if swag.IsZero(m.RestrictionCode) { // not required
		return nil
	}

	if err := m.RestrictionCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("restrictionCode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolQueueEligibilityEligibilityRestriction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolQueueEligibilityEligibilityRestriction) UnmarshalBinary(b []byte) error {
	var res LolQueueEligibilityEligibilityRestriction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
