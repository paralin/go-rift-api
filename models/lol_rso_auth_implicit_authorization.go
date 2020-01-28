// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolRsoAuthImplicitAuthorization lol rso auth implicit authorization
// swagger:model LolRsoAuthImplicitAuthorization
type LolRsoAuthImplicitAuthorization struct {

	// access token
	AccessToken *LolRsoAuthAccessToken `json:"accessToken,omitempty"`

	// id token
	IDToken *LolRsoAuthIDToken `json:"idToken,omitempty"`
}

// Validate validates this lol rso auth implicit authorization
func (m *LolRsoAuthImplicitAuthorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolRsoAuthImplicitAuthorization) validateAccessToken(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessToken) { // not required
		return nil
	}

	if m.AccessToken != nil {
		if err := m.AccessToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessToken")
			}
			return err
		}
	}

	return nil
}

func (m *LolRsoAuthImplicitAuthorization) validateIDToken(formats strfmt.Registry) error {

	if swag.IsZero(m.IDToken) { // not required
		return nil
	}

	if m.IDToken != nil {
		if err := m.IDToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idToken")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolRsoAuthImplicitAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRsoAuthImplicitAuthorization) UnmarshalBinary(b []byte) error {
	var res LolRsoAuthImplicitAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
