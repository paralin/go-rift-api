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

// LolLobbyTeamBuilderChampSelectSession lol lobby team builder champ select session
// swagger:model LolLobbyTeamBuilderChampSelectSession
type LolLobbyTeamBuilderChampSelectSession struct {

	// actions
	Actions []interface{} `json:"actions"`

	// allow battle boost
	AllowBattleBoost bool `json:"allowBattleBoost,omitempty"`

	// allow duplicate picks
	AllowDuplicatePicks bool `json:"allowDuplicatePicks,omitempty"`

	// allow locked events
	AllowLockedEvents bool `json:"allowLockedEvents,omitempty"`

	// allow rerolling
	AllowRerolling bool `json:"allowRerolling,omitempty"`

	// allow skin selection
	AllowSkinSelection bool `json:"allowSkinSelection,omitempty"`

	// bench champion ids
	BenchChampionIds []int32 `json:"benchChampionIds"`

	// bench enabled
	BenchEnabled bool `json:"benchEnabled,omitempty"`

	// boostable skin count
	BoostableSkinCount int32 `json:"boostableSkinCount,omitempty"`

	// chat details
	ChatDetails *LolLobbyTeamBuilderChampSelectChatRoomDetails `json:"chatDetails,omitempty"`

	// counter
	Counter int64 `json:"counter,omitempty"`

	// entitled feature state
	EntitledFeatureState *LolLobbyTeamBuilderEntitledFeatureState `json:"entitledFeatureState,omitempty"`

	// has simultaneous bans
	HasSimultaneousBans bool `json:"hasSimultaneousBans,omitempty"`

	// has simultaneous picks
	HasSimultaneousPicks bool `json:"hasSimultaneousPicks,omitempty"`

	// is spectating
	IsSpectating bool `json:"isSpectating,omitempty"`

	// local player cell Id
	LocalPlayerCellID int64 `json:"localPlayerCellId,omitempty"`

	// locked event index
	LockedEventIndex int32 `json:"lockedEventIndex,omitempty"`

	// my team
	MyTeam []*LolLobbyTeamBuilderChampSelectPlayerSelection `json:"myTeam"`

	// rerolls remaining
	RerollsRemaining int32 `json:"rerollsRemaining,omitempty"`

	// skip champion select
	SkipChampionSelect bool `json:"skipChampionSelect,omitempty"`

	// their team
	TheirTeam []*LolLobbyTeamBuilderChampSelectPlayerSelection `json:"theirTeam"`

	// timer
	Timer *LolLobbyTeamBuilderChampSelectTimer `json:"timer,omitempty"`

	// trades
	Trades []*LolLobbyTeamBuilderChampSelectTradeContract `json:"trades"`
}

// Validate validates this lol lobby team builder champ select session
func (m *LolLobbyTeamBuilderChampSelectSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChatDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitledFeatureState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMyTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTheirTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrades(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyTeamBuilderChampSelectSession) validateChatDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ChatDetails) { // not required
		return nil
	}

	if m.ChatDetails != nil {
		if err := m.ChatDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chatDetails")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyTeamBuilderChampSelectSession) validateEntitledFeatureState(formats strfmt.Registry) error {

	if swag.IsZero(m.EntitledFeatureState) { // not required
		return nil
	}

	if m.EntitledFeatureState != nil {
		if err := m.EntitledFeatureState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitledFeatureState")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyTeamBuilderChampSelectSession) validateMyTeam(formats strfmt.Registry) error {

	if swag.IsZero(m.MyTeam) { // not required
		return nil
	}

	for i := 0; i < len(m.MyTeam); i++ {
		if swag.IsZero(m.MyTeam[i]) { // not required
			continue
		}

		if m.MyTeam[i] != nil {
			if err := m.MyTeam[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("myTeam" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLobbyTeamBuilderChampSelectSession) validateTheirTeam(formats strfmt.Registry) error {

	if swag.IsZero(m.TheirTeam) { // not required
		return nil
	}

	for i := 0; i < len(m.TheirTeam); i++ {
		if swag.IsZero(m.TheirTeam[i]) { // not required
			continue
		}

		if m.TheirTeam[i] != nil {
			if err := m.TheirTeam[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("theirTeam" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLobbyTeamBuilderChampSelectSession) validateTimer(formats strfmt.Registry) error {

	if swag.IsZero(m.Timer) { // not required
		return nil
	}

	if m.Timer != nil {
		if err := m.Timer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timer")
			}
			return err
		}
	}

	return nil
}

func (m *LolLobbyTeamBuilderChampSelectSession) validateTrades(formats strfmt.Registry) error {

	if swag.IsZero(m.Trades) { // not required
		return nil
	}

	for i := 0; i < len(m.Trades); i++ {
		if swag.IsZero(m.Trades[i]) { // not required
			continue
		}

		if m.Trades[i] != nil {
			if err := m.Trades[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderChampSelectSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderChampSelectSession) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderChampSelectSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
