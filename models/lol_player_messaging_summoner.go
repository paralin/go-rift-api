// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPlayerMessagingSummoner lol player messaging summoner
// swagger:model LolPlayerMessagingSummoner
type LolPlayerMessagingSummoner struct {

	// acct Id
	AcctID int64 `json:"acctId,omitempty"`

	// sum Id
	SumID int64 `json:"sumId,omitempty"`
}

// Validate validates this lol player messaging summoner
func (m *LolPlayerMessagingSummoner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPlayerMessagingSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPlayerMessagingSummoner) UnmarshalBinary(b []byte) error {
	var res LolPlayerMessagingSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
