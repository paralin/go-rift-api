// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPlayerMessagingSimpleMessage lol player messaging simple message
// swagger:model LolPlayerMessagingSimpleMessage
type LolPlayerMessagingSimpleMessage struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// body code
	BodyCode string `json:"bodyCode,omitempty"`

	// msg Id
	MsgID string `json:"msgId,omitempty"`

	// params
	Params []string `json:"params"`

	// title code
	TitleCode string `json:"titleCode,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol player messaging simple message
func (m *LolPlayerMessagingSimpleMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPlayerMessagingSimpleMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPlayerMessagingSimpleMessage) UnmarshalBinary(b []byte) error {
	var res LolPlayerMessagingSimpleMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
