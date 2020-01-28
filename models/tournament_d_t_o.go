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

// TournamentDTO tournament d t o
// swagger:model TournamentDTO
type TournamentDTO struct {

	// bracket formation init delay ms
	BracketFormationInitDelayMs int64 `json:"bracketFormationInitDelayMs,omitempty"`

	// bracket formation interval ms
	BracketFormationIntervalMs int64 `json:"bracketFormationIntervalMs,omitempty"`

	// bracket size
	BracketSize string `json:"bracketSize,omitempty"`

	// buy in options
	BuyInOptions []int32 `json:"buyInOptions"`

	// buy in options premium
	BuyInOptionsPremium []int32 `json:"buyInOptionsPremium"`

	// entry fee
	EntryFee int32 `json:"entryFee,omitempty"`

	// honor restriction
	HonorRestriction bool `json:"honorRestriction,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// last theme of season
	LastThemeOfSeason bool `json:"lastThemeOfSeason,omitempty"`

	// max sub
	MaxSub int32 `json:"maxSub,omitempty"`

	// name loc key
	NameLocKey string `json:"nameLocKey,omitempty"`

	// name loc key secondary
	NameLocKeySecondary string `json:"nameLocKeySecondary,omitempty"`

	// phases
	Phases []*TournamentPhaseDTO `json:"phases"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`

	// rank restriction
	RankRestriction bool `json:"rankRestriction,omitempty"`

	// resume time
	ResumeTime int64 `json:"resumeTime,omitempty"`

	// reward config
	RewardConfig []*ClashRewardConfigClient `json:"rewardConfig"`

	// roster create deadline
	RosterCreateDeadline int64 `json:"rosterCreateDeadline,omitempty"`

	// roster size
	RosterSize int32 `json:"rosterSize,omitempty"`

	// schedule end time
	ScheduleEndTime int64 `json:"scheduleEndTime,omitempty"`

	// schedule time
	ScheduleTime int64 `json:"scheduleTime,omitempty"`

	// scouting time ms
	ScoutingTimeMs int64 `json:"scoutingTimeMs,omitempty"`

	// sms restriction
	SmsRestriction bool `json:"smsRestriction,omitempty"`

	// status
	Status TournamentStatusEnum `json:"status,omitempty"`

	// theme Id
	ThemeID int32 `json:"themeId,omitempty"`

	// tier configs
	TierConfigs []*TierConfig `json:"tierConfigs"`

	// voice enabled
	VoiceEnabled bool `json:"voiceEnabled,omitempty"`
}

// Validate validates this tournament d t o
func (m *TournamentDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TournamentDTO) validatePhases(formats strfmt.Registry) error {

	if swag.IsZero(m.Phases) { // not required
		return nil
	}

	for i := 0; i < len(m.Phases); i++ {
		if swag.IsZero(m.Phases[i]) { // not required
			continue
		}

		if m.Phases[i] != nil {
			if err := m.Phases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TournamentDTO) validateRewardConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.RewardConfig) { // not required
		return nil
	}

	for i := 0; i < len(m.RewardConfig); i++ {
		if swag.IsZero(m.RewardConfig[i]) { // not required
			continue
		}

		if m.RewardConfig[i] != nil {
			if err := m.RewardConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rewardConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TournamentDTO) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *TournamentDTO) validateTierConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.TierConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.TierConfigs); i++ {
		if swag.IsZero(m.TierConfigs[i]) { // not required
			continue
		}

		if m.TierConfigs[i] != nil {
			if err := m.TierConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tierConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TournamentDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TournamentDTO) UnmarshalBinary(b []byte) error {
	var res TournamentDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
