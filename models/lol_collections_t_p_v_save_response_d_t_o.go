// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolCollectionsTPVSaveResponseDTO lol collections t p v save response d t o
// swagger:model LolCollectionsTPVSaveResponseDTO
type LolCollectionsTPVSaveResponseDTO struct {

	// data
	Data string `json:"data,omitempty"`
}

// Validate validates this lol collections t p v save response d t o
func (m *LolCollectionsTPVSaveResponseDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolCollectionsTPVSaveResponseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCollectionsTPVSaveResponseDTO) UnmarshalBinary(b []byte) error {
	var res LolCollectionsTPVSaveResponseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
