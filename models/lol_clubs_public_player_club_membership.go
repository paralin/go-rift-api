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

// LolClubsPublicPlayerClubMembership lol clubs public player club membership
// swagger:model LolClubsPublicPlayerClubMembership
type LolClubsPublicPlayerClubMembership struct {

	// active clubs
	ActiveClubs []*LolClubsPublicPlayerClub `json:"activeClubs"`

	// info
	Info *LolClubsPublicClubPlayer `json:"info,omitempty"`

	// preferences
	Preferences *LolClubsPublicClubPreferences `json:"preferences,omitempty"`

	// secure club presence info string
	SecureClubPresenceInfoString string `json:"secureClubPresenceInfoString,omitempty"`
}

// Validate validates this lol clubs public player club membership
func (m *LolClubsPublicPlayerClubMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveClubs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClubsPublicPlayerClubMembership) validateActiveClubs(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveClubs) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveClubs); i++ {
		if swag.IsZero(m.ActiveClubs[i]) { // not required
			continue
		}

		if m.ActiveClubs[i] != nil {
			if err := m.ActiveClubs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activeClubs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClubsPublicPlayerClubMembership) validateInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.Info) { // not required
		return nil
	}

	if m.Info != nil {
		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

func (m *LolClubsPublicPlayerClubMembership) validatePreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	if m.Preferences != nil {
		if err := m.Preferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preferences")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsPublicPlayerClubMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsPublicPlayerClubMembership) UnmarshalBinary(b []byte) error {
	var res LolClubsPublicPlayerClubMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}