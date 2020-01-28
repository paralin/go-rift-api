// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderQueueGameTypeConfig lol lobby team builder queue game type config
// swagger:model LolLobbyTeamBuilderQueueGameTypeConfig
type LolLobbyTeamBuilderQueueGameTypeConfig struct {

	// advanced learning quests
	AdvancedLearningQuests bool `json:"advancedLearningQuests,omitempty"`

	// allow trades
	AllowTrades bool `json:"allowTrades,omitempty"`

	// ban mode
	BanMode string `json:"banMode,omitempty"`

	// ban timer duration
	BanTimerDuration int32 `json:"banTimerDuration,omitempty"`

	// battle boost
	BattleBoost bool `json:"battleBoost,omitempty"`

	// cross team champion pool
	CrossTeamChampionPool bool `json:"crossTeamChampionPool,omitempty"`

	// death match
	DeathMatch bool `json:"deathMatch,omitempty"`

	// do not remove
	DoNotRemove bool `json:"doNotRemove,omitempty"`

	// duplicate pick
	DuplicatePick bool `json:"duplicatePick,omitempty"`

	// exclusive pick
	ExclusivePick bool `json:"exclusivePick,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// learning quests
	LearningQuests bool `json:"learningQuests,omitempty"`

	// main pick timer duration
	MainPickTimerDuration int32 `json:"mainPickTimerDuration,omitempty"`

	// max allowable bans
	MaxAllowableBans int32 `json:"maxAllowableBans,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// onboard coop beginner
	OnboardCoopBeginner bool `json:"onboardCoopBeginner,omitempty"`

	// pick mode
	PickMode string `json:"pickMode,omitempty"`

	// post pick timer duration
	PostPickTimerDuration int32 `json:"postPickTimerDuration,omitempty"`

	// reroll
	Reroll bool `json:"reroll,omitempty"`

	// team champion pool
	TeamChampionPool bool `json:"teamChampionPool,omitempty"`
}

// Validate validates this lol lobby team builder queue game type config
func (m *LolLobbyTeamBuilderQueueGameTypeConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderQueueGameTypeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderQueueGameTypeConfig) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderQueueGameTypeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}