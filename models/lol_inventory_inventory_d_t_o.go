// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolInventoryInventoryDTO lol inventory inventory d t o
// swagger:model LolInventoryInventoryDTO
type LolInventoryInventoryDTO struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// expires
	Expires string `json:"expires,omitempty"`

	// items
	Items map[string]interface{} `json:"items,omitempty"`

	// items jwt
	ItemsJwt string `json:"itemsJwt,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol inventory inventory d t o
func (m *LolInventoryInventoryDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolInventoryInventoryDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolInventoryInventoryDTO) UnmarshalBinary(b []byte) error {
	var res LolInventoryInventoryDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
