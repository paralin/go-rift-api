// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LeaguesLcdsSimpleMessage leagues lcds simple message
// swagger:model LeaguesLcdsSimpleMessage
type LeaguesLcdsSimpleMessage struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// msg Id
	MsgID string `json:"msgId,omitempty"`

	// params
	Params []string `json:"params"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this leagues lcds simple message
func (m *LeaguesLcdsSimpleMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LeaguesLcdsSimpleMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaguesLcdsSimpleMessage) UnmarshalBinary(b []byte) error {
	var res LeaguesLcdsSimpleMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}