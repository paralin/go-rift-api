// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEndOfGameSimpleMessage lol end of game simple message
// swagger:model LolEndOfGameSimpleMessage
type LolEndOfGameSimpleMessage struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// msg Id
	MsgID string `json:"msgId,omitempty"`

	// params
	Params []string `json:"params"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol end of game simple message
func (m *LolEndOfGameSimpleMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameSimpleMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameSimpleMessage) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameSimpleMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
