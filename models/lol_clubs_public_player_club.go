// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClubsPublicPlayerClub lol clubs public player club
// swagger:model LolClubsPublicPlayerClub
type LolClubsPublicPlayerClub struct {

	// key
	Key string `json:"key,omitempty"`

	// members
	Members *LolClubsPublicClubMemberLists `json:"members,omitempty"`
}

// Validate validates this lol clubs public player club
func (m *LolClubsPublicPlayerClub) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClubsPublicPlayerClub) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	if m.Members != nil {
		if err := m.Members.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("members")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsPublicPlayerClub) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsPublicPlayerClub) UnmarshalBinary(b []byte) error {
	var res LolClubsPublicPlayerClub
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
