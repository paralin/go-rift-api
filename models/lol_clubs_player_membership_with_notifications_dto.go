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

// LolClubsPlayerMembershipWithNotificationsDto lol clubs player membership with notifications dto
// swagger:model LolClubsPlayerMembershipWithNotificationsDto
type LolClubsPlayerMembershipWithNotificationsDto struct {

	// membership notifications
	MembershipNotifications []*LolClubsMembershipNoficationsDto `json:"membershipNotifications"`

	// player membership
	PlayerMembership *LolClubsPlayerMembershipDto `json:"playerMembership,omitempty"`
}

// Validate validates this lol clubs player membership with notifications dto
func (m *LolClubsPlayerMembershipWithNotificationsDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembershipNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayerMembership(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClubsPlayerMembershipWithNotificationsDto) validateMembershipNotifications(formats strfmt.Registry) error {

	if swag.IsZero(m.MembershipNotifications) { // not required
		return nil
	}

	for i := 0; i < len(m.MembershipNotifications); i++ {
		if swag.IsZero(m.MembershipNotifications[i]) { // not required
			continue
		}

		if m.MembershipNotifications[i] != nil {
			if err := m.MembershipNotifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("membershipNotifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClubsPlayerMembershipWithNotificationsDto) validatePlayerMembership(formats strfmt.Registry) error {

	if swag.IsZero(m.PlayerMembership) { // not required
		return nil
	}

	if m.PlayerMembership != nil {
		if err := m.PlayerMembership.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("playerMembership")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsPlayerMembershipWithNotificationsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsPlayerMembershipWithNotificationsDto) UnmarshalBinary(b []byte) error {
	var res LolClubsPlayerMembershipWithNotificationsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
