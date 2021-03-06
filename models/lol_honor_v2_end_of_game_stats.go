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

// LolHonorV2EndOfGameStats lol honor v2 end of game stats
// swagger:model LolHonorV2EndOfGameStats
type LolHonorV2EndOfGameStats struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// difficulty
	Difficulty string `json:"difficulty,omitempty"`

	// game ended in early surrender
	GameEndedInEarlySurrender bool `json:"gameEndedInEarlySurrender,omitempty"`

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// game length
	GameLength int32 `json:"gameLength,omitempty"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// game mutators
	GameMutators []string `json:"gameMutators"`

	// game type
	GameType string `json:"gameType,omitempty"`

	// imbalanced teams no points
	ImbalancedTeamsNoPoints bool `json:"imbalancedTeamsNoPoints,omitempty"`

	// invalid
	Invalid bool `json:"invalid,omitempty"`

	// my team status
	MyTeamStatus string `json:"myTeamStatus,omitempty"`

	// queue type
	QueueType string `json:"queueType,omitempty"`

	// ranked
	Ranked bool `json:"ranked,omitempty"`

	// report game Id
	ReportGameID int64 `json:"reportGameId,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`

	// teams
	Teams []*LolHonorV2EndOfGameTeam `json:"teams"`
}

// Validate validates this lol honor v2 end of game stats
func (m *LolHonorV2EndOfGameStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolHonorV2EndOfGameStats) validateTeams(formats strfmt.Registry) error {

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
func (m *LolHonorV2EndOfGameStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHonorV2EndOfGameStats) UnmarshalBinary(b []byte) error {
	var res LolHonorV2EndOfGameStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
