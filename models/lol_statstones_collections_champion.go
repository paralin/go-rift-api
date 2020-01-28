// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStatstonesCollectionsChampion lol statstones collections champion
// swagger:model LolStatstonesCollectionsChampion
type LolStatstonesCollectionsChampion struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// square portrait path
	SquarePortraitPath string `json:"squarePortraitPath,omitempty"`
}

// Validate validates this lol statstones collections champion
func (m *LolStatstonesCollectionsChampion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStatstonesCollectionsChampion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStatstonesCollectionsChampion) UnmarshalBinary(b []byte) error {
	var res LolStatstonesCollectionsChampion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}