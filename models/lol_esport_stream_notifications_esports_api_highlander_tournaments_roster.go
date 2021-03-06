// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRoster lol esport stream notifications esports API highlander tournaments roster
// swagger:model LolEsportStreamNotificationsEsportsAPI_highlanderTournaments_roster
type LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRoster struct {

	// roster
	Roster string `json:"roster,omitempty"`
}

// Validate validates this lol esport stream notifications esports API highlander tournaments roster
func (m *LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRoster) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRoster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRoster) UnmarshalBinary(b []byte) error {
	var res LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRoster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
