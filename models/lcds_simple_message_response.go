// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LcdsSimpleMessageResponse lcds simple message response
// swagger:model LcdsSimpleMessageResponse
type LcdsSimpleMessageResponse struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// command
	Command string `json:"command,omitempty"`

	// msg Id
	MsgID string `json:"msgId,omitempty"`
}

// Validate validates this lcds simple message response
func (m *LcdsSimpleMessageResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LcdsSimpleMessageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LcdsSimpleMessageResponse) UnmarshalBinary(b []byte) error {
	var res LcdsSimpleMessageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
