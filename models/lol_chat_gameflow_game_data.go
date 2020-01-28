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

// LolChatGameflowGameData lol chat gameflow game data
// swagger:model LolChatGameflowGameData
type LolChatGameflowGameData struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// player champion selections
	PlayerChampionSelections []*LolChatChampSelection `json:"playerChampionSelections"`

	// queue
	Queue *LolChatQueue `json:"queue,omitempty"`

	// team one
	TeamOne []*LolChatTeamPlayerEntry `json:"teamOne"`

	// team two
	TeamTwo []*LolChatTeamPlayerEntry `json:"teamTwo"`
}

// Validate validates this lol chat gameflow game data
func (m *LolChatGameflowGameData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayerChampionSelections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamOne(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamTwo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChatGameflowGameData) validatePlayerChampionSelections(formats strfmt.Registry) error {

	if swag.IsZero(m.PlayerChampionSelections) { // not required
		return nil
	}

	for i := 0; i < len(m.PlayerChampionSelections); i++ {
		if swag.IsZero(m.PlayerChampionSelections[i]) { // not required
			continue
		}

		if m.PlayerChampionSelections[i] != nil {
			if err := m.PlayerChampionSelections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("playerChampionSelections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolChatGameflowGameData) validateQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *LolChatGameflowGameData) validateTeamOne(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamOne) { // not required
		return nil
	}

	for i := 0; i < len(m.TeamOne); i++ {
		if swag.IsZero(m.TeamOne[i]) { // not required
			continue
		}

		if m.TeamOne[i] != nil {
			if err := m.TeamOne[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teamOne" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolChatGameflowGameData) validateTeamTwo(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamTwo) { // not required
		return nil
	}

	for i := 0; i < len(m.TeamTwo); i++ {
		if swag.IsZero(m.TeamTwo[i]) { // not required
			continue
		}

		if m.TeamTwo[i] != nil {
			if err := m.TeamTwo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teamTwo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChatGameflowGameData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatGameflowGameData) UnmarshalBinary(b []byte) error {
	var res LolChatGameflowGameData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
