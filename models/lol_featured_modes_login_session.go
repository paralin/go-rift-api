// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolFeaturedModesLoginSession lol featured modes login session
// swagger:model LolFeaturedModesLoginSession
type LolFeaturedModesLoginSession struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// state
	State LolFeaturedModesLoginSessionStates `json:"state,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol featured modes login session
func (m *LolFeaturedModesLoginSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolFeaturedModesLoginSession) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolFeaturedModesLoginSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolFeaturedModesLoginSession) UnmarshalBinary(b []byte) error {
	var res LolFeaturedModesLoginSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
