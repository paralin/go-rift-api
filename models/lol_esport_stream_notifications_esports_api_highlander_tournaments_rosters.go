// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRosters lol esport stream notifications esports API highlander tournaments rosters
// swagger:model LolEsportStreamNotificationsEsportsAPI_highlanderTournaments_rosters
type LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRosters struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// team
	Team int64 `json:"team,omitempty"`
}

// Validate validates this lol esport stream notifications esports API highlander tournaments rosters
func (m *LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRosters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRosters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRosters) UnmarshalBinary(b []byte) error {
	var res LolEsportStreamNotificationsEsportsAPIHighlanderTournamentsRosters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}