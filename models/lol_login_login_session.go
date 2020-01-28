// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLoginLoginSession lol login login session
// swagger:model LolLoginLoginSession
type LolLoginLoginSession struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// connected
	Connected bool `json:"connected,omitempty"`

	// error
	Error *LolLoginLoginError `json:"error,omitempty"`

	// gas token
	GasToken interface{} `json:"gasToken,omitempty"`

	// id token
	IDToken string `json:"idToken,omitempty"`

	// is in login queue
	IsInLoginQueue bool `json:"isInLoginQueue,omitempty"`

	// is new player
	IsNewPlayer bool `json:"isNewPlayer,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`

	// state
	State LolLoginLoginSessionStates `json:"state,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// user auth token
	UserAuthToken string `json:"userAuthToken,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this lol login login session
func (m *LolLoginLoginSession) Validate(formats strfmt.Registry) error {
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

func (m *LolLoginLoginSession) validateError(formats strfmt.Registry) error {

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

func (m *LolLoginLoginSession) validateState(formats strfmt.Registry) error {

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
func (m *LolLoginLoginSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoginLoginSession) UnmarshalBinary(b []byte) error {
	var res LolLoginLoginSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}