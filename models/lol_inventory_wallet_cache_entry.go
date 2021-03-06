// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolInventoryWalletCacheEntry lol inventory wallet cache entry
// swagger:model LolInventoryWalletCacheEntry
type LolInventoryWalletCacheEntry struct {

	// expiration m s
	ExpirationMS int64 `json:"expirationMS,omitempty"`

	// issued at m s
	IssuedAtMS int64 `json:"issuedAtMS,omitempty"`

	// received at m s
	ReceivedAtMS int64 `json:"receivedAtMS,omitempty"`

	// signed balances jwt
	SignedBalancesJwt string `json:"signedBalancesJwt,omitempty"`

	// valid
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this lol inventory wallet cache entry
func (m *LolInventoryWalletCacheEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolInventoryWalletCacheEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolInventoryWalletCacheEntry) UnmarshalBinary(b []byte) error {
	var res LolInventoryWalletCacheEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
