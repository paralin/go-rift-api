// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClubsPublicClubTag lol clubs public club tag
// swagger:model LolClubsPublicClubTag
type LolClubsPublicClubTag struct {

	// club name
	ClubName string `json:"clubName,omitempty"`

	// club tag
	ClubTag string `json:"clubTag,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol clubs public club tag
func (m *LolClubsPublicClubTag) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsPublicClubTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsPublicClubTag) UnmarshalBinary(b []byte) error {
	var res LolClubsPublicClubTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}