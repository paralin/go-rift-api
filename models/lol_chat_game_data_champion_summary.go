// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatGameDataChampionSummary lol chat game data champion summary
// swagger:model LolChatGameDataChampionSummary
type LolChatGameDataChampionSummary struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this lol chat game data champion summary
func (m *LolChatGameDataChampionSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatGameDataChampionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatGameDataChampionSummary) UnmarshalBinary(b []byte) error {
	var res LolChatGameDataChampionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}