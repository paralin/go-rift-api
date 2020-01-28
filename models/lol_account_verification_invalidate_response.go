// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolAccountVerificationInvalidateResponse lol account verification invalidate response
// swagger:model LolAccountVerificationInvalidateResponse
type LolAccountVerificationInvalidateResponse struct {

	// message
	Message string `json:"message,omitempty"`

	// sms token expire duration in sec
	SmsTokenExpireDurationInSec int32 `json:"smsTokenExpireDurationInSec,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this lol account verification invalidate response
func (m *LolAccountVerificationInvalidateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolAccountVerificationInvalidateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolAccountVerificationInvalidateResponse) UnmarshalBinary(b []byte) error {
	var res LolAccountVerificationInvalidateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}