// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLeaverBusterSummoner lol leaver buster summoner
// swagger:model LolLeaverBusterSummoner
type LolLeaverBusterSummoner struct {

	// acct Id
	AcctID int64 `json:"acctId,omitempty"`

	// sum Id
	SumID int64 `json:"sumId,omitempty"`
}

// Validate validates this lol leaver buster summoner
func (m *LolLeaverBusterSummoner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLeaverBusterSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLeaverBusterSummoner) UnmarshalBinary(b []byte) error {
	var res LolLeaverBusterSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}