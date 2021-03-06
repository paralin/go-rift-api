// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EndOfGameLcdsTeamInfo end of game lcds team info
// swagger:model EndOfGameLcdsTeamInfo
type EndOfGameLcdsTeamInfo struct {

	// member status string
	MemberStatusString string `json:"memberStatusString,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// seconds until eligible for deletion
	SecondsUntilEligibleForDeletion int64 `json:"secondsUntilEligibleForDeletion,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// team Id
	TeamID *EndOfGameLcdsTeamID `json:"teamId,omitempty"`
}

// Validate validates this end of game lcds team info
func (m *EndOfGameLcdsTeamInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndOfGameLcdsTeamInfo) validateTeamID(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamID) { // not required
		return nil
	}

	if m.TeamID != nil {
		if err := m.TeamID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndOfGameLcdsTeamInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndOfGameLcdsTeamInfo) UnmarshalBinary(b []byte) error {
	var res EndOfGameLcdsTeamInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
