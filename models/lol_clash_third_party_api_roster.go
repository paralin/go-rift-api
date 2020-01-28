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

// LolClashThirdPartyAPIRoster lol clash third party Api roster
// swagger:model LolClashThirdPartyApiRoster
type LolClashThirdPartyAPIRoster struct {

	// captain
	Captain *LolClashThirdPartyAPIPlayer `json:"captain,omitempty"`

	// members
	Members []*LolClashThirdPartyAPIPlayer `json:"members"`
}

// Validate validates this lol clash third party Api roster
func (m *LolClashThirdPartyAPIRoster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaptain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashThirdPartyAPIRoster) validateCaptain(formats strfmt.Registry) error {

	if swag.IsZero(m.Captain) { // not required
		return nil
	}

	if m.Captain != nil {
		if err := m.Captain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("captain")
			}
			return err
		}
	}

	return nil
}

func (m *LolClashThirdPartyAPIRoster) validateMembers(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *LolClashThirdPartyAPIRoster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashThirdPartyAPIRoster) UnmarshalBinary(b []byte) error {
	var res LolClashThirdPartyAPIRoster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
