// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolGameflowGameflowGameData lol gameflow gameflow game data
// swagger:model LolGameflowGameflowGameData
type LolGameflowGameflowGameData struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// game name
	GameName string `json:"gameName,omitempty"`

	// is custom game
	IsCustomGame bool `json:"isCustomGame,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// player champion selections
	PlayerChampionSelections []interface{} `json:"playerChampionSelections"`

	// queue
	Queue *LolGameflowQueue `json:"queue,omitempty"`

	// spectators allowed
	SpectatorsAllowed bool `json:"spectatorsAllowed,omitempty"`

	// team one
	TeamOne []interface{} `json:"teamOne"`

	// team two
	TeamTwo []interface{} `json:"teamTwo"`
}

// Validate validates this lol gameflow gameflow game data
func (m *LolGameflowGameflowGameData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolGameflowGameflowGameData) validateQueue(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolGameflowGameflowGameData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGameflowGameflowGameData) UnmarshalBinary(b []byte) error {
	var res LolGameflowGameflowGameData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}