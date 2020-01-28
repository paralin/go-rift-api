// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRankedSplitPointsNotification lol ranked split points notification
// swagger:model LolRankedSplitPointsNotification
type LolRankedSplitPointsNotification struct {

	// next reward Id
	NextRewardID string `json:"nextRewardId,omitempty"`

	// next reward type
	NextRewardType string `json:"nextRewardType,omitempty"`

	// previous split points required
	PreviousSplitPointsRequired int32 `json:"previousSplitPointsRequired,omitempty"`

	// split points after game
	SplitPointsAfterGame int32 `json:"splitPointsAfterGame,omitempty"`

	// split points before game
	SplitPointsBeforeGame int32 `json:"splitPointsBeforeGame,omitempty"`

	// split points breakdown
	SplitPointsBreakdown map[string]int32 `json:"splitPointsBreakdown,omitempty"`

	// split points delta
	SplitPointsDelta int32 `json:"splitPointsDelta,omitempty"`

	// split points required
	SplitPointsRequired int32 `json:"splitPointsRequired,omitempty"`
}

// Validate validates this lol ranked split points notification
func (m *LolRankedSplitPointsNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRankedSplitPointsNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRankedSplitPointsNotification) UnmarshalBinary(b []byte) error {
	var res LolRankedSplitPointsNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
