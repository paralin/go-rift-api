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

// LolRankedLeagueDivisionInfo lol ranked league division info
// swagger:model LolRankedLeagueDivisionInfo
type LolRankedLeagueDivisionInfo struct {

	// apex unlock time millis
	ApexUnlockTimeMillis int64 `json:"apexUnlockTimeMillis,omitempty"`

	// division
	Division LolRankedLeagueDivision `json:"division,omitempty"`

	// max league size
	MaxLeagueSize int64 `json:"maxLeagueSize,omitempty"`

	// standings
	Standings []*LolRankedLeagueStanding `json:"standings"`

	// tier
	Tier LolRankedLeagueTier `json:"tier,omitempty"`
}

// Validate validates this lol ranked league division info
func (m *LolRankedLeagueDivisionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolRankedLeagueDivisionInfo) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if err := m.Division.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("division")
		}
		return err
	}

	return nil
}

func (m *LolRankedLeagueDivisionInfo) validateStandings(formats strfmt.Registry) error {

	if swag.IsZero(m.Standings) { // not required
		return nil
	}

	for i := 0; i < len(m.Standings); i++ {
		if swag.IsZero(m.Standings[i]) { // not required
			continue
		}

		if m.Standings[i] != nil {
			if err := m.Standings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolRankedLeagueDivisionInfo) validateTier(formats strfmt.Registry) error {

	if swag.IsZero(m.Tier) { // not required
		return nil
	}

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolRankedLeagueDivisionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRankedLeagueDivisionInfo) UnmarshalBinary(b []byte) error {
	var res LolRankedLeagueDivisionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
