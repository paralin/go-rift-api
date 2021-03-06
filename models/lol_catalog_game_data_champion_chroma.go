// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolCatalogGameDataChampionChroma lol catalog game data champion chroma
// swagger:model LolCatalogGameDataChampionChroma
type LolCatalogGameDataChampionChroma struct {

	// chroma path
	ChromaPath string `json:"chromaPath,omitempty"`

	// colors
	Colors []string `json:"colors"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this lol catalog game data champion chroma
func (m *LolCatalogGameDataChampionChroma) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolCatalogGameDataChampionChroma) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogGameDataChampionChroma) UnmarshalBinary(b []byte) error {
	var res LolCatalogGameDataChampionChroma
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
