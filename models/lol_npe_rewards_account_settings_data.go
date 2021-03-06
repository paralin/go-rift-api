// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolNpeRewardsAccountSettingsData lol npe rewards account settings data
// swagger:model LolNpeRewardsAccountSettingsData
type LolNpeRewardsAccountSettingsData struct {

	// challenges
	Challenges *LolNpeRewardsChallengesSettings `json:"challenges,omitempty"`

	// login
	Login *LolNpeRewardsLoginSeriesSettings `json:"login,omitempty"`
}

// Validate validates this lol npe rewards account settings data
func (m *LolNpeRewardsAccountSettingsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChallenges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolNpeRewardsAccountSettingsData) validateChallenges(formats strfmt.Registry) error {

	if swag.IsZero(m.Challenges) { // not required
		return nil
	}

	if m.Challenges != nil {
		if err := m.Challenges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("challenges")
			}
			return err
		}
	}

	return nil
}

func (m *LolNpeRewardsAccountSettingsData) validateLogin(formats strfmt.Registry) error {

	if swag.IsZero(m.Login) { // not required
		return nil
	}

	if m.Login != nil {
		if err := m.Login.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeRewardsAccountSettingsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeRewardsAccountSettingsData) UnmarshalBinary(b []byte) error {
	var res LolNpeRewardsAccountSettingsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
