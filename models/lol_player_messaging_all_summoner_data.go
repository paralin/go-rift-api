// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolPlayerMessagingAllSummonerData lol player messaging all summoner data
// swagger:model LolPlayerMessagingAllSummonerData
type LolPlayerMessagingAllSummonerData struct {

	// summoner
	Summoner *LolPlayerMessagingSummoner `json:"summoner,omitempty"`
}

// Validate validates this lol player messaging all summoner data
func (m *LolPlayerMessagingAllSummonerData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummoner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPlayerMessagingAllSummonerData) validateSummoner(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolPlayerMessagingAllSummonerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPlayerMessagingAllSummonerData) UnmarshalBinary(b []byte) error {
	var res LolPlayerMessagingAllSummonerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
