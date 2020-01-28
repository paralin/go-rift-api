// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolRsoAuthAuthorizationResponse lol rso auth authorization response
// swagger:model LolRsoAuthAuthorizationResponse
type LolRsoAuthAuthorizationResponse struct {

	// authorization
	Authorization *LolRsoAuthImplicitAuthorization `json:"authorization,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol rso auth authorization response
func (m *LolRsoAuthAuthorizationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolRsoAuthAuthorizationResponse) validateAuthorization(formats strfmt.Registry) error {

	if swag.IsZero(m.Authorization) { // not required
		return nil
	}

	if m.Authorization != nil {
		if err := m.Authorization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolRsoAuthAuthorizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRsoAuthAuthorizationResponse) UnmarshalBinary(b []byte) error {
	var res LolRsoAuthAuthorizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
