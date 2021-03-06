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

// LolHonorV2HonorRecipient lol honor v2 honor recipient
// swagger:model LolHonorV2HonorRecipient
type LolHonorV2HonorRecipient struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// honors
	Honors []*LolHonorV2Honor `json:"honors"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol honor v2 honor recipient
func (m *LolHonorV2HonorRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHonors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolHonorV2HonorRecipient) validateHonors(formats strfmt.Registry) error {

	if swag.IsZero(m.Honors) { // not required
		return nil
	}

	for i := 0; i < len(m.Honors); i++ {
		if swag.IsZero(m.Honors[i]) { // not required
			continue
		}

		if m.Honors[i] != nil {
			if err := m.Honors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("honors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolHonorV2HonorRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHonorV2HonorRecipient) UnmarshalBinary(b []byte) error {
	var res LolHonorV2HonorRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
