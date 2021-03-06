// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolHonorV2MutualHonorPlayer lol honor v2 mutual honor player
// swagger:model LolHonorV2MutualHonorPlayer
type LolHonorV2MutualHonorPlayer struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// skin index
	SkinIndex int32 `json:"skinIndex,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol honor v2 mutual honor player
func (m *LolHonorV2MutualHonorPlayer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolHonorV2MutualHonorPlayer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHonorV2MutualHonorPlayer) UnmarshalBinary(b []byte) error {
	var res LolHonorV2MutualHonorPlayer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
