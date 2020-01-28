// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathRewardStrategy lol npe tutorial path reward strategy
// swagger:model LolNpeTutorialPathRewardStrategy
type LolNpeTutorialPathRewardStrategy struct {

	// group strategy
	GroupStrategy string `json:"groupStrategy,omitempty"`

	// select max group count
	SelectMaxGroupCount int64 `json:"selectMaxGroupCount,omitempty"`

	// select min group count
	SelectMinGroupCount int64 `json:"selectMinGroupCount,omitempty"`
}

// Validate validates this lol npe tutorial path reward strategy
func (m *LolNpeTutorialPathRewardStrategy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathRewardStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathRewardStrategy) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathRewardStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
