// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathTutorialReward lol npe tutorial path tutorial reward
// swagger:model LolNpeTutorialPathTutorialReward
type LolNpeTutorialPathTutorialReward struct {

	// description
	Description string `json:"description,omitempty"`

	// icon Url
	IconURL string `json:"iconUrl,omitempty"`

	// item Id
	ItemID string `json:"itemId,omitempty"`

	// reward fulfilled
	RewardFulfilled bool `json:"rewardFulfilled,omitempty"`

	// reward type
	RewardType string `json:"rewardType,omitempty"`

	// sequence
	Sequence int32 `json:"sequence,omitempty"`

	// unique name
	UniqueName string `json:"uniqueName,omitempty"`
}

// Validate validates this lol npe tutorial path tutorial reward
func (m *LolNpeTutorialPathTutorialReward) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathTutorialReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathTutorialReward) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathTutorialReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}