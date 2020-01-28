// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LcdsPracticeGameSearchResult lcds practice game search result
// swagger:model LcdsPracticeGameSearchResult
type LcdsPracticeGameSearchResult struct {

	// allow spectators
	AllowSpectators string `json:"allowSpectators,omitempty"`

	// game map
	GameMap *LcdsGameMap `json:"gameMap,omitempty"`

	// game map Id
	GameMapID int32 `json:"gameMapId,omitempty"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// max num players
	MaxNumPlayers int32 `json:"maxNumPlayers,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner *LcdsPlayerParticipant `json:"owner,omitempty"`

	// pick type
	PickType string `json:"pickType,omitempty"`

	// private game
	PrivateGame bool `json:"privateGame,omitempty"`

	// spectator count
	SpectatorCount int32 `json:"spectatorCount,omitempty"`

	// team1 count
	Team1Count int32 `json:"team1Count,omitempty"`

	// team2 count
	Team2Count int32 `json:"team2Count,omitempty"`
}

// Validate validates this lcds practice game search result
func (m *LcdsPracticeGameSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LcdsPracticeGameSearchResult) validateGameMap(formats strfmt.Registry) error {

	if swag.IsZero(m.GameMap) { // not required
		return nil
	}

	if m.GameMap != nil {
		if err := m.GameMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameMap")
			}
			return err
		}
	}

	return nil
}

func (m *LcdsPracticeGameSearchResult) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LcdsPracticeGameSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LcdsPracticeGameSearchResult) UnmarshalBinary(b []byte) error {
	var res LcdsPracticeGameSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
