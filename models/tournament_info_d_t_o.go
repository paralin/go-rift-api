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

// TournamentInfoDTO tournament info d t o
// swagger:model TournamentInfoDTO
type TournamentInfoDTO struct {

	// invitee pending rosters
	InviteePendingRosters []*PendingRosterDTO `json:"inviteePendingRosters"`

	// pending roster
	PendingRoster *PendingRosterDTO `json:"pendingRoster,omitempty"`

	// roster
	Roster *RosterDTO `json:"roster,omitempty"`

	// sub rosters
	SubRosters []*RosterDTO `json:"subRosters"`

	// theme vp
	ThemeVp int32 `json:"themeVp,omitempty"`

	// tournament
	Tournament *TournamentDTO `json:"tournament,omitempty"`
}

// Validate validates this tournament info d t o
func (m *TournamentInfoDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInviteePendingRosters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingRoster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRosters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTournament(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TournamentInfoDTO) validateInviteePendingRosters(formats strfmt.Registry) error {

	if swag.IsZero(m.InviteePendingRosters) { // not required
		return nil
	}

	for i := 0; i < len(m.InviteePendingRosters); i++ {
		if swag.IsZero(m.InviteePendingRosters[i]) { // not required
			continue
		}

		if m.InviteePendingRosters[i] != nil {
			if err := m.InviteePendingRosters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inviteePendingRosters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TournamentInfoDTO) validatePendingRoster(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingRoster) { // not required
		return nil
	}

	if m.PendingRoster != nil {
		if err := m.PendingRoster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pendingRoster")
			}
			return err
		}
	}

	return nil
}

func (m *TournamentInfoDTO) validateRoster(formats strfmt.Registry) error {

	if swag.IsZero(m.Roster) { // not required
		return nil
	}

	if m.Roster != nil {
		if err := m.Roster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roster")
			}
			return err
		}
	}

	return nil
}

func (m *TournamentInfoDTO) validateSubRosters(formats strfmt.Registry) error {

	if swag.IsZero(m.SubRosters) { // not required
		return nil
	}

	for i := 0; i < len(m.SubRosters); i++ {
		if swag.IsZero(m.SubRosters[i]) { // not required
			continue
		}

		if m.SubRosters[i] != nil {
			if err := m.SubRosters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subRosters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TournamentInfoDTO) validateTournament(formats strfmt.Registry) error {

	if swag.IsZero(m.Tournament) { // not required
		return nil
	}

	if m.Tournament != nil {
		if err := m.Tournament.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tournament")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TournamentInfoDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TournamentInfoDTO) UnmarshalBinary(b []byte) error {
	var res TournamentInfoDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}