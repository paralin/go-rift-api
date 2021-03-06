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

// LolEndOfGameTFTEndOfGameViewModel lol end of game t f t end of game view model
// swagger:model LolEndOfGameTFTEndOfGameViewModel
type LolEndOfGameTFTEndOfGameViewModel struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// game length
	GameLength int32 `json:"gameLength,omitempty"`

	// is ranked
	IsRanked bool `json:"isRanked,omitempty"`

	// local player
	LocalPlayer *LolEndOfGameTFTEndOfGamePlayerViewModel `json:"localPlayer,omitempty"`

	// players
	Players []*LolEndOfGameTFTEndOfGamePlayerViewModel `json:"players"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`
}

// Validate validates this lol end of game t f t end of game view model
func (m *LolEndOfGameTFTEndOfGameViewModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalPlayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolEndOfGameTFTEndOfGameViewModel) validateLocalPlayer(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalPlayer) { // not required
		return nil
	}

	if m.LocalPlayer != nil {
		if err := m.LocalPlayer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localPlayer")
			}
			return err
		}
	}

	return nil
}

func (m *LolEndOfGameTFTEndOfGameViewModel) validatePlayers(formats strfmt.Registry) error {

	if swag.IsZero(m.Players) { // not required
		return nil
	}

	for i := 0; i < len(m.Players); i++ {
		if swag.IsZero(m.Players[i]) { // not required
			continue
		}

		if m.Players[i] != nil {
			if err := m.Players[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("players" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameTFTEndOfGameViewModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameTFTEndOfGameViewModel) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameTFTEndOfGameViewModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
