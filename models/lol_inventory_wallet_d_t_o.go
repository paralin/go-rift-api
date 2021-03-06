// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolInventoryWalletDTO lol inventory wallet d t o
// swagger:model LolInventoryWalletDTO
type LolInventoryWalletDTO struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// balances
	Balances map[string]int32 `json:"balances,omitempty"`

	// balances jwt
	BalancesJwt string `json:"balancesJwt,omitempty"`

	// expires
	Expires string `json:"expires,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`
}

// Validate validates this lol inventory wallet d t o
func (m *LolInventoryWalletDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolInventoryWalletDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolInventoryWalletDTO) UnmarshalBinary(b []byte) error {
	var res LolInventoryWalletDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
