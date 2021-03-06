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

// LolLobbyTeamBuilderMatchmakingSearchResource lol lobby team builder matchmaking search resource
// swagger:model LolLobbyTeamBuilderMatchmakingSearchResource
type LolLobbyTeamBuilderMatchmakingSearchResource struct {

	// dodge data
	DodgeData *LolLobbyTeamBuilderMatchmakingDodgeData `json:"dodgeData,omitempty"`

	// errors
	Errors []*LolLobbyTeamBuilderMatchmakingSearchErrorResource `json:"errors"`

	// estimated queue time
	EstimatedQueueTime float32 `json:"estimatedQueueTime,omitempty"`

	// is currently in queue
	IsCurrentlyInQueue bool `json:"isCurrentlyInQueue,omitempty"`

	// lobby Id
	LobbyID string `json:"lobbyId,omitempty"`

	// low priority data
	LowPriorityData *LolLobbyTeamBuilderMatchmakingLowPriorityData `json:"lowPriorityData,omitempty"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`

	// ready check
	ReadyCheck *LolLobbyTeamBuilderMatchmakingReadyCheckResource `json:"readyCheck,omitempty"`

	// search state
	SearchState LolLobbyTeamBuilderMatchmakingSearchState `json:"searchState,omitempty"`

	// time in queue
	TimeInQueue float32 `json:"timeInQueue,omitempty"`
}

// Validate validates this lol lobby team builder matchmaking search resource
func (m *LolLobbyTeamBuilderMatchmakingSearchResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDodgeData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLowPriorityData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadyCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyTeamBuilderMatchmakingSearchResource) validateDodgeData(formats strfmt.Registry) error {

	if swag.IsZero(m.DodgeData) { // not required
		return nil
	}

	if m.DodgeData != nil {
		if err := m.DodgeData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dodgeData")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyTeamBuilderMatchmakingSearchResource) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLobbyTeamBuilderMatchmakingSearchResource) validateLowPriorityData(formats strfmt.Registry) error {

	if swag.IsZero(m.LowPriorityData) { // not required
		return nil
	}

	if m.LowPriorityData != nil {
		if err := m.LowPriorityData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lowPriorityData")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyTeamBuilderMatchmakingSearchResource) validateReadyCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.ReadyCheck) { // not required
		return nil
	}

	if m.ReadyCheck != nil {
		if err := m.ReadyCheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readyCheck")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyTeamBuilderMatchmakingSearchResource) validateSearchState(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchState) { // not required
		return nil
	}

	if err := m.SearchState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("searchState")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderMatchmakingSearchResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderMatchmakingSearchResource) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderMatchmakingSearchResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
