// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolSuggestedPlayersSuggestedPlayersSuggestedPlayer lol suggested players suggested players suggested player
// swagger:model LolSuggestedPlayersSuggestedPlayersSuggestedPlayer
type LolSuggestedPlayersSuggestedPlayersSuggestedPlayer struct {

	// common friend Id
	CommonFriendID int64 `json:"commonFriendId,omitempty"`

	// common friend name
	CommonFriendName string `json:"commonFriendName,omitempty"`

	// reason
	Reason LolSuggestedPlayersSuggestedPlayersReason `json:"reason,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`

	// summoner name
	SummonerName string `json:"summonerName,omitempty"`
}

// Validate validates this lol suggested players suggested players suggested player
func (m *LolSuggestedPlayersSuggestedPlayersSuggestedPlayer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolSuggestedPlayersSuggestedPlayersSuggestedPlayer) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := m.Reason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersSuggestedPlayer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolSuggestedPlayersSuggestedPlayersSuggestedPlayer) UnmarshalBinary(b []byte) error {
	var res LolSuggestedPlayersSuggestedPlayersSuggestedPlayer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
