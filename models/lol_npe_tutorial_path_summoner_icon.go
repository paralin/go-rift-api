// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathSummonerIcon lol npe tutorial path summoner icon
// swagger:model LolNpeTutorialPathSummonerIcon
type LolNpeTutorialPathSummonerIcon struct {

	// profile icon Id
	ProfileIconID int32 `json:"profileIconId,omitempty"`
}

// Validate validates this lol npe tutorial path summoner icon
func (m *LolNpeTutorialPathSummonerIcon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathSummonerIcon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathSummonerIcon) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathSummonerIcon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
