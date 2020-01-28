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

// LolAcsAcsChampionGamesCollection lol acs acs champion games collection
// swagger:model LolAcsAcsChampionGamesCollection
type LolAcsAcsChampionGamesCollection struct {

	// champions
	Champions []*LolAcsAcsChampionGames `json:"champions"`

	// game count
	GameCount int32 `json:"gameCount,omitempty"`
}

// Validate validates this lol acs acs champion games collection
func (m *LolAcsAcsChampionGamesCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChampions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolAcsAcsChampionGamesCollection) validateChampions(formats strfmt.Registry) error {

	if swag.IsZero(m.Champions) { // not required
		return nil
	}

	for i := 0; i < len(m.Champions); i++ {
		if swag.IsZero(m.Champions[i]) { // not required
			continue
		}

		if m.Champions[i] != nil {
			if err := m.Champions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("champions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolAcsAcsChampionGamesCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolAcsAcsChampionGamesCollection) UnmarshalBinary(b []byte) error {
	var res LolAcsAcsChampionGamesCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}