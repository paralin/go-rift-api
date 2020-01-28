// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectLegacyChampSelectBannedChampions lol champ select legacy champ select banned champions
// swagger:model LolChampSelectLegacyChampSelectBannedChampions
type LolChampSelectLegacyChampSelectBannedChampions struct {

	// my team bans
	MyTeamBans []int32 `json:"myTeamBans"`

	// num bans
	NumBans int32 `json:"numBans,omitempty"`

	// their team bans
	TheirTeamBans []int32 `json:"theirTeamBans"`
}

// Validate validates this lol champ select legacy champ select banned champions
func (m *LolChampSelectLegacyChampSelectBannedChampions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectLegacyChampSelectBannedChampions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectLegacyChampSelectBannedChampions) UnmarshalBinary(b []byte) error {
	var res LolChampSelectLegacyChampSelectBannedChampions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
