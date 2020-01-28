// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolGameflowGameflowGameMap lol gameflow gameflow game map
// swagger:model LolGameflowGameflowGameMap
type LolGameflowGameflowGameMap struct {

	// assets
	Assets interface{} `json:"assets,omitempty"`

	// categorized content bundles
	CategorizedContentBundles interface{} `json:"categorizedContentBundles,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// game mode name
	GameModeName string `json:"gameModeName,omitempty"`

	// game mode short name
	GameModeShortName string `json:"gameModeShortName,omitempty"`

	// game mutator
	GameMutator string `json:"gameMutator,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is r g m
	IsRGM bool `json:"isRGM,omitempty"`

	// map string Id
	MapStringID string `json:"mapStringId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// platform name
	PlatformName string `json:"platformName,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`
}

// Validate validates this lol gameflow gameflow game map
func (m *LolGameflowGameflowGameMap) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolGameflowGameflowGameMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGameflowGameflowGameMap) UnmarshalBinary(b []byte) error {
	var res LolGameflowGameflowGameMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}