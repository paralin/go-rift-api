// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolMatchHistoryRecentlyPlayedSummoner lol match history recently played summoner
// swagger:model LolMatchHistoryRecentlyPlayedSummoner
type LolMatchHistoryRecentlyPlayedSummoner struct {

	// champion Id
	ChampionID int64 `json:"championId,omitempty"`

	// game creation date
	GameCreationDate string `json:"gameCreationDate,omitempty"`

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`

	// team Id
	TeamID int64 `json:"teamId,omitempty"`
}

// Validate validates this lol match history recently played summoner
func (m *LolMatchHistoryRecentlyPlayedSummoner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolMatchHistoryRecentlyPlayedSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMatchHistoryRecentlyPlayedSummoner) UnmarshalBinary(b []byte) error {
	var res LolMatchHistoryRecentlyPlayedSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}