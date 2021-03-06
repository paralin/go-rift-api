// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LCDSPlayerMessagingSimpleMessageResponse l c d s player messaging simple message response
// swagger:model LCDSPlayerMessagingSimpleMessageResponse
type LCDSPlayerMessagingSimpleMessageResponse struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// command
	Command string `json:"command,omitempty"`

	// msg Id
	MsgID string `json:"msgId,omitempty"`
}

// Validate validates this l c d s player messaging simple message response
func (m *LCDSPlayerMessagingSimpleMessageResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LCDSPlayerMessagingSimpleMessageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LCDSPlayerMessagingSimpleMessageResponse) UnmarshalBinary(b []byte) error {
	var res LCDSPlayerMessagingSimpleMessageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
