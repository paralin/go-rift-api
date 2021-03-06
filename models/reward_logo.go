// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RewardLogo reward logo
// swagger:model RewardLogo
type RewardLogo struct {

	// logo
	Logo int32 `json:"logo,omitempty"`

	// member owned count
	MemberOwnedCount int32 `json:"memberOwnedCount,omitempty"`
}

// Validate validates this reward logo
func (m *RewardLogo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RewardLogo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RewardLogo) UnmarshalBinary(b []byte) error {
	var res RewardLogo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
