// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolFeaturedModesCollectionsSummonerIcon lol featured modes collections summoner icon
// swagger:model LolFeaturedModesCollectionsSummonerIcon
type LolFeaturedModesCollectionsSummonerIcon struct {

	// icon Id
	IconID int32 `json:"iconId,omitempty"`

	// ownership
	Ownership *LolFeaturedModesCollectionsOwnership `json:"ownership,omitempty"`
}

// Validate validates this lol featured modes collections summoner icon
func (m *LolFeaturedModesCollectionsSummonerIcon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnership(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolFeaturedModesCollectionsSummonerIcon) validateOwnership(formats strfmt.Registry) error {

	if swag.IsZero(m.Ownership) { // not required
		return nil
	}

	if m.Ownership != nil {
		if err := m.Ownership.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownership")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolFeaturedModesCollectionsSummonerIcon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolFeaturedModesCollectionsSummonerIcon) UnmarshalBinary(b []byte) error {
	var res LolFeaturedModesCollectionsSummonerIcon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
