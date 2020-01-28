// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectChampSelectBannedChampions lol champ select champ select banned champions
// swagger:model LolChampSelectChampSelectBannedChampions
type LolChampSelectChampSelectBannedChampions struct {

	// my team bans
	MyTeamBans []int32 `json:"myTeamBans"`

	// num bans
	NumBans int32 `json:"numBans,omitempty"`

	// their team bans
	TheirTeamBans []int32 `json:"theirTeamBans"`
}

// Validate validates this lol champ select champ select banned champions
func (m *LolChampSelectChampSelectBannedChampions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectChampSelectBannedChampions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectChampSelectBannedChampions) UnmarshalBinary(b []byte) error {
	var res LolChampSelectChampSelectBannedChampions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
