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

// LolEndOfGameLoginDataPacket lol end of game login data packet
// swagger:model LolEndOfGameLoginDataPacket
type LolEndOfGameLoginDataPacket struct {

	// all summoner data
	AllSummonerData *LolEndOfGameLoginSummonerData `json:"allSummonerData,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// simple messages
	SimpleMessages []*LolEndOfGameSimpleMessage `json:"simpleMessages"`
}

// Validate validates this lol end of game login data packet
func (m *LolEndOfGameLoginDataPacket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllSummonerData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSimpleMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolEndOfGameLoginDataPacket) validateAllSummonerData(formats strfmt.Registry) error {

	if swag.IsZero(m.AllSummonerData) { // not required
		return nil
	}

	if m.AllSummonerData != nil {
		if err := m.AllSummonerData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allSummonerData")
			}
			return err
		}
	}

	return nil
}

func (m *LolEndOfGameLoginDataPacket) validateSimpleMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.SimpleMessages) { // not required
		return nil
	}

	for i := 0; i < len(m.SimpleMessages); i++ {
		if swag.IsZero(m.SimpleMessages[i]) { // not required
			continue
		}

		if m.SimpleMessages[i] != nil {
			if err := m.SimpleMessages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("simpleMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameLoginDataPacket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameLoginDataPacket) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameLoginDataPacket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
