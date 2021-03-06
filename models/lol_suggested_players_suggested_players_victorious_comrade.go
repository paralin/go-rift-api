// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolSuggestedPlayersSuggestedPlayersVictoriousComrade lol suggested players suggested players victorious comrade
// swagger:model LolSuggestedPlayersSuggestedPlayersVictoriousComrade
type LolSuggestedPlayersSuggestedPlayersVictoriousComrade struct {

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`
}

// Validate validates this lol suggested players suggested players victorious comrade
func (m *LolSuggestedPlayersSuggestedPlayersVictoriousComrade) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersVictoriousComrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersVictoriousComrade) UnmarshalBinary(b []byte) error {
	var res LolSuggestedPlayersSuggestedPlayersVictoriousComrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
