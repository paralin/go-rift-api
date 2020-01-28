// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStatstonesSummoner lol statstones summoner
// swagger:model LolStatstonesSummoner
type LolStatstonesSummoner struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol statstones summoner
func (m *LolStatstonesSummoner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesSummoner) UnmarshalBinary(b []byte) error {
	var res LolStatstonesSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
