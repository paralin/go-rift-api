// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathLoginSession lol npe tutorial path login session
// swagger:model LolNpeTutorialPathLoginSession
type LolNpeTutorialPathLoginSession struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// is new player
	IsNewPlayer bool `json:"isNewPlayer,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol npe tutorial path login session
func (m *LolNpeTutorialPathLoginSession) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathLoginSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathLoginSession) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathLoginSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
