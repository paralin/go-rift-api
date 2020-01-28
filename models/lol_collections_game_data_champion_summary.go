// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolCollectionsGameDataChampionSummary lol collections game data champion summary
// swagger:model LolCollectionsGameDataChampionSummary
type LolCollectionsGameDataChampionSummary struct {

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this lol collections game data champion summary
func (m *LolCollectionsGameDataChampionSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolCollectionsGameDataChampionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCollectionsGameDataChampionSummary) UnmarshalBinary(b []byte) error {
	var res LolCollectionsGameDataChampionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}