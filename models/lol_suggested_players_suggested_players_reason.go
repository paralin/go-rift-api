// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// LolSuggestedPlayersSuggestedPlayersReason lol suggested players suggested players reason
// swagger:model LolSuggestedPlayersSuggestedPlayersReason
type LolSuggestedPlayersSuggestedPlayersReason string

const (

	// LolSuggestedPlayersSuggestedPlayersReasonPreviousPremade captures enum value "PreviousPremade"
	LolSuggestedPlayersSuggestedPlayersReasonPreviousPremade LolSuggestedPlayersSuggestedPlayersReason = "PreviousPremade"

	// LolSuggestedPlayersSuggestedPlayersReasonOnlineFriend captures enum value "OnlineFriend"
	LolSuggestedPlayersSuggestedPlayersReasonOnlineFriend LolSuggestedPlayersSuggestedPlayersReason = "OnlineFriend"

	// LolSuggestedPlayersSuggestedPlayersReasonFriendOfFriend captures enum value "FriendOfFriend"
	LolSuggestedPlayersSuggestedPlayersReasonFriendOfFriend LolSuggestedPlayersSuggestedPlayersReason = "FriendOfFriend"

	// LolSuggestedPlayersSuggestedPlayersReasonVictoriousComrade captures enum value "VictoriousComrade"
	LolSuggestedPlayersSuggestedPlayersReasonVictoriousComrade LolSuggestedPlayersSuggestedPlayersReason = "VictoriousComrade"

	// LolSuggestedPlayersSuggestedPlayersReasonLegacyPlayAgain captures enum value "LegacyPlayAgain"
	LolSuggestedPlayersSuggestedPlayersReasonLegacyPlayAgain LolSuggestedPlayersSuggestedPlayersReason = "LegacyPlayAgain"
)

// for schema
var lolSuggestedPlayersSuggestedPlayersReasonEnum []interface{}

func init() {
	var res []LolSuggestedPlayersSuggestedPlayersReason
	if err := json.Unmarshal([]byte(`["PreviousPremade","OnlineFriend","FriendOfFriend","VictoriousComrade","LegacyPlayAgain"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolSuggestedPlayersSuggestedPlayersReasonEnum = append(lolSuggestedPlayersSuggestedPlayersReasonEnum, v)
	}
}

func (m LolSuggestedPlayersSuggestedPlayersReason) validateLolSuggestedPlayersSuggestedPlayersReasonEnum(path, location string, value LolSuggestedPlayersSuggestedPlayersReason) error {
	if err := validate.Enum(path, location, value, lolSuggestedPlayersSuggestedPlayersReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol suggested players suggested players reason
func (m LolSuggestedPlayersSuggestedPlayersReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolSuggestedPlayersSuggestedPlayersReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
