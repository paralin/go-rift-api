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

// LolRsoAuthAuthorizationRequest lol rso auth authorization request
// swagger:model LolRsoAuthAuthorizationRequest
type LolRsoAuthAuthorizationRequest struct {

	// claims
	Claims []string `json:"claims"`

	// client Id
	ClientID string `json:"clientId,omitempty"`

	// scope
	Scope []string `json:"scope"`

	// trust levels
	TrustLevels []LolRsoAuthRSOAuthorizationTrustLevel `json:"trustLevels"`
}

// Validate validates this lol rso auth authorization request
func (m *LolRsoAuthAuthorizationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrustLevels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolRsoAuthAuthorizationRequest) validateTrustLevels(formats strfmt.Registry) error {

	if swag.IsZero(m.TrustLevels) { // not required
		return nil
	}

	for i := 0; i < len(m.TrustLevels); i++ {

		if err := m.TrustLevels[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trustLevels" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolRsoAuthAuthorizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRsoAuthAuthorizationRequest) UnmarshalBinary(b []byte) error {
	var res LolRsoAuthAuthorizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
