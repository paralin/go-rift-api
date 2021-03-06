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

// PublicRosterDTO public roster d t o
// swagger:model PublicRosterDTO
type PublicRosterDTO struct {

	// id
	ID int64 `json:"id,omitempty"`

	// logo
	Logo int32 `json:"logo,omitempty"`

	// logo color
	LogoColor int32 `json:"logoColor,omitempty"`

	// member ids
	MemberIds []int64 `json:"memberIds"`

	// name
	Name string `json:"name,omitempty"`

	// phases subs
	PhasesSubs []*PublicPhaseSubsDTO `json:"phasesSubs"`

	// short name
	ShortName string `json:"shortName,omitempty"`

	// tournament Id
	TournamentID int64 `json:"tournamentId,omitempty"`
}

// Validate validates this public roster d t o
func (m *PublicRosterDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhasesSubs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicRosterDTO) validatePhasesSubs(formats strfmt.Registry) error {

	if swag.IsZero(m.PhasesSubs) { // not required
		return nil
	}

	for i := 0; i < len(m.PhasesSubs); i++ {
		if swag.IsZero(m.PhasesSubs[i]) { // not required
			continue
		}

		if m.PhasesSubs[i] != nil {
			if err := m.PhasesSubs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phasesSubs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicRosterDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicRosterDTO) UnmarshalBinary(b []byte) error {
	var res PublicRosterDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
