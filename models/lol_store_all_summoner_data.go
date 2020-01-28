// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolStoreAllSummonerData lol store all summoner data
// swagger:model LolStoreAllSummonerData
type LolStoreAllSummonerData struct {

	// summoner
	Summoner *LolStoreSummoner `json:"summoner,omitempty"`

	// summoner level and points
	SummonerLevelAndPoints *LolStoreSummonerLevelAndPoints `json:"summonerLevelAndPoints,omitempty"`
}

// Validate validates this lol store all summoner data
func (m *LolStoreAllSummonerData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummoner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummonerLevelAndPoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolStoreAllSummonerData) validateSummoner(formats strfmt.Registry) error {

	if swag.IsZero(m.Summoner) { // not required
		return nil
	}

	if m.Summoner != nil {
		if err := m.Summoner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summoner")
			}
			return err
		}
	}

	return nil
}

func (m *LolStoreAllSummonerData) validateSummonerLevelAndPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.SummonerLevelAndPoints) { // not required
		return nil
	}

	if m.SummonerLevelAndPoints != nil {
		if err := m.SummonerLevelAndPoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summonerLevelAndPoints")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreAllSummonerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreAllSummonerData) UnmarshalBinary(b []byte) error {
	var res LolStoreAllSummonerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}