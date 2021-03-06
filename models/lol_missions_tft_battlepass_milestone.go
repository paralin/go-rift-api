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

// LolMissionsTftBattlepassMilestone lol missions tft battlepass milestone
// swagger:model LolMissionsTftBattlepassMilestone
type LolMissionsTftBattlepassMilestone struct {

	// description
	Description string `json:"description,omitempty"`

	// icon image Url
	IconImageURL string `json:"iconImageUrl,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// mission Id
	MissionID string `json:"missionId,omitempty"`

	// percent complete
	PercentComplete int64 `json:"percentComplete,omitempty"`

	// points for milestone
	PointsForMilestone int32 `json:"pointsForMilestone,omitempty"`

	// rewards
	Rewards []*PlayerMissionRewardDTO `json:"rewards"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this lol missions tft battlepass milestone
func (m *LolMissionsTftBattlepassMilestone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolMissionsTftBattlepassMilestone) validateRewards(formats strfmt.Registry) error {

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
func (m *LolMissionsTftBattlepassMilestone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMissionsTftBattlepassMilestone) UnmarshalBinary(b []byte) error {
	var res LolMissionsTftBattlepassMilestone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
