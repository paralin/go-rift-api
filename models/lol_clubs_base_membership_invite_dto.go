// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClubsBaseMembershipInviteDto lol clubs base membership invite dto
// swagger:model LolClubsBaseMembershipInviteDto
type LolClubsBaseMembershipInviteDto struct {

	// club key
	ClubKey string `json:"clubKey,omitempty"`

	// club name
	ClubName string `json:"clubName,omitempty"`

	// invitee platform Id
	InviteePlatformID string `json:"inviteePlatformId,omitempty"`

	// invitee summoner Id
	InviteeSummonerID int64 `json:"inviteeSummonerId,omitempty"`

	// inviter account Id
	InviterAccountID int64 `json:"inviterAccountId,omitempty"`

	// inviter platform Id
	InviterPlatformID string `json:"inviterPlatformId,omitempty"`

	// inviter summoner Id
	InviterSummonerID int64 `json:"inviterSummonerId,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this lol clubs base membership invite dto
func (m *LolClubsBaseMembershipInviteDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsBaseMembershipInviteDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsBaseMembershipInviteDto) UnmarshalBinary(b []byte) error {
	var res LolClubsBaseMembershipInviteDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
