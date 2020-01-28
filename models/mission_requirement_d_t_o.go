// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MissionRequirementDTO mission requirement d t o
// swagger:model MissionRequirementDTO
type MissionRequirementDTO struct {

	// expression
	Expression string `json:"expression,omitempty"`
}

// Validate validates this mission requirement d t o
func (m *MissionRequirementDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MissionRequirementDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MissionRequirementDTO) UnmarshalBinary(b []byte) error {
	var res MissionRequirementDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
