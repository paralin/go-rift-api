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

// LolClashTournament lol clash tournament
// swagger:model LolClashTournament
type LolClashTournament struct {

	// allow roster creation
	AllowRosterCreation bool `json:"allowRosterCreation,omitempty"`

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

	// end time ms
	EndTimeMs int64 `json:"endTimeMs,omitempty"`

	// entry fee
	EntryFee int32 `json:"entryFee,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is honor restriction enabled
	IsHonorRestrictionEnabled bool `json:"isHonorRestrictionEnabled,omitempty"`

	// is ranked restriction enabled
	IsRankedRestrictionEnabled bool `json:"isRankedRestrictionEnabled,omitempty"`

	// is sms restriction enabled
	IsSmsRestrictionEnabled bool `json:"isSmsRestrictionEnabled,omitempty"`

	// last theme of season
	LastThemeOfSeason bool `json:"lastThemeOfSeason,omitempty"`

	// max substitutes
	MaxSubstitutes int32 `json:"maxSubstitutes,omitempty"`

	// name loc key
	NameLocKey string `json:"nameLocKey,omitempty"`

	// name loc key secondary
	NameLocKeySecondary string `json:"nameLocKeySecondary,omitempty"`

	// phases
	Phases []*LolClashTournamentPhase `json:"phases"`

	// resume time
	ResumeTime int64 `json:"resumeTime,omitempty"`

	// reward config
	RewardConfig []*ClashRewardConfigClient `json:"rewardConfig"`

	// roster create deadline
	RosterCreateDeadline int64 `json:"rosterCreateDeadline,omitempty"`

	// roster size
	RosterSize int32 `json:"rosterSize,omitempty"`

	// scouting duration ms
	ScoutingDurationMs int64 `json:"scoutingDurationMs,omitempty"`

	// start time ms
	StartTimeMs int64 `json:"startTimeMs,omitempty"`

	// status
	Status TournamentStatusEnum `json:"status,omitempty"`

	// theme Id
	ThemeID int32 `json:"themeId,omitempty"`

	// tier configs
	TierConfigs []*TierConfig `json:"tierConfigs"`
}

// Validate validates this lol clash tournament
func (m *LolClashTournament) Validate(formats strfmt.Registry) error {
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

func (m *LolClashTournament) validatePhases(formats strfmt.Registry) error {

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

func (m *LolClashTournament) validateRewardConfig(formats strfmt.Registry) error {

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

func (m *LolClashTournament) validateStatus(formats strfmt.Registry) error {

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

func (m *LolClashTournament) validateTierConfigs(formats strfmt.Registry) error {

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
func (m *LolClashTournament) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashTournament) UnmarshalBinary(b []byte) error {
	var res LolClashTournament
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
