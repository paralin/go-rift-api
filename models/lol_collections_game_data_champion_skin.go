// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolCollectionsGameDataChampionSkin lol collections game data champion skin
// swagger:model LolCollectionsGameDataChampionSkin
type LolCollectionsGameDataChampionSkin struct {

	// id
	ID int32 `json:"id,omitempty"`

	// is base
	IsBase bool `json:"isBase,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// splash path
	SplashPath string `json:"splashPath,omitempty"`

	// splash video path
	SplashVideoPath string `json:"splashVideoPath,omitempty"`
}

// Validate validates this lol collections game data champion skin
func (m *LolCollectionsGameDataChampionSkin) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolCollectionsGameDataChampionSkin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCollectionsGameDataChampionSkin) UnmarshalBinary(b []byte) error {
	var res LolCollectionsGameDataChampionSkin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
