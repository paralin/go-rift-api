// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEndOfGameGameDataChampionSummary lol end of game game data champion summary
// swagger:model LolEndOfGameGameDataChampionSummary
type LolEndOfGameGameDataChampionSummary struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// square portrait path
	SquarePortraitPath string `json:"squarePortraitPath,omitempty"`
}

// Validate validates this lol end of game game data champion summary
func (m *LolEndOfGameGameDataChampionSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameGameDataChampionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameGameDataChampionSummary) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameGameDataChampionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}