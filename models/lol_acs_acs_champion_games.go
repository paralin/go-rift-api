// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolAcsAcsChampionGames lol acs acs champion games
// swagger:model LolAcsAcsChampionGames
type LolAcsAcsChampionGames struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// lane
	Lane string `json:"lane,omitempty"`

	// queue
	Queue int32 `json:"queue,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this lol acs acs champion games
func (m *LolAcsAcsChampionGames) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolAcsAcsChampionGames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolAcsAcsChampionGames) UnmarshalBinary(b []byte) error {
	var res LolAcsAcsChampionGames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
