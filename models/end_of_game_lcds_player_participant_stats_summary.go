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

// EndOfGameLcdsPlayerParticipantStatsSummary end of game lcds player participant stats summary
// swagger:model EndOfGameLcdsPlayerParticipantStatsSummary
type EndOfGameLcdsPlayerParticipantStatsSummary struct {

	// bot player
	BotPlayer bool `json:"botPlayer,omitempty"`

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// detected team position
	DetectedTeamPosition string `json:"detectedTeamPosition,omitempty"`

	// elo
	Elo int32 `json:"elo,omitempty"`

	// elo change
	EloChange int32 `json:"eloChange,omitempty"`

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// leaver
	Leaver bool `json:"leaver,omitempty"`

	// leaves
	Leaves int32 `json:"leaves,omitempty"`

	// level
	Level int32 `json:"level,omitempty"`

	// losses
	Losses int32 `json:"losses,omitempty"`

	// profile icon Id
	ProfileIconID int32 `json:"profileIconId,omitempty"`

	// selected position
	SelectedPosition string `json:"selectedPosition,omitempty"`

	// skin index
	SkinIndex int32 `json:"skinIndex,omitempty"`

	// skin name
	SkinName string `json:"skinName,omitempty"`

	// spell1 Id
	Spell1ID int32 `json:"spell1Id,omitempty"`

	// spell2 Id
	Spell2ID int32 `json:"spell2Id,omitempty"`

	// statistics
	Statistics []*EndOfGameLcdsRawStatDTO `json:"statistics"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`

	// team Id
	TeamID int32 `json:"teamId,omitempty"`

	// user Id
	UserID int64 `json:"userId,omitempty"`

	// wins
	Wins int32 `json:"wins,omitempty"`
}

// Validate validates this end of game lcds player participant stats summary
func (m *EndOfGameLcdsPlayerParticipantStatsSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndOfGameLcdsPlayerParticipantStatsSummary) validateStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	for i := 0; i < len(m.Statistics); i++ {
		if swag.IsZero(m.Statistics[i]) { // not required
			continue
		}

		if m.Statistics[i] != nil {
			if err := m.Statistics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statistics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndOfGameLcdsPlayerParticipantStatsSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndOfGameLcdsPlayerParticipantStatsSummary) UnmarshalBinary(b []byte) error {
	var res EndOfGameLcdsPlayerParticipantStatsSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
