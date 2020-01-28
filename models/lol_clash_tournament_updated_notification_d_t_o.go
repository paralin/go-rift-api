// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashTournamentUpdatedNotificationDTO lol clash tournament updated notification d t o
// swagger:model LolClashTournamentUpdatedNotificationDTO
type LolClashTournamentUpdatedNotificationDTO struct {

	// reason
	Reason LolClashTournamentNotifyReason `json:"reason,omitempty"`

	// relevant phase Id
	RelevantPhaseID int64 `json:"relevantPhaseId,omitempty"`

	// tournament
	Tournament *TournamentDTO `json:"tournament,omitempty"`

	// tournament Id
	TournamentID int64 `json:"tournamentId,omitempty"`
}

// Validate validates this lol clash tournament updated notification d t o
func (m *LolClashTournamentUpdatedNotificationDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
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

func (m *LolClashTournamentUpdatedNotificationDTO) validateReason(formats strfmt.Registry) error {

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

func (m *LolClashTournamentUpdatedNotificationDTO) validateTournament(formats strfmt.Registry) error {

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
func (m *LolClashTournamentUpdatedNotificationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashTournamentUpdatedNotificationDTO) UnmarshalBinary(b []byte) error {
	var res LolClashTournamentUpdatedNotificationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}