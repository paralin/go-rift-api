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

// LolEsportStreamNotificationsLiveMatch lol esport stream notifications live match
// swagger:model LolEsportStreamNotificationsLiveMatch
type LolEsportStreamNotificationsLiveMatch struct {

	// id
	ID string `json:"id,omitempty"`

	// stream group
	StreamGroup string `json:"streamGroup,omitempty"`

	// teams
	Teams []*LolEsportStreamNotificationsLiveMatchTeam `json:"teams"`

	// title
	Title string `json:"title,omitempty"`

	// tournament description
	TournamentDescription string `json:"tournamentDescription,omitempty"`
}

// Validate validates this lol esport stream notifications live match
func (m *LolEsportStreamNotificationsLiveMatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolEsportStreamNotificationsLiveMatch) validateTeams(formats strfmt.Registry) error {

	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolEsportStreamNotificationsLiveMatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEsportStreamNotificationsLiveMatch) UnmarshalBinary(b []byte) error {
	var res LolEsportStreamNotificationsLiveMatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}