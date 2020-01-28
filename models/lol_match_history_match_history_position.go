// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolMatchHistoryMatchHistoryPosition lol match history match history position
// swagger:model LolMatchHistoryMatchHistoryPosition
type LolMatchHistoryMatchHistoryPosition struct {

	// x
	X int64 `json:"x,omitempty"`

	// y
	Y int64 `json:"y,omitempty"`
}

// Validate validates this lol match history match history position
func (m *LolMatchHistoryMatchHistoryPosition) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolMatchHistoryMatchHistoryPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolMatchHistoryMatchHistoryPosition) UnmarshalBinary(b []byte) error {
	var res LolMatchHistoryMatchHistoryPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}