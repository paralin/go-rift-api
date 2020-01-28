// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeRewardsProgress lol npe rewards progress
// swagger:model LolNpeRewardsProgress
type LolNpeRewardsProgress struct {

	// current progress
	CurrentProgress int32 `json:"currentProgress,omitempty"`

	// last viewed progress
	LastViewedProgress int32 `json:"lastViewedProgress,omitempty"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this lol npe rewards progress
func (m *LolNpeRewardsProgress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeRewardsProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeRewardsProgress) UnmarshalBinary(b []byte) error {
	var res LolNpeRewardsProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
