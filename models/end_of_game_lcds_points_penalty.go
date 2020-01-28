// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EndOfGameLcdsPointsPenalty end of game lcds points penalty
// swagger:model EndOfGameLcdsPointsPenalty
type EndOfGameLcdsPointsPenalty struct {

	// penalty
	Penalty float64 `json:"penalty,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this end of game lcds points penalty
func (m *EndOfGameLcdsPointsPenalty) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndOfGameLcdsPointsPenalty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndOfGameLcdsPointsPenalty) UnmarshalBinary(b []byte) error {
	var res EndOfGameLcdsPointsPenalty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
