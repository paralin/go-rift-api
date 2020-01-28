// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRecommendationsAcsChampionGames lol recommendations acs champion games
// swagger:model LolRecommendationsAcsChampionGames
type LolRecommendationsAcsChampionGames struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// lane
	Lane string `json:"lane,omitempty"`

	// queue
	Queue int32 `json:"queue,omitempty"`

	// role
	Role string `json:"role,omitempty"`
}

// Validate validates this lol recommendations acs champion games
func (m *LolRecommendationsAcsChampionGames) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRecommendationsAcsChampionGames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRecommendationsAcsChampionGames) UnmarshalBinary(b []byte) error {
	var res LolRecommendationsAcsChampionGames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}