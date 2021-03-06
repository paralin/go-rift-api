// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashMatchmakingSearchResource lol clash matchmaking search resource
// swagger:model LolClashMatchmakingSearchResource
type LolClashMatchmakingSearchResource struct {

	// dodge data
	DodgeData *LolClashMatchmakingDodgeData `json:"dodgeData,omitempty"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`

	// ready check
	ReadyCheck *LolClashMatchmakingReadyCheckResource `json:"readyCheck,omitempty"`
}

// Validate validates this lol clash matchmaking search resource
func (m *LolClashMatchmakingSearchResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDodgeData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadyCheck(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashMatchmakingSearchResource) validateDodgeData(formats strfmt.Registry) error {

	if swag.IsZero(m.DodgeData) { // not required
		return nil
	}

	if m.DodgeData != nil {
		if err := m.DodgeData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dodgeData")
			}
			return err
		}
	}

	return nil
}

func (m *LolClashMatchmakingSearchResource) validateReadyCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.ReadyCheck) { // not required
		return nil
	}

	if m.ReadyCheck != nil {
		if err := m.ReadyCheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readyCheck")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashMatchmakingSearchResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashMatchmakingSearchResource) UnmarshalBinary(b []byte) error {
	var res LolClashMatchmakingSearchResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
