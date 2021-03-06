// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolMissionsCollectionsChampionSkin lol missions collections champion skin
// swagger:model LolMissionsCollectionsChampionSkin
type LolMissionsCollectionsChampionSkin struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// ownership
	Ownership *LolMissionsCollectionsOwnership `json:"ownership,omitempty"`
}

// Validate validates this lol missions collections champion skin
func (m *LolMissionsCollectionsChampionSkin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnership(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolMissionsCollectionsChampionSkin) validateOwnership(formats strfmt.Registry) error {

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
func (m *LolMissionsCollectionsChampionSkin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMissionsCollectionsChampionSkin) UnmarshalBinary(b []byte) error {
	var res LolMissionsCollectionsChampionSkin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
