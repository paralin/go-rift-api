// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClashSuggestedInvite lol clash suggested invite
// swagger:model LolClashSuggestedInvite
type LolClashSuggestedInvite struct {

	// suggester summoner Id
	SuggesterSummonerID int64 `json:"suggesterSummonerId,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol clash suggested invite
func (m *LolClashSuggestedInvite) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClashSuggestedInvite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashSuggestedInvite) UnmarshalBinary(b []byte) error {
	var res LolClashSuggestedInvite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}