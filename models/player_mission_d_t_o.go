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

// PlayerMissionDTO player mission d t o
// swagger:model PlayerMissionDTO
type PlayerMissionDTO struct {

	// background image Url
	BackgroundImageURL string `json:"backgroundImageUrl,omitempty"`

	// celebration type
	CelebrationType string `json:"celebrationType,omitempty"`

	// client notify level
	ClientNotifyLevel string `json:"clientNotifyLevel,omitempty"`

	// completed date
	CompletedDate int64 `json:"completedDate,omitempty"`

	// completion expression
	CompletionExpression string `json:"completionExpression,omitempty"`

	// cooldown time millis
	CooldownTimeMillis int64 `json:"cooldownTimeMillis,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display
	Display *MissionDisplay `json:"display,omitempty"`

	// display type
	DisplayType string `json:"displayType,omitempty"`

	// earned date
	EarnedDate int64 `json:"earnedDate,omitempty"`

	// end time
	EndTime int64 `json:"endTime,omitempty"`

	// expiring warnings
	ExpiringWarnings []*MissionAlertDTO `json:"expiringWarnings"`

	// helper text
	HelperText string `json:"helperText,omitempty"`

	// icon image Url
	IconImageURL string `json:"iconImageUrl,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// internal name
	InternalName string `json:"internalName,omitempty"`

	// is new
	IsNew bool `json:"isNew,omitempty"`

	// last updated timestamp
	LastUpdatedTimestamp int64 `json:"lastUpdatedTimestamp,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`

	// metadata
	Metadata *MissionMetadata `json:"metadata,omitempty"`

	// mission type
	MissionType string `json:"missionType,omitempty"`

	// objectives
	Objectives []*PlayerMissionObjectiveDTO `json:"objectives"`

	// requirements
	Requirements []string `json:"requirements"`

	// reward strategy
	RewardStrategy *RewardStrategy `json:"rewardStrategy,omitempty"`

	// rewards
	Rewards []*PlayerMissionRewardDTO `json:"rewards"`

	// sequence
	Sequence int32 `json:"sequence,omitempty"`

	// series name
	SeriesName string `json:"seriesName,omitempty"`

	// start time
	StartTime int64 `json:"startTime,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// viewed
	Viewed bool `json:"viewed,omitempty"`
}

// Validate validates this player mission d t o
func (m *PlayerMissionDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiringWarnings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayerMissionDTO) validateDisplay(formats strfmt.Registry) error {

	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if m.Display != nil {
		if err := m.Display.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerMissionDTO) validateExpiringWarnings(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiringWarnings) { // not required
		return nil
	}

	for i := 0; i < len(m.ExpiringWarnings); i++ {
		if swag.IsZero(m.ExpiringWarnings[i]) { // not required
			continue
		}

		if m.ExpiringWarnings[i] != nil {
			if err := m.ExpiringWarnings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("expiringWarnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlayerMissionDTO) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerMissionDTO) validateObjectives(formats strfmt.Registry) error {

	if swag.IsZero(m.Objectives) { // not required
		return nil
	}

	for i := 0; i < len(m.Objectives); i++ {
		if swag.IsZero(m.Objectives[i]) { // not required
			continue
		}

		if m.Objectives[i] != nil {
			if err := m.Objectives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlayerMissionDTO) validateRewardStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.RewardStrategy) { // not required
		return nil
	}

	if m.RewardStrategy != nil {
		if err := m.RewardStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rewardStrategy")
			}
			return err
		}
	}

	return nil
}

func (m *PlayerMissionDTO) validateRewards(formats strfmt.Registry) error {

	if swag.IsZero(m.Rewards) { // not required
		return nil
	}

	for i := 0; i < len(m.Rewards); i++ {
		if swag.IsZero(m.Rewards[i]) { // not required
			continue
		}

		if m.Rewards[i] != nil {
			if err := m.Rewards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayerMissionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayerMissionDTO) UnmarshalBinary(b []byte) error {
	var res PlayerMissionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
