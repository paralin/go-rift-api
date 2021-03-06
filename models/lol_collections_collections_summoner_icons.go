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

// LolCollectionsCollectionsSummonerIcons lol collections collections summoner icons
// swagger:model LolCollectionsCollectionsSummonerIcons
type LolCollectionsCollectionsSummonerIcons struct {

	// icons
	Icons []int32 `json:"icons"`

	// summoner icons
	SummonerIcons []*LolCollectionsCollectionsSummonerIcon `json:"summonerIcons"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol collections collections summoner icons
func (m *LolCollectionsCollectionsSummonerIcons) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummonerIcons(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCollectionsCollectionsSummonerIcons) validateSummonerIcons(formats strfmt.Registry) error {

	if swag.IsZero(m.SummonerIcons) { // not required
		return nil
	}

	for i := 0; i < len(m.SummonerIcons); i++ {
		if swag.IsZero(m.SummonerIcons[i]) { // not required
			continue
		}

		if m.SummonerIcons[i] != nil {
			if err := m.SummonerIcons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summonerIcons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCollectionsCollectionsSummonerIcons) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCollectionsCollectionsSummonerIcons) UnmarshalBinary(b []byte) error {
	var res LolCollectionsCollectionsSummonerIcons
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
