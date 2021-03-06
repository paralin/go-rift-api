// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LCDSGlobalRewards l c d s global rewards
// swagger:model LCDSGlobalRewards
type LCDSGlobalRewards struct {

	// all champions
	AllChampions bool `json:"allChampions,omitempty"`
}

// Validate validates this l c d s global rewards
func (m *LCDSGlobalRewards) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LCDSGlobalRewards) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LCDSGlobalRewards) UnmarshalBinary(b []byte) error {
	var res LCDSGlobalRewards
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
