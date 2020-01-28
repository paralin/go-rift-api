// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClubsClubDto lol clubs club dto
// swagger:model LolClubsClubDto
type LolClubsClubDto struct {

	// club data
	ClubData *LolClubsClubDataDto `json:"clubData,omitempty"`

	// membership
	Membership *LolClubsClubMembershipDto `json:"membership,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// resource Uri
	ResourceURI string `json:"resourceUri,omitempty"`
}

// Validate validates this lol clubs club dto
func (m *LolClubsClubDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClubData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembership(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClubsClubDto) validateClubData(formats strfmt.Registry) error {

	if swag.IsZero(m.ClubData) { // not required
		return nil
	}

	if m.ClubData != nil {
		if err := m.ClubData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clubData")
			}
			return err
		}
	}

	return nil
}

func (m *LolClubsClubDto) validateMembership(formats strfmt.Registry) error {

	if swag.IsZero(m.Membership) { // not required
		return nil
	}

	if m.Membership != nil {
		if err := m.Membership.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("membership")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsClubDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsClubDto) UnmarshalBinary(b []byte) error {
	var res LolClubsClubDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}