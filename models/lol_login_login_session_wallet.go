// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLoginLoginSessionWallet lol login login session wallet
// swagger:model LolLoginLoginSessionWallet
type LolLoginLoginSessionWallet struct {

	// ip
	IP int64 `json:"ip,omitempty"`

	// rp
	Rp int64 `json:"rp,omitempty"`
}

// Validate validates this lol login login session wallet
func (m *LolLoginLoginSessionWallet) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLoginLoginSessionWallet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLoginLoginSessionWallet) UnmarshalBinary(b []byte) error {
	var res LolLoginLoginSessionWallet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
