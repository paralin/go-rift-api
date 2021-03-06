// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolSummonerXpAndLevelMessage lol summoner xp and level message
// swagger:model LolSummonerXpAndLevelMessage
type LolSummonerXpAndLevelMessage struct {

	// level
	Level *LolSummonerLevelField `json:"level,omitempty"`

	// xp
	Xp interface{} `json:"xp,omitempty"`
}

// Validate validates this lol summoner xp and level message
func (m *LolSummonerXpAndLevelMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolSummonerXpAndLevelMessage) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	if m.Level != nil {
		if err := m.Level.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("level")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolSummonerXpAndLevelMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSummonerXpAndLevelMessage) UnmarshalBinary(b []byte) error {
	var res LolSummonerXpAndLevelMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
