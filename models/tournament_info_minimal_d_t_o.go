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

// TournamentInfoMinimalDTO tournament info minimal d t o
// swagger:model TournamentInfoMinimalDTO
type TournamentInfoMinimalDTO struct {

	// time
	Time int64 `json:"time,omitempty"`

	// tournament info
	TournamentInfo []*TournamentInfoDTO `json:"tournamentInfo"`
}

// Validate validates this tournament info minimal d t o
func (m *TournamentInfoMinimalDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTournamentInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TournamentInfoMinimalDTO) validateTournamentInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.TournamentInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.TournamentInfo); i++ {
		if swag.IsZero(m.TournamentInfo[i]) { // not required
			continue
		}

		if m.TournamentInfo[i] != nil {
			if err := m.TournamentInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tournamentInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TournamentInfoMinimalDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TournamentInfoMinimalDTO) UnmarshalBinary(b []byte) error {
	var res TournamentInfoMinimalDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
