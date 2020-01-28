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

// LolClubsPlayerClubMembership lol clubs player club membership
// swagger:model LolClubsPlayerClubMembership
type LolClubsPlayerClubMembership struct {

	// active clubs
	ActiveClubs []*LolClubsPlayerClub `json:"activeClubs"`

	// clubs server config
	ClubsServerConfig *LolClubsClubsConfig `json:"clubsServerConfig,omitempty"`

	// info
	Info *LolClubsClubPlayer `json:"info,omitempty"`

	// pending invites
	PendingInvites []*LolClubsClubInvite `json:"pendingInvites"`

	// preferences
	Preferences *LolClubsClubPreferences `json:"preferences,omitempty"`

	// removed clubs
	RemovedClubs []*LolClubsClub `json:"removedClubs"`

	// revoked invite clubs
	RevokedInviteClubs []*LolClubsClub `json:"revokedInviteClubs"`

	// secure club presence info string
	SecureClubPresenceInfoString string `json:"secureClubPresenceInfoString,omitempty"`
}

// Validate validates this lol clubs player club membership
func (m *LolClubsPlayerClubMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveClubs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClubsServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingInvites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemovedClubs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevokedInviteClubs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClubsPlayerClubMembership) validateActiveClubs(formats strfmt.Registry) error {

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

func (m *LolClubsPlayerClubMembership) validateClubsServerConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ClubsServerConfig) { // not required
		return nil
	}

	if m.ClubsServerConfig != nil {
		if err := m.ClubsServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clubsServerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *LolClubsPlayerClubMembership) validateInfo(formats strfmt.Registry) error {

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

func (m *LolClubsPlayerClubMembership) validatePendingInvites(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingInvites) { // not required
		return nil
	}

	for i := 0; i < len(m.PendingInvites); i++ {
		if swag.IsZero(m.PendingInvites[i]) { // not required
			continue
		}

		if m.PendingInvites[i] != nil {
			if err := m.PendingInvites[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pendingInvites" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClubsPlayerClubMembership) validatePreferences(formats strfmt.Registry) error {

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

func (m *LolClubsPlayerClubMembership) validateRemovedClubs(formats strfmt.Registry) error {

	if swag.IsZero(m.RemovedClubs) { // not required
		return nil
	}

	for i := 0; i < len(m.RemovedClubs); i++ {
		if swag.IsZero(m.RemovedClubs[i]) { // not required
			continue
		}

		if m.RemovedClubs[i] != nil {
			if err := m.RemovedClubs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("removedClubs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClubsPlayerClubMembership) validateRevokedInviteClubs(formats strfmt.Registry) error {

	if swag.IsZero(m.RevokedInviteClubs) { // not required
		return nil
	}

	for i := 0; i < len(m.RevokedInviteClubs); i++ {
		if swag.IsZero(m.RevokedInviteClubs[i]) { // not required
			continue
		}

		if m.RevokedInviteClubs[i] != nil {
			if err := m.RevokedInviteClubs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("revokedInviteClubs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsPlayerClubMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsPlayerClubMembership) UnmarshalBinary(b []byte) error {
	var res LolClubsPlayerClubMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}