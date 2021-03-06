// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRevivalsEligibility lol revivals eligibility
// swagger:model LolRevivalsEligibility
type LolRevivalsEligibility struct {

	// is eligible
	IsEligible bool `json:"isEligible,omitempty"`

	// is first touch
	IsFirstTouch bool `json:"isFirstTouch,omitempty"`
}

// Validate validates this lol revivals eligibility
func (m *LolRevivalsEligibility) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRevivalsEligibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRevivalsEligibility) UnmarshalBinary(b []byte) error {
	var res LolRevivalsEligibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
