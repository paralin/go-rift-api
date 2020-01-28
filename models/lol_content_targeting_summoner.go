// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolContentTargetingSummoner lol content targeting summoner
// swagger:model LolContentTargetingSummoner
type LolContentTargetingSummoner struct {

	// profile icon Id
	ProfileIconID int32 `json:"profileIconId,omitempty"`

	// summoner level
	SummonerLevel int32 `json:"summonerLevel,omitempty"`
}

// Validate validates this lol content targeting summoner
func (m *LolContentTargetingSummoner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolContentTargetingSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolContentTargetingSummoner) UnmarshalBinary(b []byte) error {
	var res LolContentTargetingSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}