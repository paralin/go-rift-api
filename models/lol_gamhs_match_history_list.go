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

// LolGamhsMatchHistoryList lol gamhs match history list
// swagger:model LolGamhsMatchHistoryList
type LolGamhsMatchHistoryList struct {

	// active puuid
	ActivePuuid string `json:"active_puuid,omitempty"`

	// games
	Games []*LolGamhsMatchHistoryData `json:"games"`
}

// Validate validates this lol gamhs match history list
func (m *LolGamhsMatchHistoryList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolGamhsMatchHistoryList) validateGames(formats strfmt.Registry) error {

	if swag.IsZero(m.Games) { // not required
		return nil
	}

	for i := 0; i < len(m.Games); i++ {
		if swag.IsZero(m.Games[i]) { // not required
			continue
		}

		if m.Games[i] != nil {
			if err := m.Games[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("games" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolGamhsMatchHistoryList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGamhsMatchHistoryList) UnmarshalBinary(b []byte) error {
	var res LolGamhsMatchHistoryList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
