// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolMatchHistorySummonerIDAndName lol match history summoner Id and name
// swagger:model LolMatchHistorySummonerIdAndName
type LolMatchHistorySummonerIDAndName struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`
}

// Validate validates this lol match history summoner Id and name
func (m *LolMatchHistorySummonerIDAndName) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolMatchHistorySummonerIDAndName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMatchHistorySummonerIDAndName) UnmarshalBinary(b []byte) error {
	var res LolMatchHistorySummonerIDAndName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}