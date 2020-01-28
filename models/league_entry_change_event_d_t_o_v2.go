// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LeagueEntryChangeEventDTOV2 league entry change event d t o v2
// swagger:model LeagueEntryChangeEventDTOV2
type LeagueEntryChangeEventDTOV2 struct {

	// league item
	LeagueItem *LeaguesLcdsLeagueItemDTO `json:"leagueItem,omitempty"`

	// league points delta
	LeaguePointsDelta int32 `json:"leaguePointsDelta,omitempty"`

	// lp change reason
	LpChangeReason string `json:"lpChangeReason,omitempty"`

	// notify reason
	NotifyReason string `json:"notifyReason,omitempty"`
}

// Validate validates this league entry change event d t o v2
func (m *LeagueEntryChangeEventDTOV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeagueItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LeagueEntryChangeEventDTOV2) validateLeagueItem(formats strfmt.Registry) error {

	if swag.IsZero(m.LeagueItem) { // not required
		return nil
	}

	if m.LeagueItem != nil {
		if err := m.LeagueItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leagueItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LeagueEntryChangeEventDTOV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeagueEntryChangeEventDTOV2) UnmarshalBinary(b []byte) error {
	var res LeagueEntryChangeEventDTOV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
