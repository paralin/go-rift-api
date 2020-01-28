// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStoreSummoner lol store summoner
// swagger:model LolStoreSummoner
type LolStoreSummoner struct {

	// acct Id
	AcctID int64 `json:"acctId,omitempty"`

	// sum Id
	SumID int64 `json:"sumId,omitempty"`
}

// Validate validates this lol store summoner
func (m *LolStoreSummoner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreSummoner) UnmarshalBinary(b []byte) error {
	var res LolStoreSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
