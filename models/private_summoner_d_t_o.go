// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PrivateSummonerDTO private summoner d t o
// swagger:model PrivateSummonerDTO
type PrivateSummonerDTO struct {

	// acct Id
	AcctID int64 `json:"acctId,omitempty"`

	// internal name
	InternalName string `json:"internalName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// name change flag
	NameChangeFlag bool `json:"nameChangeFlag,omitempty"`

	// profile icon Id
	ProfileIconID int32 `json:"profileIconId,omitempty"`

	// sum Id
	SumID int64 `json:"sumId,omitempty"`
}

// Validate validates this private summoner d t o
func (m *PrivateSummonerDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateSummonerDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateSummonerDTO) UnmarshalBinary(b []byte) error {
	var res PrivateSummonerDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
