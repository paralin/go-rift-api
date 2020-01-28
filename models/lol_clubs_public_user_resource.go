// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClubsPublicUserResource lol clubs public user resource
// swagger:model LolClubsPublicUserResource
type LolClubsPublicUserResource struct {

	// availability
	Availability string `json:"availability,omitempty"`

	// icon
	Icon int32 `json:"icon,omitempty"`

	// last seen online timestamp
	LastSeenOnlineTimestamp string `json:"lastSeenOnlineTimestamp,omitempty"`

	// lol
	Lol map[string]string `json:"lol,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol clubs public user resource
func (m *LolClubsPublicUserResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsPublicUserResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsPublicUserResource) UnmarshalBinary(b []byte) error {
	var res LolClubsPublicUserResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
