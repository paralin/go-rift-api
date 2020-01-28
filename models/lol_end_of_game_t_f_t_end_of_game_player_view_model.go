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

// LolEndOfGameTFTEndOfGamePlayerViewModel lol end of game t f t end of game player view model
// swagger:model LolEndOfGameTFTEndOfGamePlayerViewModel
type LolEndOfGameTFTEndOfGamePlayerViewModel struct {

	// board pieces
	BoardPieces []*LolEndOfGameTFTEndOfGamePieceViewModel `json:"boardPieces"`

	// companion
	Companion *LolEndOfGameTFTEndOfGameCompanionViewModel `json:"companion,omitempty"`

	// ffa standing
	FfaStanding int64 `json:"ffaStanding,omitempty"`

	// health
	Health int64 `json:"health,omitempty"`

	// icon Id
	IconID int32 `json:"iconId,omitempty"`

	// is local player
	IsLocalPlayer bool `json:"isLocalPlayer,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`

	// rank
	Rank int64 `json:"rank,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`
}

// Validate validates this lol end of game t f t end of game player view model
func (m *LolEndOfGameTFTEndOfGamePlayerViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoardPieces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompanion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolEndOfGameTFTEndOfGamePlayerViewModel) validateBoardPieces(formats strfmt.Registry) error {

	if swag.IsZero(m.BoardPieces) { // not required
		return nil
	}

	for i := 0; i < len(m.BoardPieces); i++ {
		if swag.IsZero(m.BoardPieces[i]) { // not required
			continue
		}

		if m.BoardPieces[i] != nil {
			if err := m.BoardPieces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boardPieces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolEndOfGameTFTEndOfGamePlayerViewModel) validateCompanion(formats strfmt.Registry) error {

	if swag.IsZero(m.Companion) { // not required
		return nil
	}

	if m.Companion != nil {
		if err := m.Companion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("companion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameTFTEndOfGamePlayerViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameTFTEndOfGamePlayerViewModel) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameTFTEndOfGamePlayerViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
