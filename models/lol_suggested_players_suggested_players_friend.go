// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolSuggestedPlayersSuggestedPlayersFriend lol suggested players suggested players friend
// swagger:model LolSuggestedPlayersSuggestedPlayersFriend
type LolSuggestedPlayersSuggestedPlayersFriend struct {

	// availability
	Availability string `json:"availability,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol suggested players suggested players friend
func (m *LolSuggestedPlayersSuggestedPlayersFriend) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersFriend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersFriend) UnmarshalBinary(b []byte) error {
	var res LolSuggestedPlayersSuggestedPlayersFriend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
