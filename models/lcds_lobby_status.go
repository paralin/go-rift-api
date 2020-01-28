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

// LcdsLobbyStatus lcds lobby status
// swagger:model LcdsLobbyStatus
type LcdsLobbyStatus struct {

	// chat key
	ChatKey string `json:"chatKey,omitempty"`

	// game meta data
	GameMetaData string `json:"gameMetaData,omitempty"`

	// invitation Id
	InvitationID string `json:"invitationId,omitempty"`

	// invitees
	Invitees []*LcdsInvitee `json:"invitees"`

	// members
	Members []*LcdsMember `json:"members"`

	// owner
	Owner *LcdsPlayer `json:"owner,omitempty"`
}

// Validate validates this lcds lobby status
func (m *LcdsLobbyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvitees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LcdsLobbyStatus) validateInvitees(formats strfmt.Registry) error {

	if swag.IsZero(m.Invitees) { // not required
		return nil
	}

	for i := 0; i < len(m.Invitees); i++ {
		if swag.IsZero(m.Invitees[i]) { // not required
			continue
		}

		if m.Invitees[i] != nil {
			if err := m.Invitees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invitees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LcdsLobbyStatus) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LcdsLobbyStatus) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LcdsLobbyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LcdsLobbyStatus) UnmarshalBinary(b []byte) error {
	var res LcdsLobbyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}