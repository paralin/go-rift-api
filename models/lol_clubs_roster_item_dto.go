// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClubsRosterItemDto lol clubs roster item dto
// swagger:model LolClubsRosterItemDto
type LolClubsRosterItemDto struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// club role
	ClubRole string `json:"clubRole,omitempty"`

	// summoner icon Id
	SummonerIconID int32 `json:"summonerIconId,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`
}

// Validate validates this lol clubs roster item dto
func (m *LolClubsRosterItemDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsRosterItemDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsRosterItemDto) UnmarshalBinary(b []byte) error {
	var res LolClubsRosterItemDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
