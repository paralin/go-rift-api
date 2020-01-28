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

// LolGameQueuesQueueCustomGameSubcategory lol game queues queue custom game subcategory
// swagger:model LolGameQueuesQueueCustomGameSubcategory
type LolGameQueuesQueueCustomGameSubcategory struct {

	// custom spectator policies
	CustomSpectatorPolicies []LolGameQueuesQueueCustomGameSpectatorPolicy `json:"customSpectatorPolicies"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// map Id
	MapID int32 `json:"mapId,omitempty"`

	// max level
	MaxLevel int32 `json:"maxLevel,omitempty"`

	// max player count
	MaxPlayerCount int32 `json:"maxPlayerCount,omitempty"`

	// maximum participant list size
	MaximumParticipantListSize int32 `json:"maximumParticipantListSize,omitempty"`

	// min level
	MinLevel int32 `json:"minLevel,omitempty"`

	// minimum participant list size
	MinimumParticipantListSize int32 `json:"minimumParticipantListSize,omitempty"`

	// mutators
	Mutators []*LolGameQueuesQueueGameTypeConfig `json:"mutators"`

	// num players per team
	NumPlayersPerTeam int32 `json:"numPlayersPerTeam,omitempty"`

	// queue availability
	QueueAvailability LolGameQueuesQueueAvailability `json:"queueAvailability,omitempty"`
}

// Validate validates this lol game queues queue custom game subcategory
func (m *LolGameQueuesQueueCustomGameSubcategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomSpectatorPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMutators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueAvailability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolGameQueuesQueueCustomGameSubcategory) validateCustomSpectatorPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomSpectatorPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomSpectatorPolicies); i++ {

		if err := m.CustomSpectatorPolicies[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customSpectatorPolicies" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *LolGameQueuesQueueCustomGameSubcategory) validateMutators(formats strfmt.Registry) error {

	if swag.IsZero(m.Mutators) { // not required
		return nil
	}

	for i := 0; i < len(m.Mutators); i++ {
		if swag.IsZero(m.Mutators[i]) { // not required
			continue
		}

		if m.Mutators[i] != nil {
			if err := m.Mutators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mutators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolGameQueuesQueueCustomGameSubcategory) validateQueueAvailability(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueAvailability) { // not required
		return nil
	}

	if err := m.QueueAvailability.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("queueAvailability")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolGameQueuesQueueCustomGameSubcategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGameQueuesQueueCustomGameSubcategory) UnmarshalBinary(b []byte) error {
	var res LolGameQueuesQueueCustomGameSubcategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}