// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashNoShowPingResponseData lol clash no show ping response data
// swagger:model LolClashNoShowPingResponseData
type LolClashNoShowPingResponseData struct {

	// dodge time
	DodgeTime int64 `json:"dodgeTime,omitempty"`

	// gameflow state
	GameflowState LolClashGameflowPhase `json:"gameflowState,omitempty"`

	// is playmode restricted
	IsPlaymodeRestricted bool `json:"isPlaymodeRestricted,omitempty"`

	// login time
	LoginTime int64 `json:"loginTime,omitempty"`

	// ready check info
	ReadyCheckInfo *LolClashReadyCheckInfo `json:"readyCheckInfo,omitempty"`
}

// Validate validates this lol clash no show ping response data
func (m *LolClashNoShowPingResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameflowState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadyCheckInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashNoShowPingResponseData) validateGameflowState(formats strfmt.Registry) error {

	if swag.IsZero(m.GameflowState) { // not required
		return nil
	}

	if err := m.GameflowState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gameflowState")
		}
		return err
	}

	return nil
}

func (m *LolClashNoShowPingResponseData) validateReadyCheckInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ReadyCheckInfo) { // not required
		return nil
	}

	if m.ReadyCheckInfo != nil {
		if err := m.ReadyCheckInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readyCheckInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashNoShowPingResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashNoShowPingResponseData) UnmarshalBinary(b []byte) error {
	var res LolClashNoShowPingResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
