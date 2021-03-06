// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPerksGameCustomizationDTO lol perks game customization d t o
// swagger:model LolPerksGameCustomizationDTO
type LolPerksGameCustomizationDTO struct {

	// category
	Category string `json:"category,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// is teambuilder
	IsTeambuilder bool `json:"isTeambuilder,omitempty"`

	// queue type
	QueueType int64 `json:"queueType,omitempty"`
}

// Validate validates this lol perks game customization d t o
func (m *LolPerksGameCustomizationDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPerksGameCustomizationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPerksGameCustomizationDTO) UnmarshalBinary(b []byte) error {
	var res LolPerksGameCustomizationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
