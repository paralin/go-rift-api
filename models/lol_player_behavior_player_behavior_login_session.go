// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolPlayerBehaviorPlayerBehaviorLoginSession lol player behavior player behavior login session
// swagger:model LolPlayerBehaviorPlayerBehavior_LoginSession
type LolPlayerBehaviorPlayerBehaviorLoginSession struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// error
	Error *LolPlayerBehaviorPlayerBehaviorLoginError `json:"error,omitempty"`

	// state
	State LolPlayerBehaviorPlayerBehaviorLoginSessionState `json:"state,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol player behavior player behavior login session
func (m *LolPlayerBehaviorPlayerBehaviorLoginSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPlayerBehaviorPlayerBehaviorLoginSession) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *LolPlayerBehaviorPlayerBehaviorLoginSession) validateState(formats strfmt.Registry) error {

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
func (m *LolPlayerBehaviorPlayerBehaviorLoginSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPlayerBehaviorPlayerBehaviorLoginSession) UnmarshalBinary(b []byte) error {
	var res LolPlayerBehaviorPlayerBehaviorLoginSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
