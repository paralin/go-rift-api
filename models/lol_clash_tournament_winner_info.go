// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClashTournamentWinnerInfo lol clash tournament winner info
// swagger:model LolClashTournamentWinnerInfo
type LolClashTournamentWinnerInfo struct {

	// average win duration
	AverageWinDuration int64 `json:"averageWinDuration,omitempty"`

	// create time
	CreateTime int64 `json:"createTime,omitempty"`

	// logo
	Logo int32 `json:"logo,omitempty"`

	// logo color
	LogoColor int32 `json:"logoColor,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// player ids
	PlayerIds []int64 `json:"playerIds"`

	// roster Id
	RosterID int64 `json:"rosterId,omitempty"`

	// short name
	ShortName string `json:"shortName,omitempty"`

	// tier
	Tier int32 `json:"tier,omitempty"`
}

// Validate validates this lol clash tournament winner info
func (m *LolClashTournamentWinnerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClashTournamentWinnerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashTournamentWinnerInfo) UnmarshalBinary(b []byte) error {
	var res LolClashTournamentWinnerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
