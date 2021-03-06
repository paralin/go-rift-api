// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolClashRoster lol clash roster
// swagger:model LolClashRoster
type LolClashRoster struct {

	// available logos
	AvailableLogos []*RewardLogo `json:"availableLogos"`

	// captain summoner Id
	CaptainSummonerID int64 `json:"captainSummonerId,omitempty"`

	// current bracket wins
	CurrentBracketWins int32 `json:"currentBracketWins,omitempty"`

	// high tier variance
	HighTierVariance bool `json:"highTierVariance,omitempty"`

	// icon color Id
	IconColorID int32 `json:"iconColorId,omitempty"`

	// icon Id
	IconID int32 `json:"iconId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invitation Id
	InvitationID string `json:"invitationId,omitempty"`

	// is active in current phase
	IsActiveInCurrentPhase bool `json:"isActiveInCurrentPhase,omitempty"`

	// is clash banned
	IsClashBanned bool `json:"isClashBanned,omitempty"`

	// is current bracket complete
	IsCurrentBracketComplete bool `json:"isCurrentBracketComplete,omitempty"`

	// is eliminated
	IsEliminated bool `json:"isEliminated,omitempty"`

	// is registered
	IsRegistered bool `json:"isRegistered,omitempty"`

	// losses
	Losses int32 `json:"losses,omitempty"`

	// members
	Members []*LolClashRosterMember `json:"members"`

	// name
	Name string `json:"name,omitempty"`

	// num completed periods
	NumCompletedPeriods int32 `json:"numCompletedPeriods,omitempty"`

	// phase infos
	PhaseInfos []*LolClashRosterPhaseInfo `json:"phaseInfos"`

	// points
	Points int32 `json:"points,omitempty"`

	// short name
	ShortName string `json:"shortName,omitempty"`

	// suggested invites
	SuggestedInvites []*LolClashSuggestedInvite `json:"suggestedInvites"`

	// tier
	Tier int32 `json:"tier,omitempty"`

	// tournament Id
	TournamentID int64 `json:"tournamentId,omitempty"`

	// wins
	Wins int32 `json:"wins,omitempty"`

	// withdraw
	Withdraw *RosterWithdraw `json:"withdraw,omitempty"`
}

// Validate validates this lol clash roster
func (m *LolClashRoster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableLogos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuggestedInvites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithdraw(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolClashRoster) validateAvailableLogos(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableLogos) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableLogos); i++ {
		if swag.IsZero(m.AvailableLogos[i]) { // not required
			continue
		}

		if m.AvailableLogos[i] != nil {
			if err := m.AvailableLogos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableLogos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClashRoster) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClashRoster) validatePhaseInfos(formats strfmt.Registry) error {

	if swag.IsZero(m.PhaseInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.PhaseInfos); i++ {
		if swag.IsZero(m.PhaseInfos[i]) { // not required
			continue
		}

		if m.PhaseInfos[i] != nil {
			if err := m.PhaseInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phaseInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClashRoster) validateSuggestedInvites(formats strfmt.Registry) error {

	if swag.IsZero(m.SuggestedInvites) { // not required
		return nil
	}

	for i := 0; i < len(m.SuggestedInvites); i++ {
		if swag.IsZero(m.SuggestedInvites[i]) { // not required
			continue
		}

		if m.SuggestedInvites[i] != nil {
			if err := m.SuggestedInvites[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suggestedInvites" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolClashRoster) validateWithdraw(formats strfmt.Registry) error {

	if swag.IsZero(m.Withdraw) { // not required
		return nil
	}

	if m.Withdraw != nil {
		if err := m.Withdraw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("withdraw")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolClashRoster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashRoster) UnmarshalBinary(b []byte) error {
	var res LolClashRoster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
