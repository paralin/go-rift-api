// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolAccountVerificationIsVerifiedResponse lol account verification is verified response
// swagger:model LolAccountVerificationIsVerifiedResponse
type LolAccountVerificationIsVerifiedResponse struct {

	// message
	Message string `json:"message,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this lol account verification is verified response
func (m *LolAccountVerificationIsVerifiedResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolAccountVerificationIsVerifiedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolAccountVerificationIsVerifiedResponse) UnmarshalBinary(b []byte) error {
	var res LolAccountVerificationIsVerifiedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}