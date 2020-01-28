// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolSummonerSummonerDTO lol summoner summoner d t o
// swagger:model LolSummonerSummonerDTO
type LolSummonerSummonerDTO struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// exp points
	ExpPoints int32 `json:"expPoints,omitempty"`

	// exp to next level
	ExpToNextLevel int32 `json:"expToNextLevel,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// level
	Level int32 `json:"level,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// name change flag
	NameChangeFlag bool `json:"nameChangeFlag,omitempty"`

	// partner Id
	PartnerID string `json:"partnerId,omitempty"`

	// profile icon Id
	ProfileIconID int32 `json:"profileIconId,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`
}

// Validate validates this lol summoner summoner d t o
func (m *LolSummonerSummonerDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolSummonerSummonerDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSummonerSummonerDTO) UnmarshalBinary(b []byte) error {
	var res LolSummonerSummonerDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
